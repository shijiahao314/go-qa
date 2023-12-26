package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/shijiahao314/go-qa/errcode"
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/model"
	"github.com/shijiahao314/go-qa/service"
	"github.com/shijiahao314/go-qa/utils"
	"go.uber.org/zap"
)

type AdminApi struct{}

type GetUsersResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Page  int              `json:"page"`
		Size  int              `json:"size"`
		Total int64            `json:"total"`
		Users []model.UserInfo `json:"users"`
	} `json:"data"`
}

func (aa *AdminApi) AddUser(c *gin.Context) {
	type AddUserRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}
	type AddUserResponse struct {
		BaseResponse
	}
	req := AddUserRequest{}
	res := AddUserResponse{}
	// param
	if err := c.ShouldBindJSON(&req); err != nil {
		global.Logger.Info("invalid request", zap.Error(err))
		res.Code = errcode.InvalidRequest
		res.Msg = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	newUser := model.User{}
	newUser.Username = req.Username
	newUser.Password = req.Password
	newUser.Role = req.Role
	// service
	us := new(service.UserService)
	ok, err := us.UserExists(newUser.Username)
	if err != nil {
		global.Logger.Info("failed to check user exists", zap.Error(err))
		res.Code = errcode.InternalServerError
		res.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	if ok {
		global.Logger.Info("user already exists", zap.String("username", newUser.Username))
		res.Code = errcode.UserExists
		res.Msg = fmt.Sprintf("user '%s' already exists", newUser.Username)
		c.JSON(http.StatusConflict, res)
		return
	}
	if err := us.AddUser(newUser); err != nil {
		global.Logger.Info("failed to add user", zap.Error(err))
		res.Code = errcode.AddUserFailed
		res.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	// success
	res.Code = 0
	res.Msg = "success"
	c.JSON(http.StatusOK, res)
}

func (aa *AdminApi) DeleteUser(c *gin.Context) {
	type DeleteUserRequest struct {
	}
	type DeleteUserResponse struct {
		BaseResponse
	}
	// req := DeleteUserRequest{}
	res := DeleteUserResponse{}
	// param
	id, err := utils.StringToUint64(c.Param("id"))
	if err != nil {
		global.Logger.Info("invalid request", zap.Error(err))
		res.Code = errcode.InvalidRequest
		res.Msg = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	// service
	us := new(service.UserService)
	if err := us.DeleteUser(id); err != nil {
		global.Logger.Info("failed to delete user", zap.Error(err))
		res.Code = errcode.DeleteUserFailed
		res.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	// success
	global.Logger.Info("success delete user", zap.Uint64("id", id))
	res.Code = 0
	res.Msg = "success"
	c.JSON(http.StatusOK, res)
}

func (aa *AdminApi) UpdateUser(c *gin.Context) {
	type UpdateUserRequest struct {
		// Username string `json:"username"` // 不允许修改用户名
		Password string `json:"password"`
		Role     string `json:"role"`
	}
	type UpdateUserResponse struct {
		BaseResponse
	}
	req := UpdateUserRequest{}
	res := UpdateUserResponse{}
	// param
	id, err := utils.StringToUint64(c.Param("id"))
	if err != nil {
		global.Logger.Info("invalid request", zap.Error(err))
		res.Code = errcode.InvalidRequest
		res.Msg = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	// 查询当前用户权限
	as := new(service.AuthService)
	ok, err := as.UserHasPermission(id)
	if err != nil {
		global.Logger.Info("update user failed", zap.Error(err))
		res.Code = errcode.UpdateUserFailed
		res.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	if !ok {
		global.Logger.Info("permission denied", zap.Error(err))
		res.Code = errcode.PermissionDenied
		res.Msg = "permission denied"
		c.JSON(http.StatusForbidden, res) // 403
		return
	}
	// json
	if err := c.ShouldBindJSON(&req); err != nil {
		global.Logger.Info("invalid request", zap.Error(err))
		res.Code = errcode.UpdateUserFailed
		res.Msg = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	// 查询（不存在则新增，保证幂等性）
	updatedUser := model.User{}
	updatedUser.Password = req.Password
	updatedUser.Role = req.Role
	us := new(service.UserService)
	if err := us.UpdateUser(updatedUser); err != nil {
		global.Logger.Info("update user failed", zap.Error(err))
		res.Code = errcode.UpdateUserFailed
		res.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	// success
	global.Logger.Info("success update user", zap.Uint64("userid", id))
	res.Code = 0
	res.Msg = "success"
	c.JSON(http.StatusOK, res)
}

func (aa *AdminApi) GetUsers(c *gin.Context) {
	type GetUsersRequest struct {
	}
	type GetUsersResponse struct {
		BaseResponse
		Data struct {
			Page      int              `json:"page"`
			Size      int              `json:"size"`
			Total     int64            `json:"total"`
			UserInfos []model.UserInfo `json:"userInfos"`
		}
	}
	// req := GetUsersRequest{}
	res := GetUsersResponse{}
	// param
	page, size := getPageAndSize(c)
	// service
	us := new(service.UserService)
	users, total, err := us.GetUsers(page, size)
	if err != nil {
		global.Logger.Info("failed to get users", zap.Error(err))
		res.Code = errcode.GetUsersFailed
		res.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	// success
	global.Logger.Info("success get users")
	res.Code = 0
	res.Msg = "success"
	res.Data.Page = page
	res.Data.Size = size
	res.Data.Total = total
	var usersInfos []model.UserInfo
	for _, u := range users {
		usersInfos = append(usersInfos, model.UserInfo{
			UserID:   u.UserID,
			Username: u.Username,
			Role:     u.Role,
		})
	}
	res.Data.UserInfos = usersInfos
	c.JSON(http.StatusOK, res)
}
