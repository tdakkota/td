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

// InlineBotSwitchPM represents TL type `inlineBotSwitchPM#3c20629f`.
// The bot requested the user to message him in private
//
// See https://core.telegram.org/constructor/inlineBotSwitchPM for reference.
type InlineBotSwitchPM struct {
	// Text for the button that switches the user to a private chat with the bot and sends the bot a start message with the parameter start_parameter (can be empty)
	Text string
	// The parameter for the /start parameter
	StartParam string
}

// InlineBotSwitchPMTypeID is TL type id of InlineBotSwitchPM.
const InlineBotSwitchPMTypeID = 0x3c20629f

func (i *InlineBotSwitchPM) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Text == "") {
		return false
	}
	if !(i.StartParam == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InlineBotSwitchPM) String() string {
	if i == nil {
		return "InlineBotSwitchPM(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InlineBotSwitchPM")
	sb.WriteString("{\n")
	sb.WriteString("\tText: ")
	sb.WriteString(fmt.Sprint(i.Text))
	sb.WriteString(",\n")
	sb.WriteString("\tStartParam: ")
	sb.WriteString(fmt.Sprint(i.StartParam))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InlineBotSwitchPM) TypeID() uint32 {
	return InlineBotSwitchPMTypeID
}

// Encode implements bin.Encoder.
func (i *InlineBotSwitchPM) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inlineBotSwitchPM#3c20629f as nil")
	}
	b.PutID(InlineBotSwitchPMTypeID)
	b.PutString(i.Text)
	b.PutString(i.StartParam)
	return nil
}

// GetText returns value of Text field.
func (i *InlineBotSwitchPM) GetText() (value string) {
	return i.Text
}

// GetStartParam returns value of StartParam field.
func (i *InlineBotSwitchPM) GetStartParam() (value string) {
	return i.StartParam
}

// Decode implements bin.Decoder.
func (i *InlineBotSwitchPM) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inlineBotSwitchPM#3c20629f to nil")
	}
	if err := b.ConsumeID(InlineBotSwitchPMTypeID); err != nil {
		return fmt.Errorf("unable to decode inlineBotSwitchPM#3c20629f: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inlineBotSwitchPM#3c20629f: field text: %w", err)
		}
		i.Text = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inlineBotSwitchPM#3c20629f: field start_param: %w", err)
		}
		i.StartParam = value
	}
	return nil
}

// Ensuring interfaces in compile-time for InlineBotSwitchPM.
var (
	_ bin.Encoder = &InlineBotSwitchPM{}
	_ bin.Decoder = &InlineBotSwitchPM{}
)
