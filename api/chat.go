package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shijiahao314/go-qa/errcode"
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/model"
	"github.com/shijiahao314/go-qa/service"
	"github.com/shijiahao314/go-qa/utils"
	"go.uber.org/zap"
)

type ChatApi struct{}

// ChatInfo
// AddChatInfo
func (ca *ChatApi) AddChatInfo(c *gin.Context) {
	type AddChatInfoRequest struct {
		Title string `json:"title"`
	}
	type AddChatInfoResponse struct {
		BaseResponse
		Data struct {
			ChatInfo model.ChatInfo `json:"chat_info"`
		} `json:"data"`
	}
	req := AddChatInfoRequest{}
	resp := AddChatInfoResponse{}
	// param
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.Code = errcode.InvalidRequest
		resp.Msg = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	uid := c.GetString(global.USER_USER_ID_KEY)
	userid, err := utils.StringToUint64(uid)
	if err != nil {
		resp.Code = errcode.InternalServerError
		resp.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	// data
	chatInfo := model.ChatInfo{}
	chatInfo.UserID = userid
	chatInfo.Title = req.Title
	// service
	cs := new(service.ChatService)
	if err := cs.AddChatInfo(&chatInfo); err != nil {
		resp.Code = errcode.AddChatInfoFailed
		resp.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	// success
	resp.Code = 0
	resp.Msg = "success"
	resp.Data.ChatInfo = chatInfo
	c.JSON(http.StatusOK, resp)
}

// DeleteChatInfo
func (ca *ChatApi) DeleteChatInfo(c *gin.Context) {
	type DeleteChatInfoRequest struct {
	}
	type DeleteChatInfoResponse struct {
		BaseResponse
	}
	// req := DeleteChatInfoRequest{}
	resp := DeleteChatInfoResponse{}
	// param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp.Code = errcode.InvalidRequest
		resp.Msg = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	// service
	cs := new(service.ChatService)
	if err := cs.DeleteChatInfo(uint(id)); err != nil {
		resp.Code = errcode.InternalServerError
		resp.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	// success
	resp.Code = 0
	resp.Msg = "success"
	c.JSON(http.StatusOK, resp)
}

// UpdateChatInfo
func (ca *ChatApi) UpdateChatInfo(c *gin.Context) {
	type UpdateChatInfoRequest struct {
		Title string `json:"title"`
	}
	type UpdateChatInfoResponse struct {
		BaseResponse
	}
	req := UpdateChatInfoRequest{}
	resp := UpdateChatInfoResponse{}
	// param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp.Code = errcode.InvalidRequest
		resp.Msg = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.Code = errcode.InvalidRequest
		resp.Msg = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	// data
	chatInfo := model.ChatInfo{}
	chatInfo.ID = uint(id)
	chatInfo.Title = req.Title
	// service
	cs := new(service.ChatService)
	if err := cs.UpdateChatInfo(chatInfo); err != nil {
		resp.Code = errcode.InternalServerError
		resp.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	// success
	resp.Code = 0
	resp.Msg = "success"
	c.JSON(http.StatusOK, resp)
}

// GetChatInfos: 获取当前用户所有的ChatInfo
func (ca *ChatApi) GetChatInfos(c *gin.Context) {
	type GetChatInfosRequest struct {
	}
	type GetChatInfosResponse struct {
		BaseResponse
		Data struct {
			ChatInfos []model.ChatInfo `json:"chat_infos"`
		} `json:"data"`
	}
	// req := GetChatInfosRequest{}
	resp := GetChatInfosResponse{}
	// param
	uid := c.GetString(global.USER_USER_ID_KEY)
	userid, err := utils.StringToUint64(uid)
	if err != nil {
		resp.Code = errcode.InternalServerError
		resp.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	// service
	cs := new(service.ChatService)
	chatInfos, err := cs.GetChatInfos(userid)
	if err != nil {
		resp.Code = errcode.GetChatInfosFailed
		resp.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Code = 0
	resp.Msg = "success"
	resp.Data.ChatInfos = chatInfos
	c.JSON(http.StatusOK, resp)
}

// GetChatInfo
func (ca *ChatApi) GetChatInfo(c *gin.Context) {
	type GetChatInfosRequest struct {
	}
	type GetChatInfosResponse struct {
		BaseResponse
		Data struct {
			ChatInfo model.ChatInfo `json:"chat_info"`
		} `json:"data"`
	}
	// req := GetChatInfosRequest{}
	resp := GetChatInfosResponse{}
	// param
	chatInfoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp.Code = errcode.InvalidRequest
		resp.Msg = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	// service
	cs := new(service.ChatService)
	chatInfo, err := cs.GetChatInfo(uint(chatInfoId))
	if err != nil {
		resp.Code = errcode.GetChatInfosFailed
		resp.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Code = 0
	resp.Msg = "success"
	resp.Data.ChatInfo = chatInfo
	c.JSON(http.StatusOK, resp)
}

// ChatCard
// AddChatCard
func (ca *ChatApi) AddChatCard(c *gin.Context) {
	type AddChatCardRequest struct {
		model.ChatCardDTO
	}
	type AddChatCardResponse struct {
		BaseResponse
		ChatCard model.ChatCard `json:"chat_card"`
	}
	req := AddChatCardRequest{}
	resp := AddChatCardResponse{}
	// param
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.Code = errcode.InvalidRequest
		resp.Msg = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	chatCard := model.ChatCard{}
	chatCard.ChatInfoID = req.ChatInfoID
	chatCard.Content = req.Content
	chatCard.Role = req.Role
	// service
	cs := new(service.ChatService)
	if err := cs.AddChatCard(&chatCard); err != nil {
		resp.Code = errcode.InternalServerError
		resp.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	// success
	resp.Code = 0
	resp.Msg = "success"
	resp.ChatCard = chatCard
	c.JSON(http.StatusOK, resp)
}

// DeleteChatCard
func (ca *ChatApi) DeleteChatCard(c *gin.Context) {
	type DeleteChatCardRequest struct {
	}
	type DeleteChatCardResponse struct {
		BaseResponse
	}
	// req := DeleteChatCardRequest{}
	resp := DeleteChatCardResponse{}
	// param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp.Code = errcode.InvalidRequest
		resp.Msg = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	// service
	cs := new(service.ChatService)
	if err := cs.DeleteChatCard(uint(id)); err != nil {
		resp.Code = errcode.InternalServerError
		resp.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	// success
	resp.Code = 0
	resp.Msg = "success"
	c.JSON(http.StatusOK, resp)
}

// UpdateChatCard
func (ca *ChatApi) UpdateChatCard(c *gin.Context) {
	type UpdateChatCardRequest struct {
		model.ChatCardDTO
	}
	type UpdateChatCardResponse struct {
		BaseResponse
	}
	req := UpdateChatCardRequest{}
	resp := UpdateChatCardResponse{}
	// param
	chatId, err := utils.StringToUint(c.Param("id"))
	if err != nil {
		global.Logger.Info("invalid request", zap.Error(err))
		resp.Code = errcode.InvalidRequest
		resp.Msg = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.Code = errcode.InvalidRequest
		resp.Msg = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	chatCard := model.ChatCard{}
	chatCard.ID = uint(chatId)
	chatCard.ChatInfoID = req.ChatCardDTO.ChatInfoID
	chatCard.Content = req.ChatCardDTO.Content
	// service
	cs := new(service.ChatService)
	if err := cs.UpdateChatCard(chatCard); err != nil {
		resp.Code = errcode.InternalServerError
		resp.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	// success
	resp.Code = 0
	resp.Msg = "success"
	c.JSON(http.StatusOK, resp)
}

// GetChatCards
func (ca *ChatApi) GetChatCards(c *gin.Context) {
	type GetChatCardsRequest struct {
	}
	type GetChatCardsResponse struct {
		BaseResponse
		Data struct {
			ChatCards []model.ChatCard `json:"chat_cards"`
		} `json:"data"`
	}
	// req := GetChatCardsRequest{}
	resp := GetChatCardsResponse{}
	// param
	chatInfoId, err := utils.StringToUint(c.Param("id"))
	if err != nil {
		global.Logger.Info("invalid request", zap.Error(err))
		resp.Code = errcode.InvalidRequest
		resp.Msg = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	// service
	cs := new(service.ChatService)
	chatCards, err := cs.GetChatCards(chatInfoId)
	if err != nil {
		global.Logger.Info("failed to get chatcards", zap.Error(err))
		resp.Code = errcode.GetChatCardsFailed
		resp.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	// success
	resp.Code = 0
	resp.Msg = "success"
	resp.Data.ChatCards = chatCards
	c.JSON(http.StatusOK, resp)
}

// GetChatChard
func (ca *ChatApi) GetChatCard(c *gin.Context) {
	type GetChatCardRequest struct {
	}
	type GetChatCardResponse struct {
		BaseResponse
		Data struct {
			ChatCard model.ChatCard `json:"chat_card"`
		} `json:"data"`
	}
	// req := GetChatCardRequest{}
	resp := GetChatCardResponse{}
	// param
	chatId, err := utils.StringToUint(c.Param("id"))
	if err != nil {
		global.Logger.Info("invalid request", zap.Error(err))
		resp.Code = errcode.InvalidRequest
		resp.Msg = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	// service
	cs := new(service.ChatService)
	chatCard, err := cs.GetChatCard(chatId)
	if err != nil {
		global.Logger.Info("failed to get chatcards", zap.Error(err))
		resp.Code = errcode.GetChatCardFailed
		resp.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	// success
	resp.Code = 0
	resp.Msg = "success"
	resp.Data.ChatCard = chatCard
	c.JSON(http.StatusOK, resp)
}
