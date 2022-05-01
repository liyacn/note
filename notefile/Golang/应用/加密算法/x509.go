package crypt

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"log"
)

type TlsConfig struct {
	ServerName string
	Cert       string
	Key        string
	Ca         string
}

func ParseTlsConfig(cfg *TlsConfig) (*tls.Config, error) {
	if cfg == nil || (cfg.ServerName == "" && (cfg.Cert == "" || cfg.Key == "") && cfg.Ca == "") {
		return nil, nil
	}
	tlsConfig := &tls.Config{ServerName: cfg.ServerName}
	if cfg.Cert != "" && cfg.Key != "" {
		certificate, err := tls.X509KeyPair([]byte(cfg.Cert), []byte(cfg.Key))
		if err != nil {
			return nil, err
		}
		tlsConfig.Certificates = []tls.Certificate{certificate}
	}
	if cfg.Ca != "" {
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM([]byte(cfg.Ca)) {
			return nil, errors.New("failed to parse root certificate")
		}
		tlsConfig.RootCAs = pool
	}
	return tlsConfig, nil
}

func MustTlsConfig(cfg *TlsConfig) *tls.Config {
	c, err := ParseTlsConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
