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

// InputPhoneContact represents TL type `inputPhoneContact#f392b7f4`.
// Phone contact. The client_id is just an arbitrary contact ID: it should be set, for example, to an incremental number when using contacts.importContacts¹, in order to retry importing only the contacts that weren't imported successfully.
//
// Links:
//  1) https://core.telegram.org/method/contacts.importContacts
//
// See https://core.telegram.org/constructor/inputPhoneContact for reference.
type InputPhoneContact struct {
	// User identifier on the client
	ClientID int64
	// Phone number
	Phone string
	// Contact's first name
	FirstName string
	// Contact's last name
	LastName string
}

// InputPhoneContactTypeID is TL type id of InputPhoneContact.
const InputPhoneContactTypeID = 0xf392b7f4

func (i *InputPhoneContact) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ClientID == 0) {
		return false
	}
	if !(i.Phone == "") {
		return false
	}
	if !(i.FirstName == "") {
		return false
	}
	if !(i.LastName == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPhoneContact) String() string {
	if i == nil {
		return "InputPhoneContact(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputPhoneContact")
	sb.WriteString("{\n")
	sb.WriteString("\tClientID: ")
	sb.WriteString(fmt.Sprint(i.ClientID))
	sb.WriteString(",\n")
	sb.WriteString("\tPhone: ")
	sb.WriteString(fmt.Sprint(i.Phone))
	sb.WriteString(",\n")
	sb.WriteString("\tFirstName: ")
	sb.WriteString(fmt.Sprint(i.FirstName))
	sb.WriteString(",\n")
	sb.WriteString("\tLastName: ")
	sb.WriteString(fmt.Sprint(i.LastName))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (i *InputPhoneContact) TypeID() uint32 {
	return InputPhoneContactTypeID
}

// Encode implements bin.Encoder.
func (i *InputPhoneContact) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPhoneContact#f392b7f4 as nil")
	}
	b.PutID(InputPhoneContactTypeID)
	b.PutLong(i.ClientID)
	b.PutString(i.Phone)
	b.PutString(i.FirstName)
	b.PutString(i.LastName)
	return nil
}

// GetClientID returns value of ClientID field.
func (i *InputPhoneContact) GetClientID() (value int64) {
	return i.ClientID
}

// GetPhone returns value of Phone field.
func (i *InputPhoneContact) GetPhone() (value string) {
	return i.Phone
}

// GetFirstName returns value of FirstName field.
func (i *InputPhoneContact) GetFirstName() (value string) {
	return i.FirstName
}

// GetLastName returns value of LastName field.
func (i *InputPhoneContact) GetLastName() (value string) {
	return i.LastName
}

// Decode implements bin.Decoder.
func (i *InputPhoneContact) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPhoneContact#f392b7f4 to nil")
	}
	if err := b.ConsumeID(InputPhoneContactTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPhoneContact#f392b7f4: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhoneContact#f392b7f4: field client_id: %w", err)
		}
		i.ClientID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhoneContact#f392b7f4: field phone: %w", err)
		}
		i.Phone = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhoneContact#f392b7f4: field first_name: %w", err)
		}
		i.FirstName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhoneContact#f392b7f4: field last_name: %w", err)
		}
		i.LastName = value
	}
	return nil
}

// Ensuring interfaces in compile-time for InputPhoneContact.
var (
	_ bin.Encoder = &InputPhoneContact{}
	_ bin.Decoder = &InputPhoneContact{}
)
