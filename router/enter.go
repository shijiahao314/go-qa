package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shijiahao314/go-qa/api"
	"github.com/shijiahao314/go-qa/middleware"
)

type IRouter interface {
	Register(r *gin.RouterGroup)
}

func Register(r *gin.Engine) {
	apiRouter := r.Group("/api")
	authRouter := r.Group("/api")

	// 允许跨域
	// rr.Use(cors.New(cors.Config{
	// 	AllowAllOrigins: true,
	// 	// AllowOrigins:     []string{"http://10.129.246.191:3000", "http://10.112.188.168:3000"},
	// 	AllowMethods:     []string{"GET", "POST"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))

	// apiRouter
	apiRouter.Use(cors.Default(), middleware.Auth())
	rts1 := []IRouter{
		&api.AdminAPI{},
		&api.ChatAPI{},
		&api.ChatWSApi{},
		&api.SettingAPI{},
	}
	for _, rt := range rts1 {
		rt.Register(apiRouter)
	}

	// authRouter
	rts2 := []IRouter{
		&api.AuthAPI{},
	}
	for _, rt := range rts2 {
		rt.Register(authRouter)
	}
}
