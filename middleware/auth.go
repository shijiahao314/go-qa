package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shijiahao314/go-qa/errcode"
	"github.com/shijiahao314/go-qa/global"
	"github.com/shijiahao314/go-qa/helper"
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

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		mode := helper.GetMode()
		// test mode
		if mode == global.TEST {
			c.Next()
			return
		}
		session := sessions.Default(c)
		if uInfo := session.Get(global.USER_INFO_KEY); uInfo != nil {
			userInfo := uInfo.(map[string]interface{})
			// set user info
			for k := range userInfo {
				c.Set(k, userInfo[k])
			}
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": errcode.NotLogin,
			"msg":  "not login",
		})
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
