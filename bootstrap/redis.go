package bootstrap

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/shijiahao314/go-qa/global"
	"go.uber.org/zap"
)

func initRedis() *redis.Client {
	global.Logger.Info("start to init redis")
	redisCfg := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Logger.Error("failed to init redis", zap.Error(err))
		panic(err)
	}

	global.Logger.Info("successfully init redis, ping response:", zap.String("pong", pong))
	return client
}
