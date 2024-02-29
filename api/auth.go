package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shijiahao314/go-qa/errcode"
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/model"
	"github.com/shijiahao314/go-qa/service"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type AuthApi struct{}

func (aa *AuthApi) Register(rg *gin.RouterGroup) {
	r := rg.Group("/auth")
	// SignUp
	r.POST("/signup", aa.SignUp)
	// Login
	r.POST("/login", aa.Login)
	// Logout
	r.POST("/logout", aa.Logout)
	// IsLogin
	r.GET("/islogin", aa.IsLogin)
	// oauth/github
	r.GET("/oauth/github", aa.HandleGithubCallback)
}

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
		AccountType: model.AccountTypeBase,
		Username:    req.Username,
		Password:    req.Password,
		Role:        model.UserRoleUser,
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

// Github Login
func (aa *AuthApi) HandleGithubCallback(c *gin.Context) {
	type GithubLoginRequest struct {
	}
	type GithubLoginResponse struct {
		BaseResponse
	}
	resp := GithubLoginResponse{}
	// GET https://github.com/login/oauth/authorize
	conf := &oauth2.Config{
		ClientID:     global.Config.OAuthConfig.Github.ClientID,
		ClientSecret: global.Config.OAuthConfig.Github.ClientSecret,
		Scopes:       []string{},
		Endpoint:     github.Endpoint,
	}
	// 查询code
	code := c.Query("code")
	// 使用code换取token
	token, err := conf.Exchange(c, code)
	if err != nil {
		resp.Code = errcode.InternalServerError
		resp.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	// 使用token获取用户信息
	client := conf.Client(c, token)
	user, err := client.Get("https://api.github.com/user")
	if err != nil {
		resp.Code = errcode.InternalServerError
		resp.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	// defer user.Body.Close()
	data, _ := io.ReadAll(user.Body)
	userInfo := &model.GithubUserDTO{}
	if err := json.Unmarshal(data, userInfo); err != nil {
		resp.Code = errcode.InternalServerError
		resp.Msg = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	fmt.Printf("Github UserInfo: \n%v\n", userInfo)
	resp.Code = 0
	resp.Msg = "success"
	c.JSON(http.StatusOK, resp)
}
