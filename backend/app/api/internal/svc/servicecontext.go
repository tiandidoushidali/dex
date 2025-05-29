package svc

import (
	"dex/app/api/internal/config"
	"github.com/redis/go-redis/v9"
	"github.com/sony/sonyflake"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config      config.Config
	SF          *sonyflake.Sonyflake
	RedisClient *redis.Client
	MysqlClient *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{})
	if sf == nil {
		panic("sonyflake not created")
	}
	db, err := gorm.Open(mysql.Open(c.Mysql.Dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(c.Mysql.MaxIdle)
	sqlDB.SetConnMaxIdleTime(c.Mysql.MaxIdleTime)
	sqlDB.SetMaxOpenConns(c.Mysql.MaxConn)

	redisClient := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: c.Redis.Password,
		DB:       c.Redis.DB,
	})

	return &ServiceContext{
		Config:      c,
		SF:          sf,
		MysqlClient: db,
		RedisClient: redisClient,
	}
}
