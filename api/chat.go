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
	UserID uint64 `json:"userid,string"`
	Title  string `json:"title"`
}

type AddChatInfoResponse struct {
	BaseResponse
}

// 新增ChatInfo
func (ca *ChatApi) AddChatInfo(c *gin.Context) {
	// request
	var req AddChatInfoRequest
	// response
	res := AddChatInfoResponse{}
	if err := c.ShouldBindJSON(&req); err != nil {
		res.Code = http.StatusBadRequest
		res.Msg = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	// data
	chatInfo := model.ChatInfo{}
	chatInfo.Title = req.Title
	chatInfo.UserID = req.UserID
	// service
	as := new(service.AuthService)
	if ok, err := as.UserIDInDatabase(req.UserID); !ok {
		res.Code = http.StatusBadRequest
		res.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	cs := new(service.ChatService)
	if err := cs.AddChatInfo(chatInfo); err != nil {
		res.Code = http.StatusBadRequest
		res.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	// success
	res.BaseResponse = BaseResponse{
		Code: http.StatusOK,
		Msg:  "success",
	}
	c.JSON(http.StatusOK, res)
}

type DeleteChatInfoRequest struct {
}

type DeleteChatInfoResponse struct {
	BaseResponse
}

func (ca *ChatApi) DeleteChatInfo(c *gin.Context) {
	// response
	res := DeleteChatInfoResponse{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Msg = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	// service
	cs := new(service.ChatService)
	if err := cs.DeleteChatInfo(uint(id)); err != nil {
		res.Code = http.StatusBadRequest
		res.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	res.Code = http.StatusBadRequest
	res.Msg = "success"
	c.JSON(http.StatusOK, res)
}

type UpdateChatInfoRequest struct {
	Title string `json:"title"`
}

type UpdateChatInfoResponse struct {
	BaseResponse
}

func (ca *ChatApi) UpdateChatInfo(c *gin.Context) {
	// request
	var req UpdateChatInfoRequest
	// response
	res := UpdateChatInfoResponse{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Msg = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		res.Code = http.StatusBadRequest
		res.Msg = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	// data
	chatInfo := model.ChatInfo{}
	chatInfo.ID = uint(id)
	chatInfo.Title = req.Title
	// service
	cs := new(service.ChatService)
	if err := cs.UpdateChatInfo(chatInfo); err != nil {
		res.Code = http.StatusBadRequest
		res.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	// success
	res.Code = http.StatusBadRequest
	res.Msg = "success"
	c.JSON(http.StatusOK, res)
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

// Add ChatInfo
type AddChatCardRequest struct {
	ChatCard model.ChatCard `json:"chat_card"`
}

type AddChatCardResponse struct {
	BaseResponse
}

// Add ChatCard
func (ca *ChatApi) AddChatCard(c *gin.Context) {
	// request
	var req AddChatCardRequest
	// response
	res := AddChatCardResponse{}
	// data
	if err := c.ShouldBindJSON(&req); err != nil {
		res.Code = http.StatusBadRequest
		res.Msg = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	chatCard := req.ChatCard
	// service
	cs := new(service.ChatService)
	if err := cs.AddChatCard(chatCard); err != nil {
		res.Code = http.StatusBadRequest
		res.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	// success
	res.Code = http.StatusBadRequest
	res.Msg = "success"
	c.JSON(http.StatusOK, res)
}

type DeleteChatCardRequest struct {
}

type DeleteChatCardResponse struct {
	BaseResponse
}

func (ca *ChatApi) DeleteChatCard(c *gin.Context) {
	// response
	res := DeleteChatCardResponse{}
	// data
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Msg = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	// service
	cs := new(service.ChatService)
	if err := cs.DeleteChatCard(uint(id)); err != nil {
		res.Code = http.StatusBadRequest
		res.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	// success
	res.Code = http.StatusBadRequest
	res.Msg = "success"
	c.JSON(http.StatusOK, res)
}

type UpdateChatCardRequest struct {
}

type UpdateChatCardResponse struct {
}

func (ca *ChatApi) UpdateChatCard(c *gin.Context) {}

type GetChatCardsRequest struct {
	UserID uint64 `json:"userid"`
}

type GetChatCardsResponse struct {
}

// 获取该用户下所有的ChatCard
func (ca *ChatApi) GetChatCards(c *gin.Context) {
}
