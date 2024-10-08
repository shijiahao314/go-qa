package bootstrap

import (
	"log/slog"

	"github.com/boj/redistore"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/router"
)

const (
	HealthPath = "/api/health"
	SecretKey  = "95osj3fUD7fo0mlYdDbncXz4VD2igvf0"
)

// MustInitRouter 初始化路由配置
func MustInitRouter() *gin.Engine {
	r := gin.New()

	// session 设置参考：https://www.cnblogs.com/taoxiaoxin/p/17991891
	// 使用 Redis（适合多机部署）
	if global.Redis == nil {
		// Redis 不可用
		slog.Error("multi mode requires redis")
		panic("multi mode requires redis")
	}

	store, err := redis.NewStore(
		global.Config.Redis.ConnectionNum,
		"tcp",
		global.Config.Redis.Addr,
		global.Config.Redis.Password,
		[]byte(SecretKey), // 多机部署使用同一个密钥
	)
	if err != nil {
		slog.Error("failed to init redis", slog.String("err", err.Error()))
		panic(err)
	}

	// 设置序列化器，不设置会导致 gob 报错
	err, rs := redis.GetRedisStore(store)
	if err != nil {
		slog.Error("failed to get redis store", slog.String("err", err.Error()))
		panic(err)
	}
	rs.SetSerializer(redistore.JSONSerializer{})

	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 24,
		Secure:   false,
		HttpOnly: false,
	})

	r.Use(sessions.Sessions("session", store))

	r.Use(
		gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: []string{HealthPath}}),
		gin.Recovery(),
	)

	r.Use(cors.Default())

	r.GET(HealthPath, func(*gin.Context) {})

	router.Register(r)

	return r
}
