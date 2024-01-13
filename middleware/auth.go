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
			username, ok := userInfo[global.USER_USERNAME_KEY].(string)
			if !ok {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code": errcode.InternalServerError,
					"msg":  fmt.Sprintf("key '%s' not in session", global.USER_USERNAME_KEY),
				})
				return
			}
			roles, err := global.Enforcer.GetRolesForUser(username)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code": errcode.InternalServerError,
					"msg":  err.Error(),
				})
				return
			}
			if len(roles) == 0 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code": errcode.NoRoleExist,
					"msg":  "no role exists",
				})
				return
			}
			sub := username
			obj := c.Request.URL.Path
			// act := c.Request.Method
			ok, err = global.Enforcer.Enforce(sub, obj)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code": errcode.InternalServerError,
					"msg":  err.Error(),
				})
				return
			}
			if !ok {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code": errcode.Unauthorized,
					"msg":  "unauthorized",
				})
				return
			}
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
