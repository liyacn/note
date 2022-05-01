// https://gorm.io/

package db

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	glog "gorm.io/gorm/logger"
	"log"
	"net"
	"time"
)

type Config struct {
	Address  string
	Username string
	Password string
	Database string
	MaxOpen  int
	MaxIdle  int
	TraceLog bool
}

func NewMysql(cfg *Config) *gorm.DB {
	dsn := cfg.Username + ":" + cfg.Password + "@tcp(" + cfg.Address + ")/" + cfg.Database +
		"?charset=utf8mb4&parseTime=true&loc=Local"
	return newDB(mysql.Open(dsn), cfg)
}

func NewPostgres(cfg *Config) *gorm.DB {
	host, port, _ := net.SplitHostPort(cfg.Address)
	dsn := "host=" + host + " port=" + port + " user=" + cfg.Username + " password=" + cfg.Password +
		" dbname=" + cfg.Database + " sslmode=disable TimeZone=Asia/Shanghai"
	return newDB(postgres.Open(dsn), cfg)
}

func NewSqlserver(cfg *Config) *gorm.DB {
	dsn := "sqlserver://" + cfg.Username + ":" + cfg.Password +
		"@" + cfg.Address + "?database=" + cfg.Database
	return newDB(sqlserver.Open(dsn), cfg)
}

func newDB(dial gorm.Dialector, cfg *Config) *gorm.DB {
	opt := &gorm.Config{Logger: glog.Discard.LogMode(glog.Silent)}
	if cfg.TraceLog {
		opt.Logger = &gormLog{opt.Logger}
	}
	orm, err := gorm.Open(dial, opt)
	if err != nil {
		log.Fatal(err)
	}
	sqlDB, _ := orm.DB()
	if cfg.MaxOpen > 0 {
		sqlDB.SetMaxOpenConns(cfg.MaxOpen)
	}
	if cfg.MaxIdle > 0 {
		sqlDB.SetMaxIdleConns(cfg.MaxIdle)
	}
	//sqlDB.SetConnMaxLifetime(time.Hour)
	//sqlDB.SetConnMaxIdleTime(time.Minute)
	return orm
}

func Close(orm *gorm.DB) error {
	sqlDB, _ := orm.DB()
	return sqlDB.Close()
}

type gormLog struct {
	glog.Interface
}

// Trace 方法重写覆盖，自定义打印日志的方法
func (*gormLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rows int64), err error) {
	sql, rows := fc()
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println(ctx.Value("trace_id"), sql, err)
	} else {
		log.Println(ctx.Value("trace_id"), sql, rows, time.Since(begin))
	}
}
