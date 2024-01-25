package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shijiahao314/go-qa/errcode"
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/model"
	"github.com/shijiahao314/go-qa/service"
	"github.com/shijiahao314/go-qa/utils"
)

type SettingApi struct {
}

func (ca *SettingApi) Register(rg *gin.RouterGroup) {
	r := rg.Group("/setting")
	// Setting
	r.POST("/setting", ca.UpdateSetting)
}

// Setting
// UpdateSetting
func (ca *SettingApi) UpdateSetting(c *gin.Context) {
	type UpdateSettingRequest struct {
		model.UserSetting
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
	if err := ss.UpdateSetting(userid, req.UserSetting); err != nil {
		resp.Code = errcode.UpdateSettingFailed
		resp.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
}
