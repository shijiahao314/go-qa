package bootstrap

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/redis/go-redis/v9"
	"github.com/shijiahao314/go-qa/global"
)

// mustInitRedis 初始化 Redis
func mustInitRedis() *redis.Client {
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
		panic(fmt.Sprintf("failed to init redis: %s", err.Error()))
	}

	slog.Info("success init redis, ping response:", slog.String("pong", pong))
	return client
}
