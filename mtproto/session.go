package mtproto

import (
	"go.uber.org/zap"

	"github.com/gotd/td/internal/crypto"
)

// Session represents connection state.
type Session struct {
	ID   int64
	Key  crypto.AuthKey
	Salt int64
}

// Session returns current connection session info.
func (c *Conn) session() Session {
	c.sessionMux.RLock()
	defer c.sessionMux.RUnlock()
	return Session{
		Key:  c.authKey,
		Salt: c.salt,
		ID:   c.sessionID,
	}
}

// newSessionID sets session id to random value.
func (c *Conn) newSessionID() error {
	id, err := crypto.RandInt64(c.rand)
	if err != nil {
		return err
	}
	c.log.Debug("New session ID", zap.Int64("session_id", id))

	c.sessionMux.Lock()
	defer c.sessionMux.Unlock()
	c.sessionID = id

	return nil
}
