package crypto

import (
	"crypto/aes"
	"io"

	"github.com/gotd/ige"

	"github.com/gotd/td/bin"
)

func countPadding(l int) int { return 16 + (16 - (l % 16)) }

// encryptMessage encrypts plaintext using AES-IGE.
func (c Cipher) encryptMessage(k AuthKey, plaintext *bin.Buffer) (EncryptedMessage, error) {
	offset := len(plaintext.Buf)
	padding := countPadding(offset)
	messageLength := offset + padding
	plaintext.Expand(padding)
	if _, err := io.ReadFull(c.rand, plaintext.Buf[offset:messageLength]); err != nil {
		return EncryptedMessage{}, err
	}
	plaintext.Expand(messageLength)
	plaintext.Buf = plaintext.Buf[:messageLength]

	messageKey := MessageKey(k.Value, plaintext.Buf, c.encryptSide)
	key, iv := Keys(k.Value, messageKey, c.encryptSide)
	aesBlock, err := aes.NewCipher(key[:])
	if err != nil {
		return EncryptedMessage{}, err
	}
	msg := EncryptedMessage{
		AuthKeyID:     k.ID,
		MsgKey:        messageKey,
		EncryptedData: plaintext.Buf[messageLength : 2*messageLength: 2*messageLength],
	}
	ige.EncryptBlocks(aesBlock, iv[:], msg.EncryptedData, plaintext.Buf)
	return msg, nil
}

// Encrypt encrypts EncryptedMessageData using AES-IGE to given buffer.
func (c Cipher) Encrypt(key AuthKey, data EncryptedMessageData, b *bin.Buffer) error {
	b.Reset()
	if err := data.EncodeWithoutCopy(b); err != nil {
		return err
	}

	msg, err := c.encryptMessage(key, b)
	if err != nil {
		return err
	}

	b.Reset()
	if err := msg.Encode(b); err != nil {
		return err
	}

	return nil
}
