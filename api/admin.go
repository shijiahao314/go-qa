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
	"go.uber.org/zap"
)

type AdminAPI struct{}

func (aa *AdminAPI) Register(rg *gin.RouterGroup) {
	r := rg.Group("/admin")
	// User
	r.GET("/user", aa.GetUsers)
	r.POST("/user", aa.AddUser)
	r.POST("/user/:id", aa.UpdateUser)
	r.DELETE("/user/:id", aa.DeleteUser)
}

func (aa *AdminAPI) AddUser(c *gin.Context) {
	type AddUserRequest struct {
		Username string         `json:"username"`
		Password string         `json:"password"`
		Role     model.UserRole `json:"role"`
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
	newUser := model.User{
		AccountType: model.AccountTypeBase,
		Username:    req.Username,
		Password:    req.Password,
		Role:        req.Role,
	}
	// service
	us := new(service.UserService)
	ok, err := us.UsernameExists(newUser.Username)
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
	res.Code = errcode.Success
	res.Msg = errmsg.Success
	c.JSON(http.StatusOK, res)
}

func (aa *AdminAPI) DeleteUser(c *gin.Context) {
	type DeleteUserResponse struct {
		BaseResponse
	}
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
	res.Code = errcode.Success
	res.Msg = errmsg.Success
	c.JSON(http.StatusOK, res)
}

func (aa *AdminAPI) UpdateUser(c *gin.Context) {
	type UpdateUserRequest struct {
		// Username string `json:"username"` // 不允许修改用户名
		Password string         `json:"password"`
		Role     model.UserRole `json:"role"`
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
	res.Code = errcode.Success
	res.Msg = errmsg.Success
	c.JSON(http.StatusOK, res)
}

func (aa *AdminAPI) GetUsers(c *gin.Context) {
	type GetUsersResponse struct {
		BaseResponse
		Data struct {
			Total     int64           `json:"total"`
			Size      int             `json:"size"`
			Page      int             `json:"page"`
			UserInfos []model.UserDTO `json:"user_infos"`
		} `json:"data"`
	}
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
	res.Code = errcode.Success
	res.Msg = errmsg.Success
	res.Data.Page = page
	res.Data.Size = size
	res.Data.Total = total
	var usersInfos []model.UserDTO
	for i := range users {
		usersInfos = append(usersInfos, model.UserDTO{
			UserID:   users[i].UserID,
			Username: users[i].Username,
			Role:     users[i].Role,
		})
	}
	res.Data.UserInfos = usersInfos
	c.JSON(http.StatusOK, res)
}
