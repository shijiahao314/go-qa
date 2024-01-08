package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router interface {
	Register(r *gin.RouterGroup)
}

func Register(r *gin.Engine) {
	rr := r.Group("/api")

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
	rr.Use(cors.Default())

	rts := []Router{
		&AuthRouter{},
		&AdminRouter{},
		&ChatRouter{},
		&ChatWSRouter{},
	}

	for _, rt := range rts {
		rt.Register(rr)
	}
}
