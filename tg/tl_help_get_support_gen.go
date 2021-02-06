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

// HelpGetSupportRequest represents TL type `help.getSupport#9cdf08cd`.
// Returns the support user for the 'ask a question' feature.
//
// See https://core.telegram.org/method/help.getSupport for reference.
type HelpGetSupportRequest struct {
}

// HelpGetSupportRequestTypeID is TL type id of HelpGetSupportRequest.
const HelpGetSupportRequestTypeID = 0x9cdf08cd

func (g *HelpGetSupportRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *HelpGetSupportRequest) String() string {
	if g == nil {
		return "HelpGetSupportRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("HelpGetSupportRequest")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// TypeID returns MTProto type id (CRC code).
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (g *HelpGetSupportRequest) TypeID() uint32 {
	return HelpGetSupportRequestTypeID
}

// Encode implements bin.Encoder.
func (g *HelpGetSupportRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getSupport#9cdf08cd as nil")
	}
	b.PutID(HelpGetSupportRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetSupportRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getSupport#9cdf08cd to nil")
	}
	if err := b.ConsumeID(HelpGetSupportRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getSupport#9cdf08cd: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpGetSupportRequest.
var (
	_ bin.Encoder = &HelpGetSupportRequest{}
	_ bin.Decoder = &HelpGetSupportRequest{}
)

// HelpGetSupport invokes method help.getSupport#9cdf08cd returning error if any.
// Returns the support user for the 'ask a question' feature.
//
// See https://core.telegram.org/method/help.getSupport for reference.
func (c *Client) HelpGetSupport(ctx context.Context) (*HelpSupport, error) {
	var result HelpSupport

	request := &HelpGetSupportRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
