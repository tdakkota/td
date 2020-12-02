// Code generated by gotdgen, DO NOT EDIT.

package td

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// DoAuthRequest represents TL type `doAuth#fd2f6687`.
type DoAuthRequest struct {
}

// DoAuthRequestTypeID is TL type id of DoAuthRequest.
const DoAuthRequestTypeID = 0xfd2f6687

// Encode implements bin.Encoder.
func (d *DoAuthRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode doAuth#fd2f6687 as nil")
	}
	b.PutID(DoAuthRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (d *DoAuthRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode doAuth#fd2f6687 to nil")
	}
	if err := b.ConsumeID(DoAuthRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode doAuth#fd2f6687: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for DoAuthRequest.
var (
	_ bin.Encoder = &DoAuthRequest{}
	_ bin.Decoder = &DoAuthRequest{}
)

// DoAuth invokes method doAuth#fd2f6687 returning error if any.
func (c *Client) DoAuth(ctx context.Context, request *DoAuthRequest) (AuthClass, error) {
	var result AuthBox
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Auth, nil
}