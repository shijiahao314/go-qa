package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shijiahao314/go-qa/api"
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/middleware"
)

type AdminRouter struct{}

func (ar *AdminRouter) Register(rg *gin.RouterGroup) {
	r := rg.Group("/admin")
	adminApi := new(api.AdminApi)

	r.Use(middleware.Auth(), middleware.Role([]string{global.ROLE_ADMIN}))

	r.GET("/user", adminApi.GetUsers)
	r.POST("/user", adminApi.AddUser)
	r.POST("/user/:id", adminApi.UpdateUser)
	r.DELETE("/user/:id", adminApi.DeleteUser)
}
