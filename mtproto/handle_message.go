package mtproto

import (
	"fmt"

	"go.uber.org/zap"
	"golang.org/x/xerrors"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/internal/mt"
	"github.com/gotd/td/internal/proto"
)

func (c *Conn) handleMessage(b *bin.Buffer) error {
	id, err := b.PeekID()
	if err != nil {
		// Empty body.
		return xerrors.Errorf("failed to determine message type: %w", err)
	}

	c.logWithType(b).Debug("Handle message")
	switch id {
	case mt.NewSessionCreatedTypeID:
		return c.handleSessionCreated(b)
	case mt.BadMsgNotificationTypeID, mt.BadServerSaltTypeID:
		return c.handleBadMsg(b)
	case proto.MessageContainerTypeID:
		return c.handleContainer(b)
	case proto.ResultTypeID:
		return c.handleResult(b)
	case mt.PongTypeID:
		return c.handlePong(b)
	case mt.MsgsAckTypeID:
		return c.handleAck(b)
	case proto.GZIPTypeID:
		return c.handleGZIP(b)
	default:
		c.log.Info("Got message", zap.String("type_id", fmt.Sprintf("%x", id)))
		defer c.log.Info("Message consumed")

		return c.handler.OnMessage(b)
	}
}
