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

// StatsGetMessagePublicForwardsRequest represents TL type `stats.getMessagePublicForwards#5630281b`.
// Obtains a list of messages, indicating to which other public channels was a channel message forwarded.
// Will return a list of messages¹ with peer_id equal to the public channel to which this message was forwarded.
//
// Links:
//  1) https://core.telegram.org/constructor/message
//
// See https://core.telegram.org/method/stats.getMessagePublicForwards for reference.
type StatsGetMessagePublicForwardsRequest struct {
	// Source channel
	Channel InputChannelClass
	// Source message ID
	MsgID int
	// Initially 0, then set to the next_rate parameter of messages.messagesSlice¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/messages.messagesSlice
	OffsetRate int
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetPeer InputPeerClass
	// Offsets for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	OffsetID int
	// Maximum number of results to return, see pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
}

// StatsGetMessagePublicForwardsRequestTypeID is TL type id of StatsGetMessagePublicForwardsRequest.
const StatsGetMessagePublicForwardsRequestTypeID = 0x5630281b

func (g *StatsGetMessagePublicForwardsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Channel == nil) {
		return false
	}
	if !(g.MsgID == 0) {
		return false
	}
	if !(g.OffsetRate == 0) {
		return false
	}
	if !(g.OffsetPeer == nil) {
		return false
	}
	if !(g.OffsetID == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *StatsGetMessagePublicForwardsRequest) String() string {
	if g == nil {
		return "StatsGetMessagePublicForwardsRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StatsGetMessagePublicForwardsRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tChannel: ")
	sb.WriteString(fmt.Sprint(g.Channel))
	sb.WriteString(",\n")
	sb.WriteString("\tMsgID: ")
	sb.WriteString(fmt.Sprint(g.MsgID))
	sb.WriteString(",\n")
	sb.WriteString("\tOffsetRate: ")
	sb.WriteString(fmt.Sprint(g.OffsetRate))
	sb.WriteString(",\n")
	sb.WriteString("\tOffsetPeer: ")
	sb.WriteString(fmt.Sprint(g.OffsetPeer))
	sb.WriteString(",\n")
	sb.WriteString("\tOffsetID: ")
	sb.WriteString(fmt.Sprint(g.OffsetID))
	sb.WriteString(",\n")
	sb.WriteString("\tLimit: ")
	sb.WriteString(fmt.Sprint(g.Limit))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *StatsGetMessagePublicForwardsRequest) TypeID() uint32 {
	return StatsGetMessagePublicForwardsRequestTypeID
}

// Encode implements bin.Encoder.
func (g *StatsGetMessagePublicForwardsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stats.getMessagePublicForwards#5630281b as nil")
	}
	b.PutID(StatsGetMessagePublicForwardsRequestTypeID)
	if g.Channel == nil {
		return fmt.Errorf("unable to encode stats.getMessagePublicForwards#5630281b: field channel is nil")
	}
	if err := g.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.getMessagePublicForwards#5630281b: field channel: %w", err)
	}
	b.PutInt(g.MsgID)
	b.PutInt(g.OffsetRate)
	if g.OffsetPeer == nil {
		return fmt.Errorf("unable to encode stats.getMessagePublicForwards#5630281b: field offset_peer is nil")
	}
	if err := g.OffsetPeer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.getMessagePublicForwards#5630281b: field offset_peer: %w", err)
	}
	b.PutInt(g.OffsetID)
	b.PutInt(g.Limit)
	return nil
}

// GetChannel returns value of Channel field.
func (g *StatsGetMessagePublicForwardsRequest) GetChannel() (value InputChannelClass) {
	return g.Channel
}

// GetMsgID returns value of MsgID field.
func (g *StatsGetMessagePublicForwardsRequest) GetMsgID() (value int) {
	return g.MsgID
}

// GetOffsetRate returns value of OffsetRate field.
func (g *StatsGetMessagePublicForwardsRequest) GetOffsetRate() (value int) {
	return g.OffsetRate
}

// GetOffsetPeer returns value of OffsetPeer field.
func (g *StatsGetMessagePublicForwardsRequest) GetOffsetPeer() (value InputPeerClass) {
	return g.OffsetPeer
}

// GetOffsetID returns value of OffsetID field.
func (g *StatsGetMessagePublicForwardsRequest) GetOffsetID() (value int) {
	return g.OffsetID
}

// GetLimit returns value of Limit field.
func (g *StatsGetMessagePublicForwardsRequest) GetLimit() (value int) {
	return g.Limit
}

// Decode implements bin.Decoder.
func (g *StatsGetMessagePublicForwardsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stats.getMessagePublicForwards#5630281b to nil")
	}
	if err := b.ConsumeID(StatsGetMessagePublicForwardsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stats.getMessagePublicForwards#5630281b: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.getMessagePublicForwards#5630281b: field channel: %w", err)
		}
		g.Channel = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stats.getMessagePublicForwards#5630281b: field msg_id: %w", err)
		}
		g.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stats.getMessagePublicForwards#5630281b: field offset_rate: %w", err)
		}
		g.OffsetRate = value
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.getMessagePublicForwards#5630281b: field offset_peer: %w", err)
		}
		g.OffsetPeer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stats.getMessagePublicForwards#5630281b: field offset_id: %w", err)
		}
		g.OffsetID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stats.getMessagePublicForwards#5630281b: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StatsGetMessagePublicForwardsRequest.
var (
	_ bin.Encoder = &StatsGetMessagePublicForwardsRequest{}
	_ bin.Decoder = &StatsGetMessagePublicForwardsRequest{}
)

// StatsGetMessagePublicForwards invokes method stats.getMessagePublicForwards#5630281b returning error if any.
// Obtains a list of messages, indicating to which other public channels was a channel message forwarded.
// Will return a list of messages¹ with peer_id equal to the public channel to which this message was forwarded.
//
// Links:
//  1) https://core.telegram.org/constructor/message
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 MESSAGE_ID_INVALID: The provided message id is invalid
//
// See https://core.telegram.org/method/stats.getMessagePublicForwards for reference.
// Can be used by bots.
func (c *Client) StatsGetMessagePublicForwards(ctx context.Context, request *StatsGetMessagePublicForwardsRequest) (MessagesMessagesClass, error) {
	var result MessagesMessagesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
