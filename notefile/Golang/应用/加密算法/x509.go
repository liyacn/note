package crypt

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
)

func TlsConfig(servername, cert, key, ca string) (*tls.Config, error) {
	cfg := &tls.Config{ServerName: servername}
	if cert != "" && key != "" {
		certificate, err := tls.X509KeyPair([]byte(cert), []byte(key))
		if err != nil {
			return nil, err
		}
		cfg.Certificates = []tls.Certificate{certificate}
	}
	if ca != "" {
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM([]byte(ca)) {
			return nil, errors.New("failed to parse root certificate")
		}
		cfg.RootCAs = pool
	}
	return cfg, nil
}
