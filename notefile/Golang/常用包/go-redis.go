// https://redis.uptrace.dev/

package gredis

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/redis/go-redis/v9"
	"log"
)

type Config struct {
	Address       string
	Username      string
	Password      string
	DB            int
	PoolSize      int
	MinIdle       int
	MaxIdle       int
	Cert, Key, Ca string
}

func NewClient(cfg *Config) *redis.Client {
	var tlsConfig *tls.Config
	if cfg.Cert != "" && cfg.Key != "" && cfg.Ca != "" {
		certificate, err := tls.X509KeyPair([]byte(cfg.Cert), []byte(cfg.Key))
		if err != nil {
			log.Fatal(err)
		}
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM([]byte(cfg.Ca)) {
			log.Fatal("failed to parse root certificate")
		}
		tlsConfig = &tls.Config{
			Certificates: []tls.Certificate{certificate},
			RootCAs:      pool,
		}
	}
	cli := redis.NewClient(&redis.Options{
		Network:               "",           // Default is tcp.
		Addr:                  cfg.Address,  // Default localhost:6379
		Username:              cfg.Username, // Redis 6.0 以上使用
		Password:              cfg.Password, // Default empty
		DB:                    cfg.DB,       // Default 0
		MaxRetries:            0,            // Default is 3 retries
		MinRetryBackoff:       0,            // Default is 8 milliseconds
		MaxRetryBackoff:       0,            // Default is 512 milliseconds
		DialTimeout:           0,            // Default is 5 seconds
		ReadTimeout:           0,            // Default is 3 seconds
		WriteTimeout:          0,            // Default is 3 seconds
		ContextTimeoutEnabled: false,        // 是否启用context超时控制
		PoolSize:              cfg.PoolSize, // Default is 10 connections per every CPU
		PoolTimeout:           0,            // Default is ReadTimeout + 1 second.
		MinIdleConns:          cfg.MinIdle,  // 建议PoolSize的(1/10~1/5)
		MaxIdleConns:          cfg.MaxIdle,  // 最大空闲连接
		ConnMaxIdleTime:       0,            // Default is 30 minutes.
		ConnMaxLifetime:       0,            // Default is to not close idle connections.
		TLSConfig:             tlsConfig,
	})
	if err := cli.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}
	return cli
}
