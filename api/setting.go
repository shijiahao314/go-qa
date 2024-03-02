package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shijiahao314/go-qa/errcode"
	"github.com/shijiahao314/go-qa/errmsg"
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/model"
	"github.com/shijiahao314/go-qa/service"
	"github.com/shijiahao314/go-qa/utils"
)

type SettingApi struct {
}

func (ca *SettingApi) Register(rg *gin.RouterGroup) {
	r := rg.Group("/settings")
	// Setting
	r.POST("/settings", ca.UpdateSetting)
	r.GET("/settings", ca.GetSetting)
}

func (ca *SettingApi) checkChatModel(chatModel model.ChatModel) error {
	switch chatModel {
	case "gpt-3.5-turbo":
		return nil
	default:
		return fmt.Errorf("unknown model '%s'", chatModel)
	}
}

// Setting
// UpdateSetting
func (ca *SettingApi) UpdateSetting(c *gin.Context) {
	type UpdateSettingRequest struct {
		model.UserSettingDTO
	}
	type UpdateSettingResponse struct {
		BaseResponse
	}
	req := UpdateSettingRequest{}
	resp := UpdateSettingResponse{}
	// param
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.Code = errcode.InvalidRequest
		resp.Msg = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	// check model
	if err := ca.checkChatModel(req.ChatModel); err != nil {
		resp.Code = errcode.ChatModelNotExists
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
	// service
	ss := new(service.SettingService)
	if err := ss.UpdateSetting(userid, req.UserSettingDTO); err != nil {
		resp.Code = errcode.UpdateSettingFailed
		resp.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	// success
	resp.Code = errcode.Success
	resp.Msg = errmsg.Success
	c.JSON(http.StatusOK, resp)
}

// GetSetting
func (ca *SettingApi) GetSetting(c *gin.Context) {
	type GetSettingRequest struct{}
	type GetSettingResponse struct {
		BaseResponse
		UserSettingDTO model.UserSettingDTO `json:"setting"`
	}
	// req := GetSettingRequest{}
	resp := GetSettingResponse{}
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
	ss := new(service.SettingService)
	setting, err := ss.GetSetting(userid)
	if err != nil {
		resp.Code = errcode.GetSettingFailed
		resp.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	// success
	resp.Code = errcode.Success
	resp.Msg = errmsg.Success
	resp.UserSettingDTO.OpenaiApiKey = setting.OpenaiApiKey
	resp.UserSettingDTO.ChatModel = setting.ChatModel
	resp.UserSettingDTO.TestMode = setting.TestMode
	c.JSON(http.StatusOK, resp)
}
