package pool

import "go.uber.org/atomic"

type poolConn struct {
	*conn
	dc      *DC // immutable
	deleted *atomic.Bool
	id      int64
}
