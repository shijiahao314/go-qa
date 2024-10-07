package bootstrap

import (
	"crypto/rand"
	"log/slog"
	"math/big"

	"github.com/boj/redistore"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/router"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}<>?/"

// generateRandomString 生成随机字符串
func generateRandomString(length int) (string, error) {
	result := make([]byte, length)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		// 从字符集中随机选择一个字符
		index, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		result[i] = charset[index.Int64()]
	}

	return string(result), nil
}

const (
	HealthPath      = "/api/health"
	SecretKeyLength = 32
)

// MustInitRouter 初始化路由配置
func MustInitRouter() *gin.Engine {
	r := gin.New()

	secretKey, err := generateRandomString(SecretKeyLength)
	if err != nil {
		slog.Error("failed to generate random string", slog.String("err", err.Error()))
		panic(err)
	}

	// session
	var store sessions.Store
	if global.Redis != nil {
		// Redis 可用则使用
		store, err := redis.NewStore(
			global.Config.Redis.ConnectionNum,
			"tcp",
			global.Config.Redis.Addr,
			global.Config.Redis.Password,
			[]byte(secretKey),
		)
		if err != nil {
			slog.Error("failed to init redis", slog.String("err", err.Error()))
			panic(err)
		}
		_, rs := redis.GetRedisStore(store)

		rs.SetSerializer(redistore.JSONSerializer{})

		store.Options(sessions.Options{
			Path:     "/",
			MaxAge:   60 * 60 * 24,
			Secure:   false,
			HttpOnly: false,
		})
	} else {
		// 如果无法使用 Redis 则使用 cookie
		store = cookie.NewStore([]byte(secretKey))
	}
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
