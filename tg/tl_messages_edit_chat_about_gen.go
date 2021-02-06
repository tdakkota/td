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

// MessagesEditChatAboutRequest represents TL type `messages.editChatAbout#def60797`.
// Edit the description of a group/supergroup/channel¹.
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/messages.editChatAbout for reference.
type MessagesEditChatAboutRequest struct {
	// The group/supergroup/channel¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Peer InputPeerClass
	// The new description
	About string
}

// MessagesEditChatAboutRequestTypeID is TL type id of MessagesEditChatAboutRequest.
const MessagesEditChatAboutRequestTypeID = 0xdef60797

func (e *MessagesEditChatAboutRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Peer == nil) {
		return false
	}
	if !(e.About == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *MessagesEditChatAboutRequest) String() string {
	if e == nil {
		return "MessagesEditChatAboutRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesEditChatAboutRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(fmt.Sprint(e.Peer))
	sb.WriteString(",\n")
	sb.WriteString("\tAbout: ")
	sb.WriteString(fmt.Sprint(e.About))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (e *MessagesEditChatAboutRequest) TypeID() uint32 {
	return MessagesEditChatAboutRequestTypeID
}

// Encode implements bin.Encoder.
func (e *MessagesEditChatAboutRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.editChatAbout#def60797 as nil")
	}
	b.PutID(MessagesEditChatAboutRequestTypeID)
	if e.Peer == nil {
		return fmt.Errorf("unable to encode messages.editChatAbout#def60797: field peer is nil")
	}
	if err := e.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.editChatAbout#def60797: field peer: %w", err)
	}
	b.PutString(e.About)
	return nil
}

// GetPeer returns value of Peer field.
func (e *MessagesEditChatAboutRequest) GetPeer() (value InputPeerClass) {
	return e.Peer
}

// GetAbout returns value of About field.
func (e *MessagesEditChatAboutRequest) GetAbout() (value string) {
	return e.About
}

// Decode implements bin.Decoder.
func (e *MessagesEditChatAboutRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.editChatAbout#def60797 to nil")
	}
	if err := b.ConsumeID(MessagesEditChatAboutRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.editChatAbout#def60797: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.editChatAbout#def60797: field peer: %w", err)
		}
		e.Peer = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.editChatAbout#def60797: field about: %w", err)
		}
		e.About = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesEditChatAboutRequest.
var (
	_ bin.Encoder = &MessagesEditChatAboutRequest{}
	_ bin.Decoder = &MessagesEditChatAboutRequest{}
)

// MessagesEditChatAbout invokes method messages.editChatAbout#def60797 returning error if any.
// Edit the description of a group/supergroup/channel¹.
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 CHAT_ABOUT_NOT_MODIFIED: About text has not changed
//  400 CHAT_ABOUT_TOO_LONG: Chat about too long
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 CHAT_NOT_MODIFIED: The pinned message wasn't modified
//  403 CHAT_WRITE_FORBIDDEN: You can't write in this chat
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/messages.editChatAbout for reference.
// Can be used by bots.
func (c *Client) MessagesEditChatAbout(ctx context.Context, request *MessagesEditChatAboutRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
