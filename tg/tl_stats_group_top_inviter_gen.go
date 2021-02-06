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

// StatsGroupTopInviter represents TL type `statsGroupTopInviter#31962a4c`.
// Information about an active supergroup inviter
//
// See https://core.telegram.org/constructor/statsGroupTopInviter for reference.
type StatsGroupTopInviter struct {
	// User ID
	UserID int
	// Number of invitations for statistics¹ period in consideration
	//
	// Links:
	//  1) https://core.telegram.org/api/stats
	Invitations int
}

// StatsGroupTopInviterTypeID is TL type id of StatsGroupTopInviter.
const StatsGroupTopInviterTypeID = 0x31962a4c

func (s *StatsGroupTopInviter) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.UserID == 0) {
		return false
	}
	if !(s.Invitations == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StatsGroupTopInviter) String() string {
	if s == nil {
		return "StatsGroupTopInviter(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StatsGroupTopInviter")
	sb.WriteString("{\n")
	sb.WriteString("\tUserID: ")
	sb.WriteString(fmt.Sprint(s.UserID))
	sb.WriteString(",\n")
	sb.WriteString("\tInvitations: ")
	sb.WriteString(fmt.Sprint(s.Invitations))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (s *StatsGroupTopInviter) TypeID() uint32 {
	return StatsGroupTopInviterTypeID
}

// Encode implements bin.Encoder.
func (s *StatsGroupTopInviter) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsGroupTopInviter#31962a4c as nil")
	}
	b.PutID(StatsGroupTopInviterTypeID)
	b.PutInt(s.UserID)
	b.PutInt(s.Invitations)
	return nil
}

// GetUserID returns value of UserID field.
func (s *StatsGroupTopInviter) GetUserID() (value int) {
	return s.UserID
}

// GetInvitations returns value of Invitations field.
func (s *StatsGroupTopInviter) GetInvitations() (value int) {
	return s.Invitations
}

// Decode implements bin.Decoder.
func (s *StatsGroupTopInviter) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsGroupTopInviter#31962a4c to nil")
	}
	if err := b.ConsumeID(StatsGroupTopInviterTypeID); err != nil {
		return fmt.Errorf("unable to decode statsGroupTopInviter#31962a4c: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopInviter#31962a4c: field user_id: %w", err)
		}
		s.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopInviter#31962a4c: field invitations: %w", err)
		}
		s.Invitations = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StatsGroupTopInviter.
var (
	_ bin.Encoder = &StatsGroupTopInviter{}
	_ bin.Decoder = &StatsGroupTopInviter{}
)
