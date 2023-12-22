package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shijiahao314/go-qa/api"
)

type ChatRouter struct{}

func (cr *ChatRouter) Register(rg *gin.RouterGroup) {
	r := rg.Group("/chat")
	ChatApi := new(api.ChatApi)

	// r.Use(middleware.Auth(), middleware.Role([]string{global.ROLE_ADMIN}))

	// ChatInfo
	r.POST("/chatInfo", ChatApi.AddChatInfo)
	r.DELETE("/chatInfo/:id", ChatApi.DeleteChatInfo)
	r.PUT("/chatInfo/:id", ChatApi.UpdateChatInfo)
	r.GET("/chatInfo", ChatApi.GetChatInfos)
	// ChatCard
	r.POST("/chatCard", ChatApi.AddChatCard)
	r.DELETE("/chatCard/:id", ChatApi.DeleteChatCard)
	r.PUT("/chatCard", ChatApi.UpdateChatCard)
	r.GET("/chatCard", ChatApi.GetChatCards)
}
