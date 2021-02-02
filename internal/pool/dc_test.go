package pool

import (
	"context"
	"errors"
	"github.com/gotd/td/clock"
	"go.uber.org/atomic"
	"go.uber.org/zap/zaptest"
	"golang.org/x/xerrors"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/internal/tdsync"
	"github.com/gotd/td/mtproto"
	"github.com/gotd/td/tg"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

type testProtoConn struct {
	handler mtproto.Handler
	invoke  func(ctx context.Context, input bin.Encoder, output bin.Decoder) error
	run     func(ctx context.Context, f func(ctx context.Context) error) error
}

func (t testProtoConn) InvokeRaw(ctx context.Context, input bin.Encoder, output bin.Decoder) error {
	if t.invoke != nil {
		return t.invoke(ctx, input, output)
	}

	return nil
}

func (t testProtoConn) defaultRun(ctx context.Context, f func(ctx context.Context) error) error {
	grp := tdsync.NewCancellableGroup(ctx)
	grp.Go(f)
	grp.Go(func(groupCtx context.Context) error {
		return t.handler.OnSession(mtproto.Session{})
	})
	grp.Go(func(groupCtx context.Context) error {
		<-ctx.Done()
		return nil
	})
	return grp.Wait()
}

func (t testProtoConn) Run(ctx context.Context, f func(ctx context.Context) error) error {
	if t.run != nil {
		return t.run(ctx, f)
	}

	return t.defaultRun(ctx, f)
}

type handler struct {
	log *zap.Logger
}

func (h handler) OnSession(addr string, cfg tg.Config, s mtproto.Session) error {
	h.log.Info("OnSession")
	return nil
}

func (h handler) OnMessage(b *bin.Buffer) error {
	h.log.Info("OnMessage")
	return nil
}

func testDC(creator protoCreator, handle ConnHandler, options DCOptions) func(t *testing.T) {
	return func(t *testing.T) {
		testCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		dc := NewDC(2, "10", handle, options)
		dc.protoCreator = creator

		grp := tdsync.NewCancellableGroup(testCtx)
		grp.Go(func(groupCtx context.Context) error {
			if err := dc.Run(groupCtx); err != nil {
				return xerrors.Errorf("run DC: %w", err)
			}
			return nil
		})
		grp.Go(func(groupCtx context.Context) error {
			<-groupCtx.Done()
			if err := dc.Close(testCtx); err != nil {
				return xerrors.Errorf("close DC: %w", err)
			}
			return nil
		})
		grp.Go(func(groupCtx context.Context) error {
			reqs := tdsync.NewCancellableGroup(groupCtx)
			defer grp.Cancel()

			for range [5]struct{}{} {
				reqs.Go(func(reqCtx context.Context) error {
					if err := dc.InvokeRaw(reqCtx, &tg.Config{}, &tg.Config{}); err != nil {
						return xerrors.Errorf("request: %w", err)
					}
					return nil
				})
			}
			return reqs.Wait()
		})

		if err := grp.Wait(); err != nil {
			require.ErrorIs(t, err, context.Canceled)
		}

		// Check InvokeRaw not work on closed DC.
		require.Error(t, dc.InvokeRaw(testCtx, &tg.Config{}, &tg.Config{}))
	}
}

func testDCMatrix(creator protoCreator) func(t *testing.T) {
	return func(t *testing.T) {
		connections := []int{0, 1, 2, 5, 10}
		for _, limit := range connections {
			t.Run(strconv.Itoa(limit), func(t *testing.T) {
				logger := zaptest.NewLogger(t)

				testDC(creator, handler{logger.Named("handler")}, DCOptions{
					Logger: logger.Named("dc"),
					MTProto: mtproto.Options{
						Logger: logger.Named("mtproto"),
						Clock:  clock.System,
					},
					MaxOpenConnections: int64(limit),
				})(t)
			})
		}
	}
}

func TestDC(t *testing.T) {
	dieErr := errors.New("жопир")
	t.Run("Die", func(t *testing.T) {
		t.Run("OnRun", testDCMatrix(func(addr string, opt mtproto.Options) protoConn {
			conn := testProtoConn{handler: opt.Handler}
			conn.run = func(ctx context.Context, f func(ctx context.Context) error) error {
				return conn.defaultRun(ctx, f)
			}

			return conn
		}))

		dead := atomic.NewBool(false)
		t.Run("OnSession", testDCMatrix(func(addr string, opt mtproto.Options) protoConn {
			conn := testProtoConn{handler: opt.Handler}
			conn.run = func(ctx context.Context, f func(ctx context.Context) error) error {
				grp := tdsync.NewCancellableGroup(ctx)
				grp.Go(f)
				grp.Go(func(groupCtx context.Context) error {
					if !dead.Swap(true) {
						return dieErr
					}
					return conn.handler.OnSession(mtproto.Session{})
				})
				grp.Go(func(groupCtx context.Context) error {
					<-groupCtx.Done()
					return nil
				})
				return grp.Wait()
			}

			return conn
		}))

		t.Run("OnInvoke", testDCMatrix(func(addr string, opt mtproto.Options) protoConn {
			conn := testProtoConn{handler: opt.Handler}
			conn.invoke = func(ctx context.Context, input bin.Encoder, output bin.Decoder) error {
				if rand.Int()%2 == 0 {
					return dieErr
				}
				return nil
			}

			return conn
		}))
	})
}
