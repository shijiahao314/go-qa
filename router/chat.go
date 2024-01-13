package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shijiahao314/go-qa/api"
)

type ChatRouter struct{}

func (cr *ChatRouter) Register(rg *gin.RouterGroup) {
	r := rg.Group("/chat")
	ChatApi := new(api.ChatApi)

	// ChatInfo
	r.POST("/chatInfo", ChatApi.AddChatInfo)
	r.DELETE("/chatInfo/:id", ChatApi.DeleteChatInfo) // chatInfoId
	r.PUT("/chatInfo/:id", ChatApi.UpdateChatInfo)    // chatInfoId
	r.GET("/chatInfos", ChatApi.GetChatInfos)
	r.GET("/chatInfo/:id", ChatApi.GetChatInfo) // chatInfoId
	// ChatCard
	r.POST("/chatCard", ChatApi.AddChatCard)
	r.DELETE("/chatCard/:id", ChatApi.DeleteChatCard) // chatCardId
	r.PUT("/chatCard/:id", ChatApi.UpdateChatCard)    // chatCardId
	r.GET("/chatCards/:id", ChatApi.GetChatCards)     // chatInfoId
	r.GET("/chatCard/:id", ChatApi.GetChatCard)       // chatCardId
}
