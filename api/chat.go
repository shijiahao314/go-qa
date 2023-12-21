package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shijiahao314/go-qa/errcode"
	"github.com/shijiahao314/go-qa/model"
	"github.com/shijiahao314/go-qa/service"
)

type ChatApi struct{}

// Add ChatInfo
type AddChatInfoRequest struct {
}

type AddChatInfoResponse struct {
	BR BaseResponse
}

// 新增ChatInfo
func (ca *ChatApi) AddChatInfo(c *gin.Context) {
	cs := new(service.ChatService)

	chatInfo := model.ChatInfo{}
	if err := c.ShouldBindJSON(&chatInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errcode.ParamParseFailed,
			"msg":  err.Error(),
		})
		return
	}

	if err := cs.AddChatInfo(chatInfo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": errcode.AddChatInfoFailed,
			"msg":  err.Error(),
		})
		return
	}

}

type GetChatInfosRequest struct {
	UserID uint64 `json:"userid"`
}

type GetChatInfosResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		ChatInfos []model.ChatInfo `json:"chat_infos"`
	} `json:"data"`
}

// 获取该用户下所有的ChatInfo
func (ca *ChatApi) GetChatInfos(c *gin.Context) {
	cs := new(service.ChatService)

	userid_string, ok := c.GetQuery("userid")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errcode.ParamParseFailed,
			"msg":  "failed to parse userid",
		})
		return
	}
	userid, err := strconv.ParseUint(userid_string, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": errcode.ParamParseFailed,
			"msg":  "failed to parse userid_string to userid",
		})
		return
	}

	chatInfos, err := cs.GetChatInfos(userid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": errcode.GetChatInfosFailed,
			"msg":  err.Error(),
		})
		return
	}

	res := GetChatInfosResponse{}
	res.Code = http.StatusOK
	res.Data.ChatInfos = chatInfos

	c.JSON(http.StatusOK, res)
}
