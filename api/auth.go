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

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// signup
func (aa *AuthApi) SignUp(c *gin.Context) {
	type SignUpReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"Role"`
	}

	req := SignUpReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": errcode.SignupFailed,
			"msg":  err.Error(),
		})
		global.Logger.Info("invalid request", zap.Error(err))
		return
	}

	u := model.User{
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
	}
	us := new(service.UserService)
	if err := us.AddUser(u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": errcode.AddUserFailed,
			"msg":  err.Error(),
		})
		global.Logger.Info("failed to add user", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
	global.Logger.Info("success add user", zap.String("username", u.Username))
}

// login
func (aa *AuthApi) Login(c *gin.Context) {
	session := sessions.Default(c)
	if uinfo, ok := session.Get(global.USER_INFO_KEY).(model.UserInfo); ok {
		global.Logger.Info("already login", zap.String("username", uinfo.Username))
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "success",
		})
		return
	}

	req := LoginReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code": errcode.LoginFailed,
			"msg":  err.Error(),
		})
		global.Logger.Info("invalid request", zap.Error(err))
		return
	}

	as := new(service.AuthService)
	user, err := as.Login(req.Username, req.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": errcode.UsernameOrPwd,
			"msg":  err.Error(),
		})
		global.Logger.Info("invalid login auth", zap.String("username", req.Username), zap.Error(err))
		return
	}

	userInfo := model.UserInfo{
		UserID:   user.UserID,
		Username: user.Username,
		Role:     user.Role,
	}

	session.Set(global.USER_INFO_KEY, userInfo)
	if err := session.Save(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": errcode.SessionSave,
			"msg":  err.Error(),
		})
		global.Logger.Info("failed to save session", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

func (aa *AuthApi) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code": errcode.SessionSave,
			"msg":  err.Error(),
		})
		global.Logger.Info("failed to save session", zap.Error(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
	})
}

func (aa *AuthApi) IsLogin(c *gin.Context) {
	session := sessions.Default(c)
	uinfo, ok := session.Get(global.USER_INFO_KEY).(model.UserInfo)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": errcode.NotLogin,
			"msg":  "not login",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": uinfo,
	})

}
