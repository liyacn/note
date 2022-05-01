package crypt

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

const format = "%s\n%s\n"

func der2pem(b []byte, t string) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Type:  t,
		Bytes: b,
	})
}

func GenerateRsaKey(multiple int) {
	key, _ := rsa.GenerateKey(rand.Reader, 512<<(multiple&3))
	priDer1 := x509.MarshalPKCS1PrivateKey(key)
	priDer8, _ := x509.MarshalPKCS8PrivateKey(key)
	pubDer1 := x509.MarshalPKCS1PublicKey(&key.PublicKey)
	pubDer8, _ := x509.MarshalPKIXPublicKey(&key.PublicKey)

	fmt.Printf(format, "pkcs1 private row:", base64.StdEncoding.EncodeToString(priDer1))
	fmt.Printf(format, "pkcs1 private pem:", der2pem(priDer1, "RSA PRIVATE KEY"))
	fmt.Printf(format, "pkcs8 private row:", base64.StdEncoding.EncodeToString(priDer8))
	fmt.Printf(format, "pkcs8 private pem:", der2pem(priDer8, "PRIVATE KEY"))

	fmt.Printf(format, "pkcs1 public row:", base64.StdEncoding.EncodeToString(pubDer1))
	fmt.Printf(format, "pkcs1 public pem:", der2pem(pubDer1, "RSA PUBLIC KEY"))
	fmt.Printf(format, "pkcs8 public row:", base64.StdEncoding.EncodeToString(pubDer8))
	fmt.Printf(format, "pkcs8 public pem:", der2pem(pubDer8, "PUBLIC KEY"))
}

func GenerateEcdsaKey(p int) {
	var key *ecdsa.PrivateKey
	switch p {
	case 224:
		key, _ = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case 256:
		key, _ = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case 384:
		key, _ = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	default:
		key, _ = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	}
	priDer, _ := x509.MarshalPKCS8PrivateKey(key)
	pubDer, _ := x509.MarshalPKIXPublicKey(&key.PublicKey)

	fmt.Printf(format, "ecdsa private row:", base64.StdEncoding.EncodeToString(priDer))
	fmt.Printf(format, "ecdsa private pem:", der2pem(priDer, "PRIVATE KEY"))
	fmt.Printf(format, "ecdsa public row:", base64.StdEncoding.EncodeToString(pubDer))
	fmt.Printf(format, "ecdsa public pem:", der2pem(pubDer, "PUBLIC KEY"))
}

func GenerateEd25519Key() {
	pub, pri, _ := ed25519.GenerateKey(rand.Reader)
	priDer, _ := x509.MarshalPKCS8PrivateKey(pri)
	pubDer, _ := x509.MarshalPKIXPublicKey(pub)
	fmt.Printf(format, "ed25519 private row:", base64.StdEncoding.EncodeToString(priDer))
	fmt.Printf(format, "ed25519 public row:", base64.StdEncoding.EncodeToString(pubDer))
}
