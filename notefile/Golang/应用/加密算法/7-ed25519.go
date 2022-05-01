package crypt

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/base64"
)

func ParseEd25519PrivateKey(key string) (ed25519.PrivateKey, error) {
	der, err := key2der(key)
	if err != nil {
		return nil, err
	}
	pri, err := x509.ParsePKCS8PrivateKey(der)
	if err != nil {
		return nil, err
	}
	if v, ok := pri.(ed25519.PrivateKey); ok {
		return v, nil
	}
	return nil, ErrInvalidKeyType
}

func ParseEd25519PublicKey(key string) (ed25519.PublicKey, error) {
	der, err := key2der(key)
	if err != nil {
		return nil, err
	}
	pub, err := x509.ParsePKIXPublicKey(der)
	if err != nil {
		return nil, err
	}
	if v, ok := pub.(ed25519.PublicKey); ok {
		return v, nil
	}
	return nil, ErrInvalidKeyType
}

func Ed25519Sign(key ed25519.PrivateKey, msg []byte) string {
	sign := ed25519.Sign(key, msg)
	return base64.StdEncoding.EncodeToString(sign)
}

func Ed25519Verify(key ed25519.PublicKey, msg []byte, sign string) bool {
	sig, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false
	}
	return ed25519.Verify(key, msg, sig)
}
