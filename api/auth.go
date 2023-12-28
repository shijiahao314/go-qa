package api

import (
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shijiahao314/go-qa/errcode"
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/model"
	"github.com/shijiahao314/go-qa/service"
	"go.uber.org/zap"
)

type AuthApi struct{}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (aa *AuthApi) SignUp(c *gin.Context) {
	type SignUpRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}
	type SignUpResponse struct {
		BaseResponse
	}
	req := SignUpRequest{}
	res := SignUpResponse{}
	if err := c.ShouldBindJSON(&req); err != nil {
		global.Logger.Info("invalid request", zap.Error(err))
		res.Code = errcode.SignupFailed
		res.Msg = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	u := model.User{
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
	}
	us := new(service.UserService)
	if err := us.AddUser(u); err != nil {
		global.Logger.Info("failed to add user", zap.Error(err))
		res.Code = errcode.AddUserFailed
		res.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	global.Logger.Info("success add user", zap.String("username", u.Username))
	res.Code = 0
	res.Msg = "success"
	c.JSON(http.StatusOK, res)
}

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
	// request
	if err := c.ShouldBindJSON(&req); err != nil {
		global.Logger.Info("invalid request", zap.Error(err))
		res.Code = errcode.LoginFailed
		res.Msg = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	// service
	as := new(service.AuthService)
	user, err := as.Login(req.Username, req.Password)
	if err != nil {
		global.Logger.Info("invalid login auth", zap.String("username", req.Username), zap.Error(err))
		res.Code = errcode.LoginFailed
		res.Msg = err.Error()
		c.JSON(http.StatusUnauthorized, res)
		return
	}
	// session
	session := sessions.Default(c)
	userid := session.Get(global.USER_USER_ID_KEY)
	if userid != nil && userid == strconv.FormatUint(user.UserID, 10) {
		global.Logger.Info("already login", zap.String("username", req.Username))
		res.Code = 0
		res.Msg = "already login"
		c.JSON(http.StatusOK, res)
		return
	}
	userInfo := model.UserInfo{
		UserID:   user.UserID,
		Username: user.Username,
		Role:     user.Role,
	}
	session.Set(global.USER_USER_ID_KEY, strconv.FormatUint(userInfo.UserID, 10))
	session.Set(global.USER_USERNAME_KEY, userInfo.Username)
	session.Set(global.USER_ROLE_KEY, userInfo.Role)
	if err := session.Save(); err != nil {
		global.Logger.Info("failed to save session", zap.Error(err))
		res.Code = errcode.SessionSave
		res.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	// success
	res.Code = 0
	res.Msg = "success"
	c.JSON(http.StatusOK, res)
}

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
	userid := session.Get(global.USER_USER_ID_KEY)
	if userid == nil {
		res.Code = 0
		res.Msg = "not login"
		c.JSON(http.StatusUnauthorized, res)
		return
	}
	session.Clear()
	session.Options(sessions.Options{MaxAge: -1})
	session.Save()
	res.Code = 0
	res.Msg = "success"
	c.JSON(http.StatusOK, res)
}

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
	userid := session.Get(global.USER_USER_ID_KEY)
	if userid != nil {
		username := session.Get(global.USER_USERNAME_KEY).(string)
		res.Code = 0
		res.Msg = "is login"
		res.Username = username
		c.JSON(http.StatusOK, res)
		return
	}
	// success
	res.Code = errcode.NotLogin
	res.Msg = "not login"
	c.JSON(http.StatusUnauthorized, res)
}
