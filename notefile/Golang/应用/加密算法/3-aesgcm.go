package ades

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
)

// DES和3DES的块大小均为8，AES的块大小为16，GCM的块大小必须为16，所以只支持AES算法。
// 密文和nonce可以分别编码存储，这里选择合并编码存储。

type AesGCM struct {
	aead cipher.AEAD
}

func NewAesGCM(key []byte) (*AesGCM, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	c := &AesGCM{aead: aead}
	return c, nil
}

func (c *AesGCM) Encrypt(src []byte) string {
	nonce := make([]byte, c.aead.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		panic(err)
	}
	b := c.aead.Seal(nil, nonce, src, nil)
	return base64.StdEncoding.EncodeToString(append(b, nonce...))
}

func (c *AesGCM) Decrypt(s string) ([]byte, error) {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	length := len(b)
	if length < c.aead.NonceSize() {
		return nil, ErrInvalidPadLen
	}
	pos := length - c.aead.NonceSize()
	return c.aead.Open(nil, b[pos:], b[:pos], nil)
}
