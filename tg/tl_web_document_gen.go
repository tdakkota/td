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

// WebDocument represents TL type `webDocument#1c570ed1`.
// Remote document
//
// See https://core.telegram.org/constructor/webDocument for reference.
type WebDocument struct {
	// Document URL
	URL string
	// Access hash
	AccessHash int64
	// File size
	Size int
	// MIME type
	MimeType string
	// Attributes for media types
	Attributes []DocumentAttributeClass
}

// WebDocumentTypeID is TL type id of WebDocument.
const WebDocumentTypeID = 0x1c570ed1

func (w *WebDocument) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.URL == "") {
		return false
	}
	if !(w.AccessHash == 0) {
		return false
	}
	if !(w.Size == 0) {
		return false
	}
	if !(w.MimeType == "") {
		return false
	}
	if !(w.Attributes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *WebDocument) String() string {
	if w == nil {
		return "WebDocument(nil)"
	}
	var sb strings.Builder
	sb.WriteString("WebDocument")
	sb.WriteString("{\n")
	sb.WriteString("\tURL: ")
	sb.WriteString(fmt.Sprint(w.URL))
	sb.WriteString(",\n")
	sb.WriteString("\tAccessHash: ")
	sb.WriteString(fmt.Sprint(w.AccessHash))
	sb.WriteString(",\n")
	sb.WriteString("\tSize: ")
	sb.WriteString(fmt.Sprint(w.Size))
	sb.WriteString(",\n")
	sb.WriteString("\tMimeType: ")
	sb.WriteString(fmt.Sprint(w.MimeType))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range w.Attributes {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (w *WebDocument) TypeID() uint32 {
	return WebDocumentTypeID
}

// Encode implements bin.Encoder.
func (w *WebDocument) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webDocument#1c570ed1 as nil")
	}
	b.PutID(WebDocumentTypeID)
	b.PutString(w.URL)
	b.PutLong(w.AccessHash)
	b.PutInt(w.Size)
	b.PutString(w.MimeType)
	b.PutVectorHeader(len(w.Attributes))
	for idx, v := range w.Attributes {
		if v == nil {
			return fmt.Errorf("unable to encode webDocument#1c570ed1: field attributes element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode webDocument#1c570ed1: field attributes element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetURL returns value of URL field.
func (w *WebDocument) GetURL() (value string) {
	return w.URL
}

// GetAccessHash returns value of AccessHash field.
func (w *WebDocument) GetAccessHash() (value int64) {
	return w.AccessHash
}

// GetSize returns value of Size field.
func (w *WebDocument) GetSize() (value int) {
	return w.Size
}

// GetMimeType returns value of MimeType field.
func (w *WebDocument) GetMimeType() (value string) {
	return w.MimeType
}

// GetAttributes returns value of Attributes field.
func (w *WebDocument) GetAttributes() (value []DocumentAttributeClass) {
	return w.Attributes
}

// Decode implements bin.Decoder.
func (w *WebDocument) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webDocument#1c570ed1 to nil")
	}
	if err := b.ConsumeID(WebDocumentTypeID); err != nil {
		return fmt.Errorf("unable to decode webDocument#1c570ed1: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webDocument#1c570ed1: field url: %w", err)
		}
		w.URL = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode webDocument#1c570ed1: field access_hash: %w", err)
		}
		w.AccessHash = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode webDocument#1c570ed1: field size: %w", err)
		}
		w.Size = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webDocument#1c570ed1: field mime_type: %w", err)
		}
		w.MimeType = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode webDocument#1c570ed1: field attributes: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocumentAttribute(b)
			if err != nil {
				return fmt.Errorf("unable to decode webDocument#1c570ed1: field attributes: %w", err)
			}
			w.Attributes = append(w.Attributes, value)
		}
	}
	return nil
}

// construct implements constructor of WebDocumentClass.
func (w WebDocument) construct() WebDocumentClass { return &w }

// Ensuring interfaces in compile-time for WebDocument.
var (
	_ bin.Encoder = &WebDocument{}
	_ bin.Decoder = &WebDocument{}

	_ WebDocumentClass = &WebDocument{}
)

// WebDocumentNoProxy represents TL type `webDocumentNoProxy#f9c8bcc6`.
// Remote document that can be downloaded without proxying through telegram¹
//
// Links:
//  1) https://core.telegram.org/api/files
//
// See https://core.telegram.org/constructor/webDocumentNoProxy for reference.
type WebDocumentNoProxy struct {
	// Document URL
	URL string
	// File size
	Size int
	// MIME type
	MimeType string
	// Attributes for media types
	Attributes []DocumentAttributeClass
}

// WebDocumentNoProxyTypeID is TL type id of WebDocumentNoProxy.
const WebDocumentNoProxyTypeID = 0xf9c8bcc6

func (w *WebDocumentNoProxy) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.URL == "") {
		return false
	}
	if !(w.Size == 0) {
		return false
	}
	if !(w.MimeType == "") {
		return false
	}
	if !(w.Attributes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *WebDocumentNoProxy) String() string {
	if w == nil {
		return "WebDocumentNoProxy(nil)"
	}
	var sb strings.Builder
	sb.WriteString("WebDocumentNoProxy")
	sb.WriteString("{\n")
	sb.WriteString("\tURL: ")
	sb.WriteString(fmt.Sprint(w.URL))
	sb.WriteString(",\n")
	sb.WriteString("\tSize: ")
	sb.WriteString(fmt.Sprint(w.Size))
	sb.WriteString(",\n")
	sb.WriteString("\tMimeType: ")
	sb.WriteString(fmt.Sprint(w.MimeType))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range w.Attributes {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (w *WebDocumentNoProxy) TypeID() uint32 {
	return WebDocumentNoProxyTypeID
}

// Encode implements bin.Encoder.
func (w *WebDocumentNoProxy) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode webDocumentNoProxy#f9c8bcc6 as nil")
	}
	b.PutID(WebDocumentNoProxyTypeID)
	b.PutString(w.URL)
	b.PutInt(w.Size)
	b.PutString(w.MimeType)
	b.PutVectorHeader(len(w.Attributes))
	for idx, v := range w.Attributes {
		if v == nil {
			return fmt.Errorf("unable to encode webDocumentNoProxy#f9c8bcc6: field attributes element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode webDocumentNoProxy#f9c8bcc6: field attributes element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetURL returns value of URL field.
func (w *WebDocumentNoProxy) GetURL() (value string) {
	return w.URL
}

// GetSize returns value of Size field.
func (w *WebDocumentNoProxy) GetSize() (value int) {
	return w.Size
}

// GetMimeType returns value of MimeType field.
func (w *WebDocumentNoProxy) GetMimeType() (value string) {
	return w.MimeType
}

// GetAttributes returns value of Attributes field.
func (w *WebDocumentNoProxy) GetAttributes() (value []DocumentAttributeClass) {
	return w.Attributes
}

// Decode implements bin.Decoder.
func (w *WebDocumentNoProxy) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode webDocumentNoProxy#f9c8bcc6 to nil")
	}
	if err := b.ConsumeID(WebDocumentNoProxyTypeID); err != nil {
		return fmt.Errorf("unable to decode webDocumentNoProxy#f9c8bcc6: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webDocumentNoProxy#f9c8bcc6: field url: %w", err)
		}
		w.URL = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode webDocumentNoProxy#f9c8bcc6: field size: %w", err)
		}
		w.Size = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode webDocumentNoProxy#f9c8bcc6: field mime_type: %w", err)
		}
		w.MimeType = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode webDocumentNoProxy#f9c8bcc6: field attributes: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocumentAttribute(b)
			if err != nil {
				return fmt.Errorf("unable to decode webDocumentNoProxy#f9c8bcc6: field attributes: %w", err)
			}
			w.Attributes = append(w.Attributes, value)
		}
	}
	return nil
}

// construct implements constructor of WebDocumentClass.
func (w WebDocumentNoProxy) construct() WebDocumentClass { return &w }

// Ensuring interfaces in compile-time for WebDocumentNoProxy.
var (
	_ bin.Encoder = &WebDocumentNoProxy{}
	_ bin.Decoder = &WebDocumentNoProxy{}

	_ WebDocumentClass = &WebDocumentNoProxy{}
)

// WebDocumentClass represents WebDocument generic type.
//
// See https://core.telegram.org/type/WebDocument for reference.
//
// Example:
//  g, err := DecodeWebDocument(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *WebDocument: // webDocument#1c570ed1
//  case *WebDocumentNoProxy: // webDocumentNoProxy#f9c8bcc6
//  default: panic(v)
//  }
type WebDocumentClass interface {
	bin.Encoder
	bin.Decoder
	construct() WebDocumentClass

	// Document URL
	GetURL() (value string)
	// File size
	GetSize() (value int)
	// MIME type
	GetMimeType() (value string)
	// Attributes for media types
	GetAttributes() (value []DocumentAttributeClass)

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeWebDocument implements binary de-serialization for WebDocumentClass.
func DecodeWebDocument(buf *bin.Buffer) (WebDocumentClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case WebDocumentTypeID:
		// Decoding webDocument#1c570ed1.
		v := WebDocument{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode WebDocumentClass: %w", err)
		}
		return &v, nil
	case WebDocumentNoProxyTypeID:
		// Decoding webDocumentNoProxy#f9c8bcc6.
		v := WebDocumentNoProxy{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode WebDocumentClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode WebDocumentClass: %w", bin.NewUnexpectedID(id))
	}
}

// WebDocument boxes the WebDocumentClass providing a helper.
type WebDocumentBox struct {
	WebDocument WebDocumentClass
}

// Decode implements bin.Decoder for WebDocumentBox.
func (b *WebDocumentBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode WebDocumentBox to nil")
	}
	v, err := DecodeWebDocument(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.WebDocument = v
	return nil
}

// Encode implements bin.Encode for WebDocumentBox.
func (b *WebDocumentBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.WebDocument == nil {
		return fmt.Errorf("unable to encode WebDocumentClass as nil")
	}
	return b.WebDocument.Encode(buf)
}
