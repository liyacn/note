package crypt

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"hash"
	"strings"
)

var ErrInvalidKeyType = errors.New("invalid key type")

func key2der(key string) ([]byte, error) {
	if strings.Contains(key, "-----") {
		block, _ := pem.Decode([]byte(key))
		if block == nil {
			return nil, errors.New("invalid pem text")
		}
		return block.Bytes, nil
	}
	return base64.StdEncoding.DecodeString(key)
}

func ParsePkcs1PrivateKey(key string) (*rsa.PrivateKey, error) {
	der, err := key2der(key)
	if err != nil {
		return nil, err
	}
	return x509.ParsePKCS1PrivateKey(der)
}

func ParsePkcs1PublicKey(key string) (*rsa.PublicKey, error) {
	der, err := key2der(key)
	if err != nil {
		return nil, err
	}
	return x509.ParsePKCS1PublicKey(der)
}

func ParseRsaPrivateKey(key string) (*rsa.PrivateKey, error) {
	der, err := key2der(key)
	if err != nil {
		return nil, err
	}
	pri, err := x509.ParsePKCS8PrivateKey(der)
	if err != nil {
		return nil, err
	}
	if v, ok := pri.(*rsa.PrivateKey); ok {
		return v, nil
	}
	return nil, ErrInvalidKeyType
}

func ParseRsaPublicKey(key string) (*rsa.PublicKey, error) {
	der, err := key2der(key)
	if err != nil {
		return nil, err
	}
	pub, err := x509.ParsePKIXPublicKey(der)
	if err != nil {
		return nil, err
	}
	if v, ok := pub.(*rsa.PublicKey); ok {
		return v, nil
	}
	return nil, ErrInvalidKeyType
}

func RsaEncryptPKCS1v15(key *rsa.PublicKey, src []byte) (string, error) {
	chunk := key.Size() - 11 //长文本需要分段加密
	end := len(src)
	buf := bytes.NewBuffer(nil)
	l := 0
	for l < end {
		r := l + chunk
		if r > end {
			r = end
		}
		b, err := rsa.EncryptPKCS1v15(rand.Reader, key, src[l:r])
		if err != nil {
			return "", err
		}
		buf.Write(b)
		l = r
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func RsaDecryptPKCS1v15(key *rsa.PrivateKey, s string) ([]byte, error) {
	src, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	chunk := key.Size() //长文本需要分段解密
	end := len(src)
	buf := bytes.NewBuffer(nil)
	l := 0
	for l < end {
		r := l + chunk
		if r > end {
			r = end
		}
		b, err := rsa.DecryptPKCS1v15(rand.Reader, key, src[l:r])
		if err != nil {
			return nil, err
		}
		buf.Write(b)
		l = r
	}
	return buf.Bytes(), nil
}

func RsaEncryptOAEP(h hash.Hash, key *rsa.PublicKey, src []byte, label []byte) (string, error) {
	chunk := key.Size() - 2*h.Size() - 2 //长文本需要分段加密
	if chunk < 1 {
		return "", errors.New("hash size is too long for this key size")
	}
	end := len(src)
	buf := bytes.NewBuffer(nil)
	l := 0
	for l < end {
		r := l + chunk
		if r > end {
			r = end
		}
		b, err := rsa.EncryptOAEP(h, rand.Reader, key, src[l:r], label)
		if err != nil {
			return "", err
		}
		buf.Write(b)
		l = r
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func RsaDecryptOAEP(h hash.Hash, key *rsa.PrivateKey, s string, label []byte) ([]byte, error) {
	src, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	chunk := key.Size() //长文本需要分段解密
	end := len(src)
	buf := bytes.NewBuffer(nil)
	l := 0
	for l < end {
		r := l + chunk
		if r > end {
			r = end
		}
		b, err := rsa.DecryptOAEP(h, rand.Reader, key, src[l:r], label)
		if err != nil {
			return nil, err
		}
		buf.Write(b)
		l = r
	}
	return buf.Bytes(), nil
}

func RsaSignPKCS1v15(key *rsa.PrivateKey, h crypto.Hash, data []byte) (string, error) {
	w := h.New()
	w.Write(data)
	sign, err := rsa.SignPKCS1v15(rand.Reader, key, h, w.Sum(nil))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sign), nil
}

func RsaVerifyPKCS1v15(key *rsa.PublicKey, h crypto.Hash, data []byte, sign string) error {
	sig, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	w := h.New()
	w.Write(data)
	return rsa.VerifyPKCS1v15(key, h, w.Sum(nil), sig)
}

func RsaSignPSS(key *rsa.PrivateKey, h crypto.Hash, data []byte) (string, error) {
	w := h.New()
	w.Write(data)
	sign, err := rsa.SignPSS(rand.Reader, key, h, w.Sum(nil), nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sign), nil
}

func RsaVerifyPSS(key *rsa.PublicKey, h crypto.Hash, data []byte, sign string) error {
	sig, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	w := h.New()
	w.Write(data)
	return rsa.VerifyPSS(key, h, w.Sum(nil), sig, nil)
}
