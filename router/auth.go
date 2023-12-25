package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shijiahao314/go-qa/api"
)

type AuthRouter struct{}

func (*AuthRouter) Register(r *gin.RouterGroup) {
	rt := r.Group("/auth")
	authApi := &api.AuthApi{}

	// rt.Use(middleware.Auth(), middleware.Role([]string{global.ROLE_ADMIN}))
	// store
	// rt.User(sessions.Sessions("mySession", global.Redis))

	rt.POST("/signup", authApi.SignUp)
	rt.POST("/login", authApi.Login)
	rt.POST("/logout", authApi.Logout)
	rt.GET("/islogin", authApi.IsLogin)
}
