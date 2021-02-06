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

// UpdatesChannelDifferenceEmpty represents TL type `updates.channelDifferenceEmpty#3e11affb`.
// There are no new updates
//
// See https://core.telegram.org/constructor/updates.channelDifferenceEmpty for reference.
type UpdatesChannelDifferenceEmpty struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether there are more updates that must be fetched (always false)
	Final bool
	// The latest PTS¹
	//
	// Links:
	//  1) https://core.telegram.org/api/updates
	Pts int
	// Clients are supposed to refetch the channel difference after timeout seconds have elapsed
	//
	// Use SetTimeout and GetTimeout helpers.
	Timeout int
}

// UpdatesChannelDifferenceEmptyTypeID is TL type id of UpdatesChannelDifferenceEmpty.
const UpdatesChannelDifferenceEmptyTypeID = 0x3e11affb

func (c *UpdatesChannelDifferenceEmpty) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Final == false) {
		return false
	}
	if !(c.Pts == 0) {
		return false
	}
	if !(c.Timeout == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *UpdatesChannelDifferenceEmpty) String() string {
	if c == nil {
		return "UpdatesChannelDifferenceEmpty(nil)"
	}
	var sb strings.Builder
	sb.WriteString("UpdatesChannelDifferenceEmpty")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(c.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tPts: ")
	sb.WriteString(fmt.Sprint(c.Pts))
	sb.WriteString(",\n")
	if c.Flags.Has(1) {
		sb.WriteString("\tTimeout: ")
		sb.WriteString(fmt.Sprint(c.Timeout))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *UpdatesChannelDifferenceEmpty) TypeID() uint32 {
	return UpdatesChannelDifferenceEmptyTypeID
}

// Encode implements bin.Encoder.
func (c *UpdatesChannelDifferenceEmpty) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode updates.channelDifferenceEmpty#3e11affb as nil")
	}
	b.PutID(UpdatesChannelDifferenceEmptyTypeID)
	if !(c.Final == false) {
		c.Flags.Set(0)
	}
	if !(c.Timeout == 0) {
		c.Flags.Set(1)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.channelDifferenceEmpty#3e11affb: field flags: %w", err)
	}
	b.PutInt(c.Pts)
	if c.Flags.Has(1) {
		b.PutInt(c.Timeout)
	}
	return nil
}

// SetFinal sets value of Final conditional field.
func (c *UpdatesChannelDifferenceEmpty) SetFinal(value bool) {
	if value {
		c.Flags.Set(0)
		c.Final = true
	} else {
		c.Flags.Unset(0)
		c.Final = false
	}
}

// GetFinal returns value of Final conditional field.
func (c *UpdatesChannelDifferenceEmpty) GetFinal() (value bool) {
	return c.Flags.Has(0)
}

// GetPts returns value of Pts field.
func (c *UpdatesChannelDifferenceEmpty) GetPts() (value int) {
	return c.Pts
}

// SetTimeout sets value of Timeout conditional field.
func (c *UpdatesChannelDifferenceEmpty) SetTimeout(value int) {
	c.Flags.Set(1)
	c.Timeout = value
}

// GetTimeout returns value of Timeout conditional field and
// boolean which is true if field was set.
func (c *UpdatesChannelDifferenceEmpty) GetTimeout() (value int, ok bool) {
	if !c.Flags.Has(1) {
		return value, false
	}
	return c.Timeout, true
}

// Decode implements bin.Decoder.
func (c *UpdatesChannelDifferenceEmpty) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode updates.channelDifferenceEmpty#3e11affb to nil")
	}
	if err := b.ConsumeID(UpdatesChannelDifferenceEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.channelDifferenceEmpty#3e11affb: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceEmpty#3e11affb: field flags: %w", err)
		}
	}
	c.Final = c.Flags.Has(0)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceEmpty#3e11affb: field pts: %w", err)
		}
		c.Pts = value
	}
	if c.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceEmpty#3e11affb: field timeout: %w", err)
		}
		c.Timeout = value
	}
	return nil
}

// construct implements constructor of UpdatesChannelDifferenceClass.
func (c UpdatesChannelDifferenceEmpty) construct() UpdatesChannelDifferenceClass { return &c }

// Ensuring interfaces in compile-time for UpdatesChannelDifferenceEmpty.
var (
	_ bin.Encoder = &UpdatesChannelDifferenceEmpty{}
	_ bin.Decoder = &UpdatesChannelDifferenceEmpty{}

	_ UpdatesChannelDifferenceClass = &UpdatesChannelDifferenceEmpty{}
)

// UpdatesChannelDifferenceTooLong represents TL type `updates.channelDifferenceTooLong#a4bcc6fe`.
// The provided pts + limit < remote pts. Simply, there are too many updates to be fetched (more than limit), the client has to resolve the update gap in one of the following ways:
//
// See https://core.telegram.org/constructor/updates.channelDifferenceTooLong for reference.
type UpdatesChannelDifferenceTooLong struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether there are more updates that must be fetched (always false)
	Final bool
	// Clients are supposed to refetch the channel difference after timeout seconds have elapsed
	//
	// Use SetTimeout and GetTimeout helpers.
	Timeout int
	// Dialog containing the latest PTS¹ that can be used to reset the channel state
	//
	// Links:
	//  1) https://core.telegram.org/api/updates
	Dialog DialogClass
	// The latest messages
	Messages []MessageClass
	// Chats from messages
	Chats []ChatClass
	// Users from messages
	Users []UserClass
}

// UpdatesChannelDifferenceTooLongTypeID is TL type id of UpdatesChannelDifferenceTooLong.
const UpdatesChannelDifferenceTooLongTypeID = 0xa4bcc6fe

func (c *UpdatesChannelDifferenceTooLong) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Final == false) {
		return false
	}
	if !(c.Timeout == 0) {
		return false
	}
	if !(c.Dialog == nil) {
		return false
	}
	if !(c.Messages == nil) {
		return false
	}
	if !(c.Chats == nil) {
		return false
	}
	if !(c.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *UpdatesChannelDifferenceTooLong) String() string {
	if c == nil {
		return "UpdatesChannelDifferenceTooLong(nil)"
	}
	var sb strings.Builder
	sb.WriteString("UpdatesChannelDifferenceTooLong")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(c.Flags))
	sb.WriteString(",\n")
	if c.Flags.Has(1) {
		sb.WriteString("\tTimeout: ")
		sb.WriteString(fmt.Sprint(c.Timeout))
		sb.WriteString(",\n")
	}
	sb.WriteString("\tDialog: ")
	sb.WriteString(fmt.Sprint(c.Dialog))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range c.Messages {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range c.Chats {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range c.Users {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *UpdatesChannelDifferenceTooLong) TypeID() uint32 {
	return UpdatesChannelDifferenceTooLongTypeID
}

// Encode implements bin.Encoder.
func (c *UpdatesChannelDifferenceTooLong) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode updates.channelDifferenceTooLong#a4bcc6fe as nil")
	}
	b.PutID(UpdatesChannelDifferenceTooLongTypeID)
	if !(c.Final == false) {
		c.Flags.Set(0)
	}
	if !(c.Timeout == 0) {
		c.Flags.Set(1)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field flags: %w", err)
	}
	if c.Flags.Has(1) {
		b.PutInt(c.Timeout)
	}
	if c.Dialog == nil {
		return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field dialog is nil")
	}
	if err := c.Dialog.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field dialog: %w", err)
	}
	b.PutVectorHeader(len(c.Messages))
	for idx, v := range c.Messages {
		if v == nil {
			return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Chats))
	for idx, v := range c.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Users))
	for idx, v := range c.Users {
		if v == nil {
			return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.channelDifferenceTooLong#a4bcc6fe: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// SetFinal sets value of Final conditional field.
func (c *UpdatesChannelDifferenceTooLong) SetFinal(value bool) {
	if value {
		c.Flags.Set(0)
		c.Final = true
	} else {
		c.Flags.Unset(0)
		c.Final = false
	}
}

// GetFinal returns value of Final conditional field.
func (c *UpdatesChannelDifferenceTooLong) GetFinal() (value bool) {
	return c.Flags.Has(0)
}

// SetTimeout sets value of Timeout conditional field.
func (c *UpdatesChannelDifferenceTooLong) SetTimeout(value int) {
	c.Flags.Set(1)
	c.Timeout = value
}

// GetTimeout returns value of Timeout conditional field and
// boolean which is true if field was set.
func (c *UpdatesChannelDifferenceTooLong) GetTimeout() (value int, ok bool) {
	if !c.Flags.Has(1) {
		return value, false
	}
	return c.Timeout, true
}

// GetDialog returns value of Dialog field.
func (c *UpdatesChannelDifferenceTooLong) GetDialog() (value DialogClass) {
	return c.Dialog
}

// GetMessages returns value of Messages field.
func (c *UpdatesChannelDifferenceTooLong) GetMessages() (value []MessageClass) {
	return c.Messages
}

// GetChats returns value of Chats field.
func (c *UpdatesChannelDifferenceTooLong) GetChats() (value []ChatClass) {
	return c.Chats
}

// GetUsers returns value of Users field.
func (c *UpdatesChannelDifferenceTooLong) GetUsers() (value []UserClass) {
	return c.Users
}

// Decode implements bin.Decoder.
func (c *UpdatesChannelDifferenceTooLong) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode updates.channelDifferenceTooLong#a4bcc6fe to nil")
	}
	if err := b.ConsumeID(UpdatesChannelDifferenceTooLongTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field flags: %w", err)
		}
	}
	c.Final = c.Flags.Has(0)
	if c.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field timeout: %w", err)
		}
		c.Timeout = value
	}
	{
		value, err := DecodeDialog(b)
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field dialog: %w", err)
		}
		c.Dialog = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field messages: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field messages: %w", err)
			}
			c.Messages = append(c.Messages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field chats: %w", err)
			}
			c.Chats = append(c.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.channelDifferenceTooLong#a4bcc6fe: field users: %w", err)
			}
			c.Users = append(c.Users, value)
		}
	}
	return nil
}

// construct implements constructor of UpdatesChannelDifferenceClass.
func (c UpdatesChannelDifferenceTooLong) construct() UpdatesChannelDifferenceClass { return &c }

// Ensuring interfaces in compile-time for UpdatesChannelDifferenceTooLong.
var (
	_ bin.Encoder = &UpdatesChannelDifferenceTooLong{}
	_ bin.Decoder = &UpdatesChannelDifferenceTooLong{}

	_ UpdatesChannelDifferenceClass = &UpdatesChannelDifferenceTooLong{}
)

// UpdatesChannelDifference represents TL type `updates.channelDifference#2064674e`.
// The new updates
//
// See https://core.telegram.org/constructor/updates.channelDifference for reference.
type UpdatesChannelDifference struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether there are more updates to be fetched using getDifference, starting from the provided pts
	Final bool
	// The PTS¹ from which to start getting updates the next time
	//
	// Links:
	//  1) https://core.telegram.org/api/updates
	Pts int
	// Clients are supposed to refetch the channel difference after timeout seconds have elapsed
	//
	// Use SetTimeout and GetTimeout helpers.
	Timeout int
	// New messages
	NewMessages []MessageClass
	// Other updates
	OtherUpdates []UpdateClass
	// Chats
	Chats []ChatClass
	// Users
	Users []UserClass
}

// UpdatesChannelDifferenceTypeID is TL type id of UpdatesChannelDifference.
const UpdatesChannelDifferenceTypeID = 0x2064674e

func (c *UpdatesChannelDifference) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Final == false) {
		return false
	}
	if !(c.Pts == 0) {
		return false
	}
	if !(c.Timeout == 0) {
		return false
	}
	if !(c.NewMessages == nil) {
		return false
	}
	if !(c.OtherUpdates == nil) {
		return false
	}
	if !(c.Chats == nil) {
		return false
	}
	if !(c.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *UpdatesChannelDifference) String() string {
	if c == nil {
		return "UpdatesChannelDifference(nil)"
	}
	var sb strings.Builder
	sb.WriteString("UpdatesChannelDifference")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(c.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tPts: ")
	sb.WriteString(fmt.Sprint(c.Pts))
	sb.WriteString(",\n")
	if c.Flags.Has(1) {
		sb.WriteString("\tTimeout: ")
		sb.WriteString(fmt.Sprint(c.Timeout))
		sb.WriteString(",\n")
	}
	sb.WriteByte('[')
	for _, v := range c.NewMessages {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range c.OtherUpdates {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range c.Chats {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range c.Users {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (c *UpdatesChannelDifference) TypeID() uint32 {
	return UpdatesChannelDifferenceTypeID
}

// Encode implements bin.Encoder.
func (c *UpdatesChannelDifference) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode updates.channelDifference#2064674e as nil")
	}
	b.PutID(UpdatesChannelDifferenceTypeID)
	if !(c.Final == false) {
		c.Flags.Set(0)
	}
	if !(c.Timeout == 0) {
		c.Flags.Set(1)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field flags: %w", err)
	}
	b.PutInt(c.Pts)
	if c.Flags.Has(1) {
		b.PutInt(c.Timeout)
	}
	b.PutVectorHeader(len(c.NewMessages))
	for idx, v := range c.NewMessages {
		if v == nil {
			return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field new_messages element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field new_messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.OtherUpdates))
	for idx, v := range c.OtherUpdates {
		if v == nil {
			return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field other_updates element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field other_updates element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Chats))
	for idx, v := range c.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Users))
	for idx, v := range c.Users {
		if v == nil {
			return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode updates.channelDifference#2064674e: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// SetFinal sets value of Final conditional field.
func (c *UpdatesChannelDifference) SetFinal(value bool) {
	if value {
		c.Flags.Set(0)
		c.Final = true
	} else {
		c.Flags.Unset(0)
		c.Final = false
	}
}

// GetFinal returns value of Final conditional field.
func (c *UpdatesChannelDifference) GetFinal() (value bool) {
	return c.Flags.Has(0)
}

// GetPts returns value of Pts field.
func (c *UpdatesChannelDifference) GetPts() (value int) {
	return c.Pts
}

// SetTimeout sets value of Timeout conditional field.
func (c *UpdatesChannelDifference) SetTimeout(value int) {
	c.Flags.Set(1)
	c.Timeout = value
}

// GetTimeout returns value of Timeout conditional field and
// boolean which is true if field was set.
func (c *UpdatesChannelDifference) GetTimeout() (value int, ok bool) {
	if !c.Flags.Has(1) {
		return value, false
	}
	return c.Timeout, true
}

// GetNewMessages returns value of NewMessages field.
func (c *UpdatesChannelDifference) GetNewMessages() (value []MessageClass) {
	return c.NewMessages
}

// GetOtherUpdates returns value of OtherUpdates field.
func (c *UpdatesChannelDifference) GetOtherUpdates() (value []UpdateClass) {
	return c.OtherUpdates
}

// GetChats returns value of Chats field.
func (c *UpdatesChannelDifference) GetChats() (value []ChatClass) {
	return c.Chats
}

// GetUsers returns value of Users field.
func (c *UpdatesChannelDifference) GetUsers() (value []UserClass) {
	return c.Users
}

// Decode implements bin.Decoder.
func (c *UpdatesChannelDifference) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode updates.channelDifference#2064674e to nil")
	}
	if err := b.ConsumeID(UpdatesChannelDifferenceTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.channelDifference#2064674e: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field flags: %w", err)
		}
	}
	c.Final = c.Flags.Has(0)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field pts: %w", err)
		}
		c.Pts = value
	}
	if c.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field timeout: %w", err)
		}
		c.Timeout = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field new_messages: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessage(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field new_messages: %w", err)
			}
			c.NewMessages = append(c.NewMessages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field other_updates: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUpdate(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field other_updates: %w", err)
			}
			c.OtherUpdates = append(c.OtherUpdates, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field chats: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field chats: %w", err)
			}
			c.Chats = append(c.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode updates.channelDifference#2064674e: field users: %w", err)
			}
			c.Users = append(c.Users, value)
		}
	}
	return nil
}

// construct implements constructor of UpdatesChannelDifferenceClass.
func (c UpdatesChannelDifference) construct() UpdatesChannelDifferenceClass { return &c }

// Ensuring interfaces in compile-time for UpdatesChannelDifference.
var (
	_ bin.Encoder = &UpdatesChannelDifference{}
	_ bin.Decoder = &UpdatesChannelDifference{}

	_ UpdatesChannelDifferenceClass = &UpdatesChannelDifference{}
)

// UpdatesChannelDifferenceClass represents updates.ChannelDifference generic type.
//
// See https://core.telegram.org/type/updates.ChannelDifference for reference.
//
// Example:
//  g, err := DecodeUpdatesChannelDifference(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *UpdatesChannelDifferenceEmpty: // updates.channelDifferenceEmpty#3e11affb
//  case *UpdatesChannelDifferenceTooLong: // updates.channelDifferenceTooLong#a4bcc6fe
//  case *UpdatesChannelDifference: // updates.channelDifference#2064674e
//  default: panic(v)
//  }
type UpdatesChannelDifferenceClass interface {
	bin.Encoder
	bin.Decoder
	construct() UpdatesChannelDifferenceClass

	// Whether there are more updates that must be fetched (always false)
	GetFinal() (value bool)
	// Clients are supposed to refetch the channel difference after timeout seconds have elapsed
	GetTimeout() (value int, ok bool)

	// TypeID returns MTProto type id (CRC code).
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeUpdatesChannelDifference implements binary de-serialization for UpdatesChannelDifferenceClass.
func DecodeUpdatesChannelDifference(buf *bin.Buffer) (UpdatesChannelDifferenceClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case UpdatesChannelDifferenceEmptyTypeID:
		// Decoding updates.channelDifferenceEmpty#3e11affb.
		v := UpdatesChannelDifferenceEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UpdatesChannelDifferenceClass: %w", err)
		}
		return &v, nil
	case UpdatesChannelDifferenceTooLongTypeID:
		// Decoding updates.channelDifferenceTooLong#a4bcc6fe.
		v := UpdatesChannelDifferenceTooLong{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UpdatesChannelDifferenceClass: %w", err)
		}
		return &v, nil
	case UpdatesChannelDifferenceTypeID:
		// Decoding updates.channelDifference#2064674e.
		v := UpdatesChannelDifference{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UpdatesChannelDifferenceClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode UpdatesChannelDifferenceClass: %w", bin.NewUnexpectedID(id))
	}
}

// UpdatesChannelDifference boxes the UpdatesChannelDifferenceClass providing a helper.
type UpdatesChannelDifferenceBox struct {
	ChannelDifference UpdatesChannelDifferenceClass
}

// Decode implements bin.Decoder for UpdatesChannelDifferenceBox.
func (b *UpdatesChannelDifferenceBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode UpdatesChannelDifferenceBox to nil")
	}
	v, err := DecodeUpdatesChannelDifference(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChannelDifference = v
	return nil
}

// Encode implements bin.Encode for UpdatesChannelDifferenceBox.
func (b *UpdatesChannelDifferenceBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChannelDifference == nil {
		return fmt.Errorf("unable to encode UpdatesChannelDifferenceClass as nil")
	}
	return b.ChannelDifference.Encode(buf)
}
