package pool

import (
	"context"
	"fmt"
	"sync"

	"go.uber.org/atomic"
	"go.uber.org/zap"
	"golang.org/x/xerrors"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/internal/crypto"
	"github.com/gotd/td/internal/tdsync"
	"github.com/gotd/td/mtproto"
	"github.com/gotd/td/tg"
)

type protoCreator func(addr string, opt mtproto.Options) protoConn

// DCOptions is a Telegram data center connections pool options.
type DCOptions struct {
	// InitConnection parameters.
	// AppID of Telegram application.
	AppID int
	// Telegram device information.
	Device DeviceConfig

	// Logger is instance of zap.Logger. No logs by default.
	Logger *zap.Logger
	// MTProto options for connections.
	MTProto mtproto.Options
	// Opened connection limit to the DC.
	MaxOpenConnections int64
}

// DC represents connection pool to one data center.
type DC struct {
	id dcID
	// MTProto connection options.
	addr string          // immutable
	opts mtproto.Options // immutable
	// MTProto session parameters. Unique for DC.
	authKey    crypto.AuthKey
	salt       int64
	sessionMux sync.Mutex

	// InitConnection parameters.
	appID int // immutable
	// Telegram device information.
	device DeviceConfig // immutable

	// Creator of MTProto connections.
	protoCreator // immutable

	// Wrappers for external world, like logs or PRNG.
	log *zap.Logger // immutable

	// Handler passed by client.
	handler ConnHandler // immutable

	// DC context. Will be canceled by Run on exit.
	ctx    context.Context    // immutable
	cancel context.CancelFunc // immutable

	// Connections supervisor.
	grp *tdsync.Supervisor
	// Free connections.
	free []*poolConn
	// Total connections.
	total int64
	// Connection id monotonic counter.
	nextConn atomic.Int64
	freeReq  *reqMap
	// DC mutex.
	mu sync.Mutex

	// Limit of connections.
	max int64 // immutable

	// Signal connection for cases when all connections are dead, but all requests waiting for
	// free connection in 3rd acquire case.
	stuck *tdsync.ResetReady

	// Requests wait group.
	ongoing sync.WaitGroup

	// State of DC.
	ready  *tdsync.ResetReady
	closed atomic.Bool
}

// NewDC creates new uninitialized DC.
func NewDC(id dcID, addr string, handler ConnHandler, opts DCOptions) *DC {
	ctx, cancel := context.WithCancel(context.Background())

	return &DC{
		id:      id,
		addr:    addr,
		opts:    opts.MTProto,
		authKey: opts.MTProto.Key,
		salt:    opts.MTProto.Salt,
		appID:   opts.AppID,
		device:  opts.Device,
		protoCreator: protoCreator(func(addr string, opt mtproto.Options) protoConn {
			return mtproto.New(addr, opt)
		}),
		log:     opts.Logger,
		handler: handler,
		ctx:     ctx,
		cancel:  cancel,
		grp:     tdsync.NewSupervisor(ctx),
		freeReq: newReqMap(),
		max:     opts.MaxOpenConnections,
		stuck:   tdsync.NewResetReady(),
		ready:   tdsync.NewResetReady(),
	}
}

// OnSession implements ConnHandler.
func (c *DC) OnSession(addr string, cfg tg.Config, s mtproto.Session) error {
	c.sessionMux.Lock()
	c.salt = s.Salt
	noSessionBefore := c.authKey.Zero()
	c.authKey = s.Key
	c.sessionMux.Unlock()

	if noSessionBefore {
		c.log.Debug("DC Session saved")
	}

	return c.handler.OnSession(addr, cfg, s)
}

// OnMessage implements ConnHandler.
func (c *DC) OnMessage(b *bin.Buffer) error {
	id, _ := b.PeekID()
	c.log.Info("Got message", zap.String("type_id", fmt.Sprintf("%x", id)))
	defer c.log.Info("Message consumed")

	return c.handler.OnMessage(b)
}

func (c *DC) createConnection(id int64, mode connMode) *poolConn {
	opts := c.opts
	opts.Logger = c.log.Named("conn").With(zap.Int64("conn_id", id))
	// Load stored session.
	c.sessionMux.Lock()
	opts.Salt = c.salt
	opts.Key = c.authKey
	c.sessionMux.Unlock()

	conn := &poolConn{
		conn:    newConn(c, c.addr, mode, opts, c.protoCreator),
		dc:      c,
		id:      id,
		deleted: atomic.NewBool(false),
	}

	c.grp.Go(func(groupCtx context.Context) (err error) {
		if mode != connModeUpdates {
			defer c.dead(conn, err)
		}
		return conn.Init(groupCtx, c.appID, c.device)
	})

	return conn
}

func (c *DC) dead(r *poolConn, deadErr error) {
	if r.deleted.Swap(true) {
		return // Already deleted.
	}

	c.stuck.Reset()
	c.mu.Lock()
	defer c.mu.Unlock()

	c.total--
	remaining := c.total
	if remaining < 0 {
		panic("unreachable: remaining can'be less than zero")
	}

	idx := -1
	for i, conn := range c.free {
		// Search connection by pointer.
		if conn.id == r.id {
			idx = i
		}
	}

	if idx >= 0 {
		// Delete by index from slice tricks.
		copy(c.free[idx:], c.free[idx+1:])
		// Delete reference to prevent resource leaking.
		c.free[len(c.free)-1] = nil
		c.free = c.free[:len(c.free)-1]
	}

	r.log.Debug(
		"Connection died",
		zap.String("addr", c.addr),
		zap.Int64("remaining", remaining),
		zap.Int64("conn_id", r.id),
		zap.Error(deadErr),
	)
}

func (c *DC) pop() (r *poolConn, ok bool) {
	l := len(c.free)
	if l > 0 {
		r, c.free = c.free[l-1], c.free[:l-1]

		return r, true
	}

	return
}

func (c *DC) release(r *poolConn) {
	if r == nil {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.freeReq.transfer(r) {
		c.log.Debug("Transfer connection to requester", zap.Int64("conn_id", r.id))
		return
	}
	c.log.Debug("Connection released", zap.Int64("conn_id", r.id))
	c.free = append(c.free, r)
}

func (c *DC) request() (key reqKey, ch chan *poolConn) {
	return c.freeReq.request()
}

var errDCIsClosed = xerrors.New("DC is closed")

func (c *DC) acquire(ctx context.Context) (r *poolConn, err error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-c.ctx.Done(): // If DC forcibly closed — exit.
		return nil, xerrors.Errorf("DC closed: %w", c.ctx.Err())
	case <-c.ready.Ready(): // Await for ready if it not a initialization connection.
		c.log.Debug("Wait for DC ready")
	}

retry:
	c.mu.Lock()
	// 1st case: have free connections.
	if r, ok := c.pop(); ok {
		c.mu.Unlock()
		select {
		case <-r.Dead():
			c.dead(r, nil)
			goto retry
		default:
		}
		c.log.Debug("Re-using free connection", zap.Int64("conn_id", r.id))
		return r, nil
	}

	// 2nd case: no free connections, but can create one.
	// c.max < 1 means unlimited
	if c.max < 1 || c.total < c.max {
		c.total++
		c.mu.Unlock()

		id := c.nextConn.Inc()
		c.log.Debug(
			"Creating new connection",
			zap.String("addr", c.addr),
			zap.Int64("conn_id", id),
		)
		conn := c.createConnection(id, connModeData)
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-c.ctx.Done():
			return nil, xerrors.Errorf("DC closed: %w", c.ctx.Err())
		case <-conn.Ready():
			return conn, nil
		case <-conn.Dead():
			c.dead(conn, nil)
			goto retry
		}
	}

	// 3rd case: no free connections, can't create yet one, wait for free.
	key, ch := c.freeReq.request()
	c.mu.Unlock()
	c.log.Debug("Waiting for free connect", zap.Int64("request_id", int64(key)))

	select {
	case conn := <-ch:
		c.log.Debug("Got connection for request",
			zap.Int64("conn_id", conn.id),
			zap.Int64("request_id", int64(key)),
		)
		return conn, nil
	case <-c.stuck.Ready():
		c.log.Debug("Some connection dead, try to create new connection, cancel waiting")

		c.freeReq.delete(key)
		select {
		default:
		case conn, ok := <-ch:
			if ok && conn != nil {
				c.release(conn)
			}
		}

		goto retry
	case <-ctx.Done():
		err = ctx.Err()
	case <-c.ctx.Done():
		err = xerrors.Errorf("DC closed: %w", c.ctx.Err())
	}

	// Executed only if at least one of context is Done.
	c.freeReq.delete(key)
	select {
	default:
	case conn, ok := <-ch:
		if ok && conn != nil {
			c.release(conn)
		}
	}

	return nil, err
}

// Ready sends ready signal when DC is initialized.
func (c *DC) Ready() <-chan struct{} {
	return c.ready.Ready()
}

func (c *DC) keepAlive(ctx context.Context) error {
	c.log.Info("Starting updates connection")
	defer c.log.Info("Stopping updates connection")
	for {
		conn := c.createConnection(0, connModeUpdates)
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c.ctx.Done():
			return c.ctx.Err()
		case <-conn.Dead():
			c.log.Info("Updates connection dead on session creation, retry to connect", zap.Int64("conn_id", conn.id))
			continue
		case <-conn.Ready():
			if _, err := tg.NewClient(conn).HelpGetConfig(ctx); err != nil {
				c.log.Info("Get config failed", zap.Error(err))
			}
		}

		c.ready.Signal()
		c.log.Debug("Waiting for updates", zap.Int64("conn_id", conn.id))
		select {
		case <-ctx.Done():
			return ctx.Err()

		case <-c.ctx.Done():
			return c.ctx.Err()

		case <-conn.Dead():
			c.log.Info("Updates connection dead, retry to connect", zap.Int64("conn_id", conn.id))
			continue
		}
	}
}

// Run initialize connection pool.
func (c *DC) Run(ctx context.Context) error {
	if c.closed.Load() {
		return errDCIsClosed
	}

	return c.keepAlive(ctx)
}

// InvokeRaw sends MTProto request using one of pool connection.
func (c *DC) InvokeRaw(ctx context.Context, input bin.Encoder, output bin.Decoder) error {
	if c.closed.Load() {
		return errDCIsClosed
	}

	c.ongoing.Add(1)
	defer c.ongoing.Done()

	for {
		conn, err := c.acquire(ctx)
		if err != nil {
			if xerrors.Is(err, errConnDead) {
				continue
			}
			return xerrors.Errorf("acquire connection: %w", err)
		}

		c.log.Debug("DC Invoke")
		err = conn.InvokeRaw(ctx, input, output)
		c.release(conn)
		if err != nil {
			var rpcErr *mtproto.Error
			if !xerrors.As(err, &rpcErr) {
				c.log.Info("DC Invoke failed", zap.Error(err))
				continue
			}
		}

		c.log.Debug("DC Invoke complete")
		return err
	}
}

// Close waits while all ongoing requests will be done or until given context is done.
// Then, closes the DC.
func (c *DC) Close(closeCtx context.Context) error {
	if c.closed.Swap(true) {
		return xerrors.New("DC already closed")
	}
	c.log.Debug("Closing DC")
	defer c.log.Debug("DC closed")

	closed, cancel := context.WithCancel(closeCtx)
	go func() {
		c.ongoing.Wait()
		cancel()
	}()

	<-closed.Done()

	c.cancel()
	return c.grp.Wait()
}
