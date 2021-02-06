// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is

// StatsGetMessageStatsRequest represents TL type `stats.getMessageStats#b6e0a3f5`.
// Get message statistics¹
//
// Links:
//  1) https://core.telegram.org/api/stats
//
// See https://core.telegram.org/method/stats.getMessageStats for reference.
type StatsGetMessageStatsRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to enable dark theme for graph colors
	Dark bool
	// Channel ID
	Channel InputChannelClass
	// Message ID
	MsgID int
}

// StatsGetMessageStatsRequestTypeID is TL type id of StatsGetMessageStatsRequest.
const StatsGetMessageStatsRequestTypeID = 0xb6e0a3f5

func (g *StatsGetMessageStatsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Flags.Zero()) {
		return false
	}
	if !(g.Dark == false) {
		return false
	}
	if !(g.Channel == nil) {
		return false
	}
	if !(g.MsgID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *StatsGetMessageStatsRequest) String() string {
	if g == nil {
		return "StatsGetMessageStatsRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StatsGetMessageStatsRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(g.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tChannel: ")
	sb.WriteString(fmt.Sprint(g.Channel))
	sb.WriteString(",\n")
	sb.WriteString("\tMsgID: ")
	sb.WriteString(fmt.Sprint(g.MsgID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *StatsGetMessageStatsRequest) TypeID() uint32 {
	return StatsGetMessageStatsRequestTypeID
}

// Encode implements bin.Encoder.
func (g *StatsGetMessageStatsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stats.getMessageStats#b6e0a3f5 as nil")
	}
	b.PutID(StatsGetMessageStatsRequestTypeID)
	if !(g.Dark == false) {
		g.Flags.Set(0)
	}
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.getMessageStats#b6e0a3f5: field flags: %w", err)
	}
	if g.Channel == nil {
		return fmt.Errorf("unable to encode stats.getMessageStats#b6e0a3f5: field channel is nil")
	}
	if err := g.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.getMessageStats#b6e0a3f5: field channel: %w", err)
	}
	b.PutInt(g.MsgID)
	return nil
}

// SetDark sets value of Dark conditional field.
func (g *StatsGetMessageStatsRequest) SetDark(value bool) {
	if value {
		g.Flags.Set(0)
		g.Dark = true
	} else {
		g.Flags.Unset(0)
		g.Dark = false
	}
}

// GetDark returns value of Dark conditional field.
func (g *StatsGetMessageStatsRequest) GetDark() (value bool) {
	return g.Flags.Has(0)
}

// GetChannel returns value of Channel field.
func (g *StatsGetMessageStatsRequest) GetChannel() (value InputChannelClass) {
	return g.Channel
}

// GetMsgID returns value of MsgID field.
func (g *StatsGetMessageStatsRequest) GetMsgID() (value int) {
	return g.MsgID
}

// Decode implements bin.Decoder.
func (g *StatsGetMessageStatsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stats.getMessageStats#b6e0a3f5 to nil")
	}
	if err := b.ConsumeID(StatsGetMessageStatsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stats.getMessageStats#b6e0a3f5: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stats.getMessageStats#b6e0a3f5: field flags: %w", err)
		}
	}
	g.Dark = g.Flags.Has(0)
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.getMessageStats#b6e0a3f5: field channel: %w", err)
		}
		g.Channel = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stats.getMessageStats#b6e0a3f5: field msg_id: %w", err)
		}
		g.MsgID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StatsGetMessageStatsRequest.
var (
	_ bin.Encoder = &StatsGetMessageStatsRequest{}
	_ bin.Decoder = &StatsGetMessageStatsRequest{}
)

// StatsGetMessageStats invokes method stats.getMessageStats#b6e0a3f5 returning error if any.
// Get message statistics¹
//
// Links:
//  1) https://core.telegram.org/api/stats
//
// Possible errors:
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//
// See https://core.telegram.org/method/stats.getMessageStats for reference.
// Can be used by bots.
func (c *Client) StatsGetMessageStats(ctx context.Context, request *StatsGetMessageStatsRequest) (*StatsMessageStats, error) {
	var result StatsMessageStats

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
