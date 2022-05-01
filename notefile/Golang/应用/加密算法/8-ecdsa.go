package crypt

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/base64"
)

func ParseEcdsaPrivateKey(key string) (*ecdsa.PrivateKey, error) {
	der, err := key2der(key)
	if err != nil {
		return nil, err
	}
	pri, err := x509.ParsePKCS8PrivateKey(der)
	if err != nil {
		return nil, err
	}
	if v, ok := pri.(*ecdsa.PrivateKey); ok {
		return v, nil
	}
	return nil, ErrInvalidKeyType
}

func ParseEcdsaPublicKey(key string) (*ecdsa.PublicKey, error) {
	der, err := key2der(key)
	if err != nil {
		return nil, err
	}
	pub, err := x509.ParsePKIXPublicKey(der)
	if err != nil {
		return nil, err
	}
	if v, ok := pub.(*ecdsa.PublicKey); ok {
		return v, nil
	}
	return nil, ErrInvalidKeyType
}

func EcdsaSignASN1(key *ecdsa.PrivateKey, hashed []byte) (string, error) {
	sign, err := ecdsa.SignASN1(rand.Reader, key, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sign), nil
}

func EcdsaVerifyASN1(key *ecdsa.PublicKey, hashed []byte, sign string) bool {
	sig, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false
	}
	return ecdsa.VerifyASN1(key, hashed, sig)
}
