// Code generated by gotdgen, DO NOT EDIT.

package td

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

// TestVectorInt represents TL type `testVectorInt#df9eb113`.
//
// See https://localhost:80/doc/constructor/testVectorInt for reference.
type TestVectorInt struct {
	// Vector of numbers
	Value []int32
}

// TestVectorIntTypeID is TL type id of TestVectorInt.
const TestVectorIntTypeID = 0xdf9eb113

func (t *TestVectorInt) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Value == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TestVectorInt) String() string {
	if t == nil {
		return "TestVectorInt(nil)"
	}
	var sb strings.Builder
	sb.WriteString("TestVectorInt")
	sb.WriteString("{\n")
	sb.WriteByte('[')
	for _, v := range t.Value {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (t *TestVectorInt) TypeID() uint32 {
	return TestVectorIntTypeID
}

// Encode implements bin.Encoder.
func (t *TestVectorInt) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testVectorInt#df9eb113 as nil")
	}
	b.PutID(TestVectorIntTypeID)
	b.PutVectorHeader(len(t.Value))
	for _, v := range t.Value {
		b.PutInt32(v)
	}
	return nil
}

// GetValue returns value of Value field.
func (t *TestVectorInt) GetValue() (value []int32) {
	return t.Value
}

// Decode implements bin.Decoder.
func (t *TestVectorInt) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testVectorInt#df9eb113 to nil")
	}
	if err := b.ConsumeID(TestVectorIntTypeID); err != nil {
		return fmt.Errorf("unable to decode testVectorInt#df9eb113: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode testVectorInt#df9eb113: field value: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode testVectorInt#df9eb113: field value: %w", err)
			}
			t.Value = append(t.Value, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for TestVectorInt.
var (
	_ bin.Encoder = &TestVectorInt{}
	_ bin.Decoder = &TestVectorInt{}
)
