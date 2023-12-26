package router

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	Register(r *gin.RouterGroup)
}

func Register(r *gin.Engine) {
	rr := r.Group("/api")

	// 允许跨域
	// rr.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"https://10.29.146.70"},
	// 	AllowMethods:     []string{"GET", "POST"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "https://github.com"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))

	rts := []Router{
		&AuthRouter{},
		&AdminRouter{},
		&ChatRouter{},
	}

	for _, rt := range rts {
		rt.Register(rr)
	}
}
