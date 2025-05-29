package config

import (
	"github.com/zeromicro/go-zero/rest"
	"time"
)

type Config struct {
	rest.RestConf
	JwtAuth JwtAuth
	Redis   RedisConf
	Mysql   MysqlConf
}

type JwtAuth struct {
	AccessSecret string `json:",default=secret"`
	AccessExpire int64  `json:",default=3600"`
}

type RedisConf struct {
	Host     string `default:"localhost"`
	Password string `default:"password"`
	Port     int    `default:"6379"`
	DB       int    `default:"0"`
}

type MysqlConf struct {
	Dsn         string        `default:"dsn"`
	MaxIdle     int           `default:"10"` // 空闲链接 空闲链接 + 活跃链接 <= 最大链接
	MaxIdleTime time.Duration `default:"1800"`
	//MaxLifeTime int    `default:"3600"` // 最大存活时间 （包含空闲等待时间）
	MaxConn int `default:"100"` // 最大链接数
}
