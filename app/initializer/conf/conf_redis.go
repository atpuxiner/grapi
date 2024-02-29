package conf

import (
	"time"
)

type RedisConf struct {
	Addr         string        `mapstructure:"addr"`
	Password     string        `mapstructure:"password"`
	DB           int           `mapstructure:"db"`
	DialTimeout  time.Duration `mapstructure:"dialTimeout"`
	ReadTimeout  time.Duration `mapstructure:"readTimeout"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout"`
	PoolSize     int           `mapstructure:"poolSize"`
	PoolTimeout  time.Duration `mapstructure:"poolTimeout"`
}
