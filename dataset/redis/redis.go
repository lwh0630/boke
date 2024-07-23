package redis

import (
	"bluebell/config"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

var (
	rdb *redis.Client
	Nil = redis.Nil
)

func InitRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Cfg.RedisConfig.Host, config.Cfg.RedisConfig.Port),
		Password: config.Cfg.RedisConfig.Password,
		DB:       config.Cfg.RedisConfig.DB,
		PoolSize: config.Cfg.RedisConfig.PoolSize,
	})
	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		zap.L().Error("redis init failed", zap.Error(err))
	}
	return err
}

func CloseRedis() {
	err := rdb.Close()
	if err != nil {
		fmt.Printf("redis close failed: %s\n", err)
	}
	fmt.Println("redis closed")
}
