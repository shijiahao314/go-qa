package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/shijiahao314/go-qa/errcode"
	"github.com/shijiahao314/go-qa/global"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		mode := global.Mode
		// test mode
		if mode == global.TEST {
			c.Next()
			return
		}
		session := sessions.Default(c)
		if uInfo := session.Get(global.UserInfoKey); uInfo != nil {
			userInfo := uInfo.(map[string]any)
			username, ok := userInfo[global.UserUsernameKey].(string)
			if !ok {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code": errcode.InternalServerError,
					"msg":  fmt.Sprintf("key '%s' not in session", global.UserUsernameKey),
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
