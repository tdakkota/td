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

// ChannelsDeleteChannelRequest represents TL type `channels.deleteChannel#c0111fe3`.
// Delete a channel/supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.deleteChannel for reference.
type ChannelsDeleteChannelRequest struct {
	// Channel/supergroup¹ to delete
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Channel InputChannelClass
}

// ChannelsDeleteChannelRequestTypeID is TL type id of ChannelsDeleteChannelRequest.
const ChannelsDeleteChannelRequestTypeID = 0xc0111fe3

func (d *ChannelsDeleteChannelRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Channel == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *ChannelsDeleteChannelRequest) String() string {
	if d == nil {
		return "ChannelsDeleteChannelRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelsDeleteChannelRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tChannel: ")
	sb.WriteString(fmt.Sprint(d.Channel))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (d *ChannelsDeleteChannelRequest) TypeID() uint32 {
	return ChannelsDeleteChannelRequestTypeID
}

// Encode implements bin.Encoder.
func (d *ChannelsDeleteChannelRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode channels.deleteChannel#c0111fe3 as nil")
	}
	b.PutID(ChannelsDeleteChannelRequestTypeID)
	if d.Channel == nil {
		return fmt.Errorf("unable to encode channels.deleteChannel#c0111fe3: field channel is nil")
	}
	if err := d.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.deleteChannel#c0111fe3: field channel: %w", err)
	}
	return nil
}

// GetChannel returns value of Channel field.
func (d *ChannelsDeleteChannelRequest) GetChannel() (value InputChannelClass) {
	return d.Channel
}

// Decode implements bin.Decoder.
func (d *ChannelsDeleteChannelRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode channels.deleteChannel#c0111fe3 to nil")
	}
	if err := b.ConsumeID(ChannelsDeleteChannelRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.deleteChannel#c0111fe3: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.deleteChannel#c0111fe3: field channel: %w", err)
		}
		d.Channel = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsDeleteChannelRequest.
var (
	_ bin.Encoder = &ChannelsDeleteChannelRequest{}
	_ bin.Decoder = &ChannelsDeleteChannelRequest{}
)

// ChannelsDeleteChannel invokes method channels.deleteChannel#c0111fe3 returning error if any.
// Delete a channel/supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 CHANNEL_TOO_LARGE: Channel is too large to be deleted; this error is issued when trying to delete channels with more than 1000 members (subject to change)
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  403 CHAT_WRITE_FORBIDDEN: You can't write in this chat
//
// See https://core.telegram.org/method/channels.deleteChannel for reference.
func (c *Client) ChannelsDeleteChannel(ctx context.Context, channel InputChannelClass) (UpdatesClass, error) {
	var result UpdatesBox

	request := &ChannelsDeleteChannelRequest{
		Channel: channel,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
