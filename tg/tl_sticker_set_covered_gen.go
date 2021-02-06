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

// StickerSetCovered represents TL type `stickerSetCovered#6410a5d2`.
// Stickerset, with a specific sticker as preview
//
// See https://core.telegram.org/constructor/stickerSetCovered for reference.
type StickerSetCovered struct {
	// Stickerset
	Set StickerSet
	// Preview
	Cover DocumentClass
}

// StickerSetCoveredTypeID is TL type id of StickerSetCovered.
const StickerSetCoveredTypeID = 0x6410a5d2

func (s *StickerSetCovered) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Set.Zero()) {
		return false
	}
	if !(s.Cover == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StickerSetCovered) String() string {
	if s == nil {
		return "StickerSetCovered(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StickerSetCovered")
	sb.WriteString("{\n")
	sb.WriteString("\tSet: ")
	sb.WriteString(fmt.Sprint(s.Set))
	sb.WriteString(",\n")
	sb.WriteString("\tCover: ")
	sb.WriteString(fmt.Sprint(s.Cover))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *StickerSetCovered) TypeID() uint32 {
	return StickerSetCoveredTypeID
}

// Encode implements bin.Encoder.
func (s *StickerSetCovered) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickerSetCovered#6410a5d2 as nil")
	}
	b.PutID(StickerSetCoveredTypeID)
	if err := s.Set.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickerSetCovered#6410a5d2: field set: %w", err)
	}
	if s.Cover == nil {
		return fmt.Errorf("unable to encode stickerSetCovered#6410a5d2: field cover is nil")
	}
	if err := s.Cover.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickerSetCovered#6410a5d2: field cover: %w", err)
	}
	return nil
}

// GetSet returns value of Set field.
func (s *StickerSetCovered) GetSet() (value StickerSet) {
	return s.Set
}

// GetCover returns value of Cover field.
func (s *StickerSetCovered) GetCover() (value DocumentClass) {
	return s.Cover
}

// Decode implements bin.Decoder.
func (s *StickerSetCovered) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickerSetCovered#6410a5d2 to nil")
	}
	if err := b.ConsumeID(StickerSetCoveredTypeID); err != nil {
		return fmt.Errorf("unable to decode stickerSetCovered#6410a5d2: %w", err)
	}
	{
		if err := s.Set.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stickerSetCovered#6410a5d2: field set: %w", err)
		}
	}
	{
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode stickerSetCovered#6410a5d2: field cover: %w", err)
		}
		s.Cover = value
	}
	return nil
}

// construct implements constructor of StickerSetCoveredClass.
func (s StickerSetCovered) construct() StickerSetCoveredClass { return &s }

// Ensuring interfaces in compile-time for StickerSetCovered.
var (
	_ bin.Encoder = &StickerSetCovered{}
	_ bin.Decoder = &StickerSetCovered{}

	_ StickerSetCoveredClass = &StickerSetCovered{}
)

// StickerSetMultiCovered represents TL type `stickerSetMultiCovered#3407e51b`.
// Stickerset, with a specific stickers as preview
//
// See https://core.telegram.org/constructor/stickerSetMultiCovered for reference.
type StickerSetMultiCovered struct {
	// Stickerset
	Set StickerSet
	// Preview stickers
	Covers []DocumentClass
}

// StickerSetMultiCoveredTypeID is TL type id of StickerSetMultiCovered.
const StickerSetMultiCoveredTypeID = 0x3407e51b

func (s *StickerSetMultiCovered) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Set.Zero()) {
		return false
	}
	if !(s.Covers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StickerSetMultiCovered) String() string {
	if s == nil {
		return "StickerSetMultiCovered(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StickerSetMultiCovered")
	sb.WriteString("{\n")
	sb.WriteString("\tSet: ")
	sb.WriteString(fmt.Sprint(s.Set))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range s.Covers {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *StickerSetMultiCovered) TypeID() uint32 {
	return StickerSetMultiCoveredTypeID
}

// Encode implements bin.Encoder.
func (s *StickerSetMultiCovered) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickerSetMultiCovered#3407e51b as nil")
	}
	b.PutID(StickerSetMultiCoveredTypeID)
	if err := s.Set.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickerSetMultiCovered#3407e51b: field set: %w", err)
	}
	b.PutVectorHeader(len(s.Covers))
	for idx, v := range s.Covers {
		if v == nil {
			return fmt.Errorf("unable to encode stickerSetMultiCovered#3407e51b: field covers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stickerSetMultiCovered#3407e51b: field covers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetSet returns value of Set field.
func (s *StickerSetMultiCovered) GetSet() (value StickerSet) {
	return s.Set
}

// GetCovers returns value of Covers field.
func (s *StickerSetMultiCovered) GetCovers() (value []DocumentClass) {
	return s.Covers
}

// Decode implements bin.Decoder.
func (s *StickerSetMultiCovered) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickerSetMultiCovered#3407e51b to nil")
	}
	if err := b.ConsumeID(StickerSetMultiCoveredTypeID); err != nil {
		return fmt.Errorf("unable to decode stickerSetMultiCovered#3407e51b: %w", err)
	}
	{
		if err := s.Set.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stickerSetMultiCovered#3407e51b: field set: %w", err)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSetMultiCovered#3407e51b: field covers: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocument(b)
			if err != nil {
				return fmt.Errorf("unable to decode stickerSetMultiCovered#3407e51b: field covers: %w", err)
			}
			s.Covers = append(s.Covers, value)
		}
	}
	return nil
}

// construct implements constructor of StickerSetCoveredClass.
func (s StickerSetMultiCovered) construct() StickerSetCoveredClass { return &s }

// Ensuring interfaces in compile-time for StickerSetMultiCovered.
var (
	_ bin.Encoder = &StickerSetMultiCovered{}
	_ bin.Decoder = &StickerSetMultiCovered{}

	_ StickerSetCoveredClass = &StickerSetMultiCovered{}
)

// StickerSetCoveredClass represents StickerSetCovered generic type.
//
// See https://core.telegram.org/type/StickerSetCovered for reference.
//
// Example:
//  g, err := DecodeStickerSetCovered(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *StickerSetCovered: // stickerSetCovered#6410a5d2
//  case *StickerSetMultiCovered: // stickerSetMultiCovered#3407e51b
//  default: panic(v)
//  }
type StickerSetCoveredClass interface {
	bin.Encoder
	bin.Decoder
	construct() StickerSetCoveredClass

	// Stickerset
	GetSet() (value StickerSet)

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeStickerSetCovered implements binary de-serialization for StickerSetCoveredClass.
func DecodeStickerSetCovered(buf *bin.Buffer) (StickerSetCoveredClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case StickerSetCoveredTypeID:
		// Decoding stickerSetCovered#6410a5d2.
		v := StickerSetCovered{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StickerSetCoveredClass: %w", err)
		}
		return &v, nil
	case StickerSetMultiCoveredTypeID:
		// Decoding stickerSetMultiCovered#3407e51b.
		v := StickerSetMultiCovered{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StickerSetCoveredClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StickerSetCoveredClass: %w", bin.NewUnexpectedID(id))
	}
}

// StickerSetCovered boxes the StickerSetCoveredClass providing a helper.
type StickerSetCoveredBox struct {
	StickerSetCovered StickerSetCoveredClass
}

// Decode implements bin.Decoder for StickerSetCoveredBox.
func (b *StickerSetCoveredBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode StickerSetCoveredBox to nil")
	}
	v, err := DecodeStickerSetCovered(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StickerSetCovered = v
	return nil
}

// Encode implements bin.Encode for StickerSetCoveredBox.
func (b *StickerSetCoveredBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.StickerSetCovered == nil {
		return fmt.Errorf("unable to encode StickerSetCoveredClass as nil")
	}
	return b.StickerSetCovered.Encode(buf)
}
