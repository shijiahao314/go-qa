package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/service"
	"github.com/shijiahao314/go-qa/utils"
	"go.uber.org/zap"
)

type ChatWSApi struct{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type MsgType int

const (
	GetMsgs MsgType = 1
)

type WSChatReceiveMessage struct {
	Type       MsgType `json:"type"`
	ChatInfoID string  `json:"chat_info_id"`
	Content    string  `json:"content"`
}

func (rm *WSChatReceiveMessage) String() string {
	b, _ := json.Marshal(rm)
	return string(b)
}

type WSChatSendMessage struct {
	Content string `json:"content"`
}

// WebSocket
func (ca *ChatWSApi) ChatWebSocket(c *gin.Context) {
	//
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		global.Logger.Error("failed to upgrade websocket", zap.Error(err))
	}
	ticker := time.NewTicker(pingPeriod)
	// defer close
	defer func() {
		conn.Close()
		ticker.Stop()
	}()
	go func() {
		<-c.Done()
		global.Logger.Info("ws lost connection")
	}()

	for {
		var rcvMsg WSChatReceiveMessage
		err := conn.ReadJSON(&rcvMsg)
		if err != nil {
			global.Logger.Error("failed to read message", zap.Error(err))
			break
		}
		fmt.Printf("receive: %+v\n", rcvMsg)
		if rcvMsg.Type == GetMsgs {
			// GetMsgs
			cs := new(service.ChatService)
			chatInfoID, err := utils.StringToUint(rcvMsg.ChatInfoID)
			if err != nil {
				global.Logger.Error("invalid request", zap.Error(err))
				return
			}
			chatCards, err := cs.GetChatCards(chatInfoID)
			if err != nil {
				global.Logger.Error("failed to get chatcards", zap.Error(err))
				return
			}
			for _, chatCard := range chatCards {
				fmt.Printf("send: %+v\n", chatCard)
				err = conn.WriteJSON(chatCard)
				if err != nil {
					global.Logger.Error("failed to write message", zap.Error(err))
					break
				}
			}
		} else {
			sendMsg := WSChatSendMessage{
				Content: rcvMsg.Content,
			}
			err = conn.WriteJSON(sendMsg)
			if err != nil {
				global.Logger.Error("failed to write message", zap.Error(err))
				break
			}
			fmt.Printf("reply: %s\n", sendMsg)
		}
	}
}
