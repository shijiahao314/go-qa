package bootstrap

import (
	"github.com/boj/redistore"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	// gSessions "github.com/gorilla/sessions"
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/router"
)

const (
	HealthPath = "/api/health"
	SecretKey  = "iV6pNvjdHVUVc5Q*Wi4S&" // random
)

// type MySerializer struct{}

// func (ms MySerializer) Serialize(ss *gSessions.Session) ([]byte, error) {
// 	m := make(map[string]interface{}, len(ss.Values))
// 	for k, v := range ss.Values {
// 		ks, ok := k.(string)
// 		if !ok {
// 			err := fmt.Errorf("Non-string key value, cannot serialize session to JSON: %v", k)
// 			fmt.Printf("redistore.JSONSerializer.serialize() Error: %v", err)
// 			return nil, err
// 		}
// 		fmt.Printf("[%s]=[%s]\n", k, v)
// 		m[ks] = v
// 	}
// 	return json.Marshal(m)
// }

// func (ms MySerializer) Deserialize(d []byte, ss *gSessions.Session) error {
// 	m := make(map[string]interface{})
// 	err := json.Unmarshal(d, &m)
// 	if err != nil {
// 		fmt.Printf("redistore.JSONSerializer.deserialize() Error: %v", err)
// 		return err
// 	}
// 	for k, v := range m {
// 		fmt.Printf("[%s]=[%s]\n", k, v)
// 		ss.Values[k] = v
// 	}
// 	return nil
// }

func InitRouter() *gin.Engine {
	r := gin.New()

	// session
	store, err := redis.NewStore(
		global.Config.Redis.ConnectionNum,
		"tcp",
		global.Config.Redis.Addr,
		global.Config.Redis.Password,
		[]byte(SecretKey),
	)
	if err != nil {
		global.Logger.Error("failed to init redis", zap.Error(err))
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
