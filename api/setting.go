package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shijiahao314/go-qa/errcode"
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/model"
	"github.com/shijiahao314/go-qa/utils"
)

type SettingApi struct {
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
	}
}
