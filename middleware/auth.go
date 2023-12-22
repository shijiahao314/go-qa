package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shijiahao314/go-qa/errcode"
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/helper"
	"go.uber.org/zap"
)

func debugUserinfo(s sessions.Session) (map[string]interface{}, error) {
	uinfo := make(map[string]interface{}, 0)
	uinfo["id"] = 1
	uinfo["username"] = "admin"
	uinfo["role"] = "admin"

	s.Set(global.USER_INFO_KEY, uinfo)
	err := s.Save()
	return uinfo, fmt.Errorf("debugUserinfo failed to save session: %w", err)
}

func saveUserInfo(uinfo map[string]interface{}, c *gin.Context) {
	c.Set("username", uinfo["username"])
	c.Set("ID", uinfo["id"])
	c.Set("role", uinfo["role"])
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uinfo, ok := session.Get(global.USER_INFO_KEY).(map[string]interface{})
		mode := helper.GetMode()

		// if debug mode on
		if debug := c.Query("debug"); !ok && len(debug) != 0 && (mode == global.DEV || mode == global.TEST) {
			userinfo, err := debugUserinfo(session)
			if err != nil {
				global.Logger.Error("failed to debugUserinfo", zap.Error(err))
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code": errcode.SessionSave,
					"msg":  err.Error(),
				})
				return
			}
			saveUserInfo(userinfo, c)
			c.Next()
			return
		}

		// or just return unauthorized
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": errcode.NotLogin,
				"msg":  "not login",
			})
			return
		}

		saveUserInfo(uinfo, c)
		c.Next()
	}
}

// func UserExist() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		userId, ok := c.Get("UserID").(uint64)
// 		if !ok {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"code": errcode.NotLogin,
// 				"msg":  "not login",
// 			})
// 			return
// 		}

// 		if !UserIDInDatabase() {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
// 				"code": errcode.NotLogin,
// 				"msg":  "not login",
// 			})
// 			return
// 		}
// 		c.Next()
// 	}
// }
