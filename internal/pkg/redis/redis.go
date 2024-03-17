package redis

import (
	"fmt"
	"go.uber.org/zap"

	"github.com/go-redis/redis"
	"wild/configs"
)

var RDB *redis.Client

func Connect(config *configs.RedisConfig) *redis.Client {

	address := fmt.Sprintf("%s:%d", config.Host, config.Port)

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
		PoolSize: config.PoolSize,
	})

	zap.L().Info(`🍟: Successfully connected to Redis at ` + address)

	return rdb

}

func InitRedis() error {

	RDB = Connect(configs.Conf.RedisConfig)

	return nil
}

func CloseRedis() error {
	if err := RDB.Close(); err != nil {
		return err
	}
	return nil
}
