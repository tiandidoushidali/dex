package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Host     string
	Port     string
	UserName string
	Password string
}

type RedisConfig struct {
	Host     string
	Port     string
	UserName string
	Password string
	Database string
}

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "aba", "path to config file")
}

func Setup(configPath string) {
	fmt.Println("flag config path:", configPath)
	viper.AddConfigPath(configPath)
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
}
