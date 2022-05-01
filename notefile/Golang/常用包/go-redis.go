// https://redis.uptrace.dev/

package gredis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/redis/go-redis/v9/maintnotifications"
	"log"
	"log/slog"
	"net"
	"project/pkg/crypt"
)

type Config struct {
	Address  string
	Username string
	Password string
	DB       int
	PoolSize int
	MinIdle  int
	MaxIdle  int
	TraceLog bool
	TLS      *crypt.TlsConfig
}

func NewClient(cfg *Config) *redis.Client {
	cli := redis.NewClient(&redis.Options{
		Addr:                  cfg.Address,                  // default: "localhost:6379"
		Username:              cfg.Username,                 //
		Password:              cfg.Password,                 //
		DB:                    cfg.DB,                       //
		MaxRetries:            0,                            // default: 3
		MinRetryBackoff:       0,                            // default: 8 milliseconds
		MaxRetryBackoff:       0,                            // default: 512 milliseconds
		DialTimeout:           0,                            // default: 5 seconds
		DialerRetries:         0,                            // default: 5
		DialerRetryTimeout:    0,                            // default: 100 milliseconds
		ReadTimeout:           0,                            // default: 3 seconds
		WriteTimeout:          0,                            // default: same as ReadTimeout
		ReadBufferSize:        0,                            // default: 32KiB
		WriteBufferSize:       0,                            // default: 32KiB
		PoolSize:              cfg.PoolSize,                 // default: 10 * runtime.GOMAXPROCS(0)
		PoolTimeout:           0,                            // default: ReadTimeout + 1 second
		MinIdleConns:          cfg.MinIdle,                  //
		MaxIdleConns:          cfg.MaxIdle,                  //
		MaxActiveConns:        0,                            //
		ConnMaxIdleTime:       0,                            // default: 30 minutes
		ConnMaxLifetime:       0,                            //
		TLSConfig:             crypt.MustTlsConfig(cfg.TLS), //
		FailingTimeoutSeconds: 0,                            // default: 15
		MaintNotificationsConfig: &maintnotifications.Config{
			Mode: maintnotifications.ModeDisabled, // default: ModeAuto
		},
	})
	if err := cli.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}
	if cfg.TraceLog {
		cli.AddHook(debugHook{})
	}
	return cli
}

type debugHook struct{}

func (debugHook) DialHook(next redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		return next(ctx, network, addr)
	}
}
func (debugHook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		slog.DebugContext(ctx, "redis", "cmd", cmd.String())
		return next(ctx, cmd)
	}
}
func (debugHook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		for _, cmd := range cmds {
			slog.DebugContext(ctx, "redis", "cmd", cmd.String())
		}
		return next(ctx, cmds)
	}
}
