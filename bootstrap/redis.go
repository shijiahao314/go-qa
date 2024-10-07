package bootstrap

import (
	"context"
	"log/slog"

	"github.com/redis/go-redis/v9"
	"github.com/shijiahao314/go-qa/global"
)

// initRedis 初始化 Redis
func initRedis() (*redis.Client, error) {
	slog.Info("start to init redis")

	redisCfg := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		slog.Error("failed to init redis", slog.String("err", err.Error()))
		return nil, err
	}

	slog.Info("success init redis, ping response:", slog.String("pong", pong))
	return client, nil
}
