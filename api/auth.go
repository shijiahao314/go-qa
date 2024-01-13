package api

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shijiahao314/go-qa/errcode"
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/model"
	"github.com/shijiahao314/go-qa/service"
	"go.uber.org/zap"
)

type AuthApi struct{}

// SignUp
func (aa *AuthApi) SignUp(c *gin.Context) {
	type SignUpRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	type SignUpResponse struct {
		BaseResponse
	}
	req := SignUpRequest{}
	resp := SignUpResponse{}
	// param
	if err := c.ShouldBindJSON(&req); err != nil {
		global.Logger.Info("invalid request", zap.Error(err))
		resp.Code = errcode.InvalidRequest
		resp.Msg = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	if len(req.Username) < 3 || len(req.Password) < 6 {
		resp.Code = errcode.UsernameTooShort
		resp.Msg = "username or password should not less than 6 chatacters"
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	// service
	u := model.User{
		Username: req.Username,
		Password: req.Password,
		Role:     model.UserRoleUser,
	}
	us := new(service.UserService)
	if err := us.AddUser(u); err != nil {
		global.Logger.Info("failed to add user", zap.Error(err))
		resp.Code = errcode.AddUserFailed
		resp.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	// success
	global.Logger.Info("success add user", zap.String("username", u.Username))
	resp.Code = 0
	resp.Msg = "success"
	c.JSON(http.StatusOK, resp)
}

// Login
func (aa *AuthApi) Login(c *gin.Context) {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	type LoginResponse struct {
		BaseResponse
	}
	req := LoginRequest{}
	res := LoginResponse{}
	// param
	if err := c.ShouldBindJSON(&req); err != nil {
		global.Logger.Info("invalid request", zap.Error(err))
		res.Code = errcode.InvalidRequest
		res.Msg = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	// service
	as := new(service.AuthService)
	user, err := as.Login(req.Username, req.Password)
	if err != nil {
		global.Logger.Info("login failed", zap.String("username", req.Username), zap.Error(err))
		res.Code = errcode.LoginFailed
		res.Msg = err.Error()
		c.JSON(http.StatusUnauthorized, res)
		return
	}
	// session
	session := sessions.Default(c)
	if userInfo := session.Get(global.USER_INFO_KEY); userInfo != nil {
		global.Logger.Info("already login", zap.String("username", req.Username))
		res.Code = 0
		res.Msg = "already login"
		c.JSON(http.StatusOK, res)
		return
	}
	userInfo := model.UserDTO{
		UserID:   user.UserID,
		Username: user.Username,
		Role:     user.Role,
	}
	session.Set(global.USER_INFO_KEY, userInfo)
	if err := session.Save(); err != nil {
		global.Logger.Info("failed to save session", zap.Error(err))
		res.Code = errcode.SessionSaveFailed
		res.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	// success
	res.Code = 0
	res.Msg = "success"
	c.JSON(http.StatusOK, res)
}

// Logout
func (aa *AuthApi) Logout(c *gin.Context) {
	type LogoutRequest struct {
	}
	type LogoutResponse struct {
		BaseResponse
	}
	// req := LogoutRequest{}
	res := LogoutResponse{}
	// session
	session := sessions.Default(c)
	if userInfo := session.Get(global.USER_INFO_KEY); userInfo == nil {
		res.Code = 0
		res.Msg = "not login"
		c.JSON(http.StatusUnauthorized, res)
		return
	}
	// success
	session.Clear()
	session.Options(sessions.Options{MaxAge: -1}) // clear client cookie
	session.Save()
	res.Code = 0
	res.Msg = "success"
	c.JSON(http.StatusOK, res)
}

// IsLogin
func (aa *AuthApi) IsLogin(c *gin.Context) {
	type IsLoginRequest struct {
	}
	type IsLoginResponse struct {
		BaseResponse
		Username string `json:"username"`
	}
	// req := IsLoginRequest{}
	res := IsLoginResponse{}
	// session
	session := sessions.Default(c)
	uInfo := session.Get(global.USER_INFO_KEY)
	if uInfo == nil {
		res.Code = errcode.NotLogin
		res.Msg = "not login"
		c.JSON(http.StatusUnauthorized, res)
		return
	}
	userInfo := uInfo.(map[string]interface{})
	// success
	res.Code = 0
	res.Msg = "is login"
	res.Username = userInfo[global.USER_USERNAME_KEY].(string)
	c.JSON(http.StatusOK, res)
}
