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

// InlineQueryPeerTypeSameBotPM represents TL type `inlineQueryPeerTypeSameBotPM#3081ed9d`.
//
// See https://core.telegram.org/constructor/inlineQueryPeerTypeSameBotPM for reference.
type InlineQueryPeerTypeSameBotPM struct {
}

// InlineQueryPeerTypeSameBotPMTypeID is TL type id of InlineQueryPeerTypeSameBotPM.
const InlineQueryPeerTypeSameBotPMTypeID = 0x3081ed9d

func (i *InlineQueryPeerTypeSameBotPM) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InlineQueryPeerTypeSameBotPM) String() string {
	if i == nil {
		return "InlineQueryPeerTypeSameBotPM(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InlineQueryPeerTypeSameBotPM")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InlineQueryPeerTypeSameBotPM) TypeID() uint32 {
	return InlineQueryPeerTypeSameBotPMTypeID
}

// Encode implements bin.Encoder.
func (i *InlineQueryPeerTypeSameBotPM) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineQueryPeerTypeSameBotPM#3081ed9d as nil")
	}
	b.PutID(InlineQueryPeerTypeSameBotPMTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InlineQueryPeerTypeSameBotPM) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineQueryPeerTypeSameBotPM#3081ed9d to nil")
	}
	if err := b.ConsumeID(InlineQueryPeerTypeSameBotPMTypeID); err != nil {
		return fmt.Errorf("unable to decode inlineQueryPeerTypeSameBotPM#3081ed9d: %w", err)
	}
	return nil
}

// construct implements constructor of InlineQueryPeerTypeClass.
func (i InlineQueryPeerTypeSameBotPM) construct() InlineQueryPeerTypeClass { return &i }

// Ensuring interfaces in compile-time for InlineQueryPeerTypeSameBotPM.
var (
	_ bin.Encoder = &InlineQueryPeerTypeSameBotPM{}
	_ bin.Decoder = &InlineQueryPeerTypeSameBotPM{}

	_ InlineQueryPeerTypeClass = &InlineQueryPeerTypeSameBotPM{}
)

// InlineQueryPeerTypePM represents TL type `inlineQueryPeerTypePM#833c0fac`.
//
// See https://core.telegram.org/constructor/inlineQueryPeerTypePM for reference.
type InlineQueryPeerTypePM struct {
}

// InlineQueryPeerTypePMTypeID is TL type id of InlineQueryPeerTypePM.
const InlineQueryPeerTypePMTypeID = 0x833c0fac

func (i *InlineQueryPeerTypePM) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InlineQueryPeerTypePM) String() string {
	if i == nil {
		return "InlineQueryPeerTypePM(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InlineQueryPeerTypePM")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InlineQueryPeerTypePM) TypeID() uint32 {
	return InlineQueryPeerTypePMTypeID
}

// Encode implements bin.Encoder.
func (i *InlineQueryPeerTypePM) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineQueryPeerTypePM#833c0fac as nil")
	}
	b.PutID(InlineQueryPeerTypePMTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InlineQueryPeerTypePM) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineQueryPeerTypePM#833c0fac to nil")
	}
	if err := b.ConsumeID(InlineQueryPeerTypePMTypeID); err != nil {
		return fmt.Errorf("unable to decode inlineQueryPeerTypePM#833c0fac: %w", err)
	}
	return nil
}

// construct implements constructor of InlineQueryPeerTypeClass.
func (i InlineQueryPeerTypePM) construct() InlineQueryPeerTypeClass { return &i }

// Ensuring interfaces in compile-time for InlineQueryPeerTypePM.
var (
	_ bin.Encoder = &InlineQueryPeerTypePM{}
	_ bin.Decoder = &InlineQueryPeerTypePM{}

	_ InlineQueryPeerTypeClass = &InlineQueryPeerTypePM{}
)

// InlineQueryPeerTypeChat represents TL type `inlineQueryPeerTypeChat#d766c50a`.
//
// See https://core.telegram.org/constructor/inlineQueryPeerTypeChat for reference.
type InlineQueryPeerTypeChat struct {
}

// InlineQueryPeerTypeChatTypeID is TL type id of InlineQueryPeerTypeChat.
const InlineQueryPeerTypeChatTypeID = 0xd766c50a

func (i *InlineQueryPeerTypeChat) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InlineQueryPeerTypeChat) String() string {
	if i == nil {
		return "InlineQueryPeerTypeChat(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InlineQueryPeerTypeChat")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InlineQueryPeerTypeChat) TypeID() uint32 {
	return InlineQueryPeerTypeChatTypeID
}

// Encode implements bin.Encoder.
func (i *InlineQueryPeerTypeChat) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineQueryPeerTypeChat#d766c50a as nil")
	}
	b.PutID(InlineQueryPeerTypeChatTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InlineQueryPeerTypeChat) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineQueryPeerTypeChat#d766c50a to nil")
	}
	if err := b.ConsumeID(InlineQueryPeerTypeChatTypeID); err != nil {
		return fmt.Errorf("unable to decode inlineQueryPeerTypeChat#d766c50a: %w", err)
	}
	return nil
}

// construct implements constructor of InlineQueryPeerTypeClass.
func (i InlineQueryPeerTypeChat) construct() InlineQueryPeerTypeClass { return &i }

// Ensuring interfaces in compile-time for InlineQueryPeerTypeChat.
var (
	_ bin.Encoder = &InlineQueryPeerTypeChat{}
	_ bin.Decoder = &InlineQueryPeerTypeChat{}

	_ InlineQueryPeerTypeClass = &InlineQueryPeerTypeChat{}
)

// InlineQueryPeerTypeMegagroup represents TL type `inlineQueryPeerTypeMegagroup#5ec4be43`.
//
// See https://core.telegram.org/constructor/inlineQueryPeerTypeMegagroup for reference.
type InlineQueryPeerTypeMegagroup struct {
}

// InlineQueryPeerTypeMegagroupTypeID is TL type id of InlineQueryPeerTypeMegagroup.
const InlineQueryPeerTypeMegagroupTypeID = 0x5ec4be43

func (i *InlineQueryPeerTypeMegagroup) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InlineQueryPeerTypeMegagroup) String() string {
	if i == nil {
		return "InlineQueryPeerTypeMegagroup(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InlineQueryPeerTypeMegagroup")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InlineQueryPeerTypeMegagroup) TypeID() uint32 {
	return InlineQueryPeerTypeMegagroupTypeID
}

// Encode implements bin.Encoder.
func (i *InlineQueryPeerTypeMegagroup) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineQueryPeerTypeMegagroup#5ec4be43 as nil")
	}
	b.PutID(InlineQueryPeerTypeMegagroupTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InlineQueryPeerTypeMegagroup) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineQueryPeerTypeMegagroup#5ec4be43 to nil")
	}
	if err := b.ConsumeID(InlineQueryPeerTypeMegagroupTypeID); err != nil {
		return fmt.Errorf("unable to decode inlineQueryPeerTypeMegagroup#5ec4be43: %w", err)
	}
	return nil
}

// construct implements constructor of InlineQueryPeerTypeClass.
func (i InlineQueryPeerTypeMegagroup) construct() InlineQueryPeerTypeClass { return &i }

// Ensuring interfaces in compile-time for InlineQueryPeerTypeMegagroup.
var (
	_ bin.Encoder = &InlineQueryPeerTypeMegagroup{}
	_ bin.Decoder = &InlineQueryPeerTypeMegagroup{}

	_ InlineQueryPeerTypeClass = &InlineQueryPeerTypeMegagroup{}
)

// InlineQueryPeerTypeBroadcast represents TL type `inlineQueryPeerTypeBroadcast#6334ee9a`.
//
// See https://core.telegram.org/constructor/inlineQueryPeerTypeBroadcast for reference.
type InlineQueryPeerTypeBroadcast struct {
}

// InlineQueryPeerTypeBroadcastTypeID is TL type id of InlineQueryPeerTypeBroadcast.
const InlineQueryPeerTypeBroadcastTypeID = 0x6334ee9a

func (i *InlineQueryPeerTypeBroadcast) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InlineQueryPeerTypeBroadcast) String() string {
	if i == nil {
		return "InlineQueryPeerTypeBroadcast(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InlineQueryPeerTypeBroadcast")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InlineQueryPeerTypeBroadcast) TypeID() uint32 {
	return InlineQueryPeerTypeBroadcastTypeID
}

// Encode implements bin.Encoder.
func (i *InlineQueryPeerTypeBroadcast) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineQueryPeerTypeBroadcast#6334ee9a as nil")
	}
	b.PutID(InlineQueryPeerTypeBroadcastTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InlineQueryPeerTypeBroadcast) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineQueryPeerTypeBroadcast#6334ee9a to nil")
	}
	if err := b.ConsumeID(InlineQueryPeerTypeBroadcastTypeID); err != nil {
		return fmt.Errorf("unable to decode inlineQueryPeerTypeBroadcast#6334ee9a: %w", err)
	}
	return nil
}

// construct implements constructor of InlineQueryPeerTypeClass.
func (i InlineQueryPeerTypeBroadcast) construct() InlineQueryPeerTypeClass { return &i }

// Ensuring interfaces in compile-time for InlineQueryPeerTypeBroadcast.
var (
	_ bin.Encoder = &InlineQueryPeerTypeBroadcast{}
	_ bin.Decoder = &InlineQueryPeerTypeBroadcast{}

	_ InlineQueryPeerTypeClass = &InlineQueryPeerTypeBroadcast{}
)

// InlineQueryPeerTypeClass represents InlineQueryPeerType generic type.
//
// See https://core.telegram.org/type/InlineQueryPeerType for reference.
//
// Example:
//  g, err := DecodeInlineQueryPeerType(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InlineQueryPeerTypeSameBotPM: // inlineQueryPeerTypeSameBotPM#3081ed9d
//  case *InlineQueryPeerTypePM: // inlineQueryPeerTypePM#833c0fac
//  case *InlineQueryPeerTypeChat: // inlineQueryPeerTypeChat#d766c50a
//  case *InlineQueryPeerTypeMegagroup: // inlineQueryPeerTypeMegagroup#5ec4be43
//  case *InlineQueryPeerTypeBroadcast: // inlineQueryPeerTypeBroadcast#6334ee9a
//  default: panic(v)
//  }
type InlineQueryPeerTypeClass interface {
	bin.Encoder
	bin.Decoder
	construct() InlineQueryPeerTypeClass

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeInlineQueryPeerType implements binary de-serialization for InlineQueryPeerTypeClass.
func DecodeInlineQueryPeerType(buf *bin.Buffer) (InlineQueryPeerTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InlineQueryPeerTypeSameBotPMTypeID:
		// Decoding inlineQueryPeerTypeSameBotPM#3081ed9d.
		v := InlineQueryPeerTypeSameBotPM{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InlineQueryPeerTypeClass: %w", err)
		}
		return &v, nil
	case InlineQueryPeerTypePMTypeID:
		// Decoding inlineQueryPeerTypePM#833c0fac.
		v := InlineQueryPeerTypePM{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InlineQueryPeerTypeClass: %w", err)
		}
		return &v, nil
	case InlineQueryPeerTypeChatTypeID:
		// Decoding inlineQueryPeerTypeChat#d766c50a.
		v := InlineQueryPeerTypeChat{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InlineQueryPeerTypeClass: %w", err)
		}
		return &v, nil
	case InlineQueryPeerTypeMegagroupTypeID:
		// Decoding inlineQueryPeerTypeMegagroup#5ec4be43.
		v := InlineQueryPeerTypeMegagroup{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InlineQueryPeerTypeClass: %w", err)
		}
		return &v, nil
	case InlineQueryPeerTypeBroadcastTypeID:
		// Decoding inlineQueryPeerTypeBroadcast#6334ee9a.
		v := InlineQueryPeerTypeBroadcast{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InlineQueryPeerTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InlineQueryPeerTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// InlineQueryPeerType boxes the InlineQueryPeerTypeClass providing a helper.
type InlineQueryPeerTypeBox struct {
	InlineQueryPeerType InlineQueryPeerTypeClass
}

// Decode implements bin.Decoder for InlineQueryPeerTypeBox.
func (b *InlineQueryPeerTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InlineQueryPeerTypeBox to nil")
	}
	v, err := DecodeInlineQueryPeerType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InlineQueryPeerType = v
	return nil
}

// Encode implements bin.Encode for InlineQueryPeerTypeBox.
func (b *InlineQueryPeerTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InlineQueryPeerType == nil {
		return fmt.Errorf("unable to encode InlineQueryPeerTypeClass as nil")
	}
	return b.InlineQueryPeerType.Encode(buf)
}
