package redis

import (
	"github.com/atpuxiner/grapi/app/initializer/conf"
	goredis "github.com/go-redis/redis"
	"time"
)

var Redis *goredis.Client

func InitRedis() {
	Redis = goredis.NewClient(&goredis.Options{
		Addr:         conf.Conf.Redis.Addr,
		Password:     conf.Conf.Redis.Password,
		DB:           conf.Conf.Redis.DB,
		DialTimeout:  conf.Conf.Redis.DialTimeout * time.Second,
		ReadTimeout:  conf.Conf.Redis.ReadTimeout * time.Second,
		WriteTimeout: conf.Conf.Redis.WriteTimeout * time.Second,
		PoolSize:     conf.Conf.Redis.PoolSize,
		PoolTimeout:  conf.Conf.Redis.PoolTimeout * time.Second,
	})
	if err := Redis.Ping().Err(); err != nil {
		panic("Redis连接失败：" + err.Error())
	}
}
