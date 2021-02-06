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

// InputFolderPeer represents TL type `inputFolderPeer#fbd2c296`.
// Peer in a folder
//
// See https://core.telegram.org/constructor/inputFolderPeer for reference.
type InputFolderPeer struct {
	// Peer
	Peer InputPeerClass
	// Peer folder ID, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders#peer-folders
	FolderID int
}

// InputFolderPeerTypeID is TL type id of InputFolderPeer.
const InputFolderPeerTypeID = 0xfbd2c296

func (i *InputFolderPeer) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Peer == nil) {
		return false
	}
	if !(i.FolderID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputFolderPeer) String() string {
	if i == nil {
		return "InputFolderPeer(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputFolderPeer")
	sb.WriteString("{\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(fmt.Sprint(i.Peer))
	sb.WriteString(",\n")
	sb.WriteString("\tFolderID: ")
	sb.WriteString(fmt.Sprint(i.FolderID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputFolderPeer) TypeID() uint32 {
	return InputFolderPeerTypeID
}

// Encode implements bin.Encoder.
func (i *InputFolderPeer) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputFolderPeer#fbd2c296 as nil")
	}
	b.PutID(InputFolderPeerTypeID)
	if i.Peer == nil {
		return fmt.Errorf("unable to encode inputFolderPeer#fbd2c296: field peer is nil")
	}
	if err := i.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputFolderPeer#fbd2c296: field peer: %w", err)
	}
	b.PutInt(i.FolderID)
	return nil
}

// GetPeer returns value of Peer field.
func (i *InputFolderPeer) GetPeer() (value InputPeerClass) {
	return i.Peer
}

// GetFolderID returns value of FolderID field.
func (i *InputFolderPeer) GetFolderID() (value int) {
	return i.FolderID
}

// Decode implements bin.Decoder.
func (i *InputFolderPeer) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputFolderPeer#fbd2c296 to nil")
	}
	if err := b.ConsumeID(InputFolderPeerTypeID); err != nil {
		return fmt.Errorf("unable to decode inputFolderPeer#fbd2c296: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputFolderPeer#fbd2c296: field peer: %w", err)
		}
		i.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputFolderPeer#fbd2c296: field folder_id: %w", err)
		}
		i.FolderID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for InputFolderPeer.
var (
	_ bin.Encoder = &InputFolderPeer{}
	_ bin.Decoder = &InputFolderPeer{}
)
