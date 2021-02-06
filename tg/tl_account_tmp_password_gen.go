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

// AccountTmpPassword represents TL type `account.tmpPassword#db64fd34`.
// Temporary payment password
//
// See https://core.telegram.org/constructor/account.tmpPassword for reference.
type AccountTmpPassword struct {
	// Temporary password
	TmpPassword []byte
	// Validity period
	ValidUntil int
}

// AccountTmpPasswordTypeID is TL type id of AccountTmpPassword.
const AccountTmpPasswordTypeID = 0xdb64fd34

func (t *AccountTmpPassword) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.TmpPassword == nil) {
		return false
	}
	if !(t.ValidUntil == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *AccountTmpPassword) String() string {
	if t == nil {
		return "AccountTmpPassword(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountTmpPassword")
	sb.WriteString("{\n")
	sb.WriteString("\tTmpPassword: ")
	sb.WriteString(fmt.Sprint(t.TmpPassword))
	sb.WriteString(",\n")
	sb.WriteString("\tValidUntil: ")
	sb.WriteString(fmt.Sprint(t.ValidUntil))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (t *AccountTmpPassword) TypeID() uint32 {
	return AccountTmpPasswordTypeID
}

// Encode implements bin.Encoder.
func (t *AccountTmpPassword) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode account.tmpPassword#db64fd34 as nil")
	}
	b.PutID(AccountTmpPasswordTypeID)
	b.PutBytes(t.TmpPassword)
	b.PutInt(t.ValidUntil)
	return nil
}

// GetTmpPassword returns value of TmpPassword field.
func (t *AccountTmpPassword) GetTmpPassword() (value []byte) {
	return t.TmpPassword
}

// GetValidUntil returns value of ValidUntil field.
func (t *AccountTmpPassword) GetValidUntil() (value int) {
	return t.ValidUntil
}

// Decode implements bin.Decoder.
func (t *AccountTmpPassword) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode account.tmpPassword#db64fd34 to nil")
	}
	if err := b.ConsumeID(AccountTmpPasswordTypeID); err != nil {
		return fmt.Errorf("unable to decode account.tmpPassword#db64fd34: %w", err)
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode account.tmpPassword#db64fd34: field tmp_password: %w", err)
		}
		t.TmpPassword = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.tmpPassword#db64fd34: field valid_until: %w", err)
		}
		t.ValidUntil = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountTmpPassword.
var (
	_ bin.Encoder = &AccountTmpPassword{}
	_ bin.Decoder = &AccountTmpPassword{}
)
