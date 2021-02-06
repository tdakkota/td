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

// ChatOnlines represents TL type `chatOnlines#f041e250`.
// Number of online users in a chat
//
// See https://core.telegram.org/constructor/chatOnlines for reference.
type ChatOnlines struct {
	// Number of online users
	Onlines int
}

// ChatOnlinesTypeID is TL type id of ChatOnlines.
const ChatOnlinesTypeID = 0xf041e250

func (c *ChatOnlines) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Onlines == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatOnlines) String() string {
	if c == nil {
		return "ChatOnlines(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChatOnlines")
	sb.WriteString("{\n")
	sb.WriteString("\tOnlines: ")
	sb.WriteString(fmt.Sprint(c.Onlines))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *ChatOnlines) TypeID() uint32 {
	return ChatOnlinesTypeID
}

// Encode implements bin.Encoder.
func (c *ChatOnlines) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatOnlines#f041e250 as nil")
	}
	b.PutID(ChatOnlinesTypeID)
	b.PutInt(c.Onlines)
	return nil
}

// GetOnlines returns value of Onlines field.
func (c *ChatOnlines) GetOnlines() (value int) {
	return c.Onlines
}

// Decode implements bin.Decoder.
func (c *ChatOnlines) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatOnlines#f041e250 to nil")
	}
	if err := b.ConsumeID(ChatOnlinesTypeID); err != nil {
		return fmt.Errorf("unable to decode chatOnlines#f041e250: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatOnlines#f041e250: field onlines: %w", err)
		}
		c.Onlines = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChatOnlines.
var (
	_ bin.Encoder = &ChatOnlines{}
	_ bin.Decoder = &ChatOnlines{}
)
