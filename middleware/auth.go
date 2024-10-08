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
		if global.Mode == global.TEST {
			// TEST 模式默认通过
			c.Next()
			return
		}

		// session 鉴权
		session := sessions.Default(c)
		if uInfo := session.Get(global.UserInfoKey); uInfo != nil {
			userInfo := uInfo.(map[string]any)
			username, ok := userInfo[global.UserUsernameKey].(string)
			if !ok {
				// key 不存在
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code": errcode.InternalServerError,
					"msg":  fmt.Sprintf("key '%s' not in session", global.UserUsernameKey),
				})
				return
			}

			// casbin 鉴权
			roles, err := global.Enforcer.GetRolesForUser(username)
			if err != nil {
				// 获取角色失败
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code": errcode.InternalServerError,
					"msg":  err.Error(),
				})
				return
			}
			if len(roles) == 0 {
				// 没有角色
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
				// casbin 鉴权失败
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code": errcode.InternalServerError,
					"msg":  err.Error(),
				})
				return
			}
			if !ok {
				// 没有对应资源的权限
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code": errcode.Unauthorized,
					"msg":  "unauthorized",
				})
				return
			}

			// 设置用户信息
			for k := range userInfo {
				c.Set(k, userInfo[k])
			}
			c.Next()
			return
		}

		// 未登录
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code": errcode.NotLogin,
			"msg":  "not login",
		})
	}
}
