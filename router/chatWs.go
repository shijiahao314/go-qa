package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shijiahao314/go-qa/api"
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/middleware"
)

type ChatWSRouter struct{}

func (cr *ChatWSRouter) Register(rg *gin.RouterGroup) {
	r := rg.Group("/chat")
	ChatWSApi := new(api.ChatWSApi)

	r.Use(middleware.Auth(), middleware.Role([]string{global.ROLE_ADMIN, global.ROLE_USER}))

	r.GET("/ws", ChatWSApi.ChatWebSocket)
	// WebSocket
}
