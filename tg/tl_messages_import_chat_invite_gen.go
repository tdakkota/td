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

// MessagesImportChatInviteRequest represents TL type `messages.importChatInvite#6c50051c`.
// Import a chat invite and join a private chat/supergroup/channel
//
// See https://core.telegram.org/method/messages.importChatInvite for reference.
type MessagesImportChatInviteRequest struct {
	// hash from t.me/joinchat/hash
	Hash string
}

// MessagesImportChatInviteRequestTypeID is TL type id of MessagesImportChatInviteRequest.
const MessagesImportChatInviteRequestTypeID = 0x6c50051c

func (i *MessagesImportChatInviteRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Hash == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *MessagesImportChatInviteRequest) String() string {
	if i == nil {
		return "MessagesImportChatInviteRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesImportChatInviteRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tHash: ")
	sb.WriteString(fmt.Sprint(i.Hash))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *MessagesImportChatInviteRequest) TypeID() uint32 {
	return MessagesImportChatInviteRequestTypeID
}

// Encode implements bin.Encoder.
func (i *MessagesImportChatInviteRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode messages.importChatInvite#6c50051c as nil")
	}
	b.PutID(MessagesImportChatInviteRequestTypeID)
	b.PutString(i.Hash)
	return nil
}

// GetHash returns value of Hash field.
func (i *MessagesImportChatInviteRequest) GetHash() (value string) {
	return i.Hash
}

// Decode implements bin.Decoder.
func (i *MessagesImportChatInviteRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode messages.importChatInvite#6c50051c to nil")
	}
	if err := b.ConsumeID(MessagesImportChatInviteRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.importChatInvite#6c50051c: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.importChatInvite#6c50051c: field hash: %w", err)
		}
		i.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesImportChatInviteRequest.
var (
	_ bin.Encoder = &MessagesImportChatInviteRequest{}
	_ bin.Decoder = &MessagesImportChatInviteRequest{}
)

// MessagesImportChatInvite invokes method messages.importChatInvite#6c50051c returning error if any.
// Import a chat invite and join a private chat/supergroup/channel
//
// Possible errors:
//  400 CHANNELS_TOO_MUCH: You have joined too many channels/supergroups
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 INVITE_HASH_EMPTY: The invite hash is empty
//  400 INVITE_HASH_EXPIRED: The invite link has expired
//  400 INVITE_HASH_INVALID: The invite hash is invalid
//  400 MSG_ID_INVALID: Invalid message ID provided
//  400 PEER_ID_INVALID: The provided peer id is invalid
//  400 USERS_TOO_MUCH: The maximum number of users has been exceeded (to create a chat, for example)
//  400 USER_ALREADY_PARTICIPANT: The user is already in the group
//  400 USER_CHANNELS_TOO_MUCH: One of the users you tried to add is already in too many channels/supergroups
//
// See https://core.telegram.org/method/messages.importChatInvite for reference.
func (c *Client) MessagesImportChatInvite(ctx context.Context, hash string) (UpdatesClass, error) {
	var result UpdatesBox

	request := &MessagesImportChatInviteRequest{
		Hash: hash,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
