// https://redis.uptrace.dev/

package gredis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

type Config struct {
	Address    string
	MasterName string
	Addrs      []string
	Username   string
	Password   string
	DB         int
	PoolSize   int
	MinIdle    int
	MaxIdle    int
}

func NewClient(cfg *Config) *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:               cfg.Address,  // default: "localhost:6379"
		Username:           cfg.Username, // redis 6.0+
		Password:           cfg.Password, //
		DB:                 cfg.DB,       //
		MaxRetries:         0,            // default: 3
		MinRetryBackoff:    0,            // default: 8 milliseconds
		MaxRetryBackoff:    0,            // default: 512 milliseconds
		DialTimeout:        0,            // default: 5 seconds
		DialerRetries:      0,            // default: 5
		DialerRetryTimeout: 0,            // default: 100 milliseconds
		ReadTimeout:        0,            // default: 3 seconds
		WriteTimeout:       0,            // default: same as ReadTimeout
		ReadBufferSize:     0,            // default: 32KiB
		WriteBufferSize:    0,            // default: 32KiB
		PoolSize:           cfg.PoolSize, // default: 10 * runtime.GOMAXPROCS(0)
		PoolTimeout:        0,            // default: ReadTimeout + 1 second
		MinIdleConns:       cfg.MinIdle,  //
		MaxIdleConns:       cfg.MaxIdle,  //
		ConnMaxIdleTime:    0,            // default: 30 minutes
		TLSConfig:          nil,          //
	})
	if err := cli.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}
	return cli
}

func NewClusterClient(cfg *Config) *redis.ClusterClient {
	cli := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:                 cfg.Addrs,    //
		MaxRedirects:          0,            // default: 3
		Username:              cfg.Username, // default: ""
		Password:              cfg.Password, // default: ""
		MaxRetries:            0,            // default: 3
		MinRetryBackoff:       0,            // default: 8 milliseconds
		MaxRetryBackoff:       0,            // default: 512 milliseconds
		DialTimeout:           0,            // default: 5 seconds
		ReadTimeout:           0,            // default: 3 seconds
		WriteTimeout:          0,            // default: same as ReadTimeout
		PoolSize:              0,            // default: 5 * runtime.GOMAXPROCS(0)
		PoolTimeout:           0,            // default: ReadTimeout + 1 second
		MinIdleConns:          cfg.MinIdle,  //
		MaxIdleConns:          cfg.MaxIdle,  //
		ConnMaxIdleTime:       0,            // default: 30 minutes
		ReadBufferSize:        0,            // default: 32KiB
		WriteBufferSize:       0,            // default: 32KiB
		TLSConfig:             nil,          //
		FailingTimeoutSeconds: 0,            // default: 15
	})
	if err := cli.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}
	return cli
}

func NewFailoverClient(cfg *Config) *redis.Client {
	cli := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:       cfg.MasterName,
		SentinelAddrs:    cfg.Addrs,
		SentinelUsername: "",
		SentinelPassword: "",
		Username:         cfg.Username,
		Password:         cfg.Password,
		DB:               cfg.DB,
		MaxRetries:       0,
		MinRetryBackoff:  0,
		MaxRetryBackoff:  0,
		DialTimeout:      0,
		ReadTimeout:      0,
		WriteTimeout:     0,
		ReadBufferSize:   0,
		WriteBufferSize:  0,
		PoolSize:         cfg.PoolSize,
		PoolTimeout:      0,
		MinIdleConns:     cfg.MinIdle,
		MaxIdleConns:     cfg.MaxIdle,
		ConnMaxLifetime:  0,
		TLSConfig:        nil,
	}) // 默认值同 NewClient
	if err := cli.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}
	return cli
}
