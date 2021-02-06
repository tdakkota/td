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

// MessagesSetEncryptedTypingRequest represents TL type `messages.setEncryptedTyping#791451ed`.
// Send typing event by the current user to a secret chat.
//
// See https://core.telegram.org/method/messages.setEncryptedTyping for reference.
type MessagesSetEncryptedTypingRequest struct {
	// Secret chat ID
	Peer InputEncryptedChat
	// Typing.Possible values:(boolTrue)¹, if the user started typing and more than 5 seconds have passed since the last request(boolFalse)², if the user stopped typing
	//
	// Links:
	//  1) https://core.telegram.org/constructor/boolTrue
	//  2) https://core.telegram.org/constructor/boolFalse
	Typing bool
}

// MessagesSetEncryptedTypingRequestTypeID is TL type id of MessagesSetEncryptedTypingRequest.
const MessagesSetEncryptedTypingRequestTypeID = 0x791451ed

func (s *MessagesSetEncryptedTypingRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Peer.Zero()) {
		return false
	}
	if !(s.Typing == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSetEncryptedTypingRequest) String() string {
	if s == nil {
		return "MessagesSetEncryptedTypingRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesSetEncryptedTypingRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(fmt.Sprint(s.Peer))
	sb.WriteString(",\n")
	sb.WriteString("\tTyping: ")
	sb.WriteString(fmt.Sprint(s.Typing))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *MessagesSetEncryptedTypingRequest) TypeID() uint32 {
	return MessagesSetEncryptedTypingRequestTypeID
}

// Encode implements bin.Encoder.
func (s *MessagesSetEncryptedTypingRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setEncryptedTyping#791451ed as nil")
	}
	b.PutID(MessagesSetEncryptedTypingRequestTypeID)
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setEncryptedTyping#791451ed: field peer: %w", err)
	}
	b.PutBool(s.Typing)
	return nil
}

// GetPeer returns value of Peer field.
func (s *MessagesSetEncryptedTypingRequest) GetPeer() (value InputEncryptedChat) {
	return s.Peer
}

// GetTyping returns value of Typing field.
func (s *MessagesSetEncryptedTypingRequest) GetTyping() (value bool) {
	return s.Typing
}

// Decode implements bin.Decoder.
func (s *MessagesSetEncryptedTypingRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setEncryptedTyping#791451ed to nil")
	}
	if err := b.ConsumeID(MessagesSetEncryptedTypingRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.setEncryptedTyping#791451ed: %w", err)
	}
	{
		if err := s.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.setEncryptedTyping#791451ed: field peer: %w", err)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setEncryptedTyping#791451ed: field typing: %w", err)
		}
		s.Typing = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSetEncryptedTypingRequest.
var (
	_ bin.Encoder = &MessagesSetEncryptedTypingRequest{}
	_ bin.Decoder = &MessagesSetEncryptedTypingRequest{}
)

// MessagesSetEncryptedTyping invokes method messages.setEncryptedTyping#791451ed returning error if any.
// Send typing event by the current user to a secret chat.
//
// Possible errors:
//  400 CHAT_ID_INVALID: The provided chat id is invalid
//
// See https://core.telegram.org/method/messages.setEncryptedTyping for reference.
func (c *Client) MessagesSetEncryptedTyping(ctx context.Context, request *MessagesSetEncryptedTypingRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
