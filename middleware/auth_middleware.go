package middleware

import (
	constant "hulio-user-service/constants"
	"hulio-user-service/constants/bizerror"
	"hulio-user-service/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get request path
		path := c.FullPath()
		// if whitelist contains path
		if slices.Contains(constant.ApiWhiteList, path) {
			c.Next()
			return
		}

		// get token
		token := c.GetHeader("Authorization")
		if token == "" {
			c.Error(bizerror.NoAuthHeader)
			c.Abort()
			return
		}

		if claims, err := utils.ParseToken(token); err != nil {
			c.Error(err)
			c.Abort()
			return
		} else {
			// 将解析后的 claims 存储到上下文中，供后续处理使用
			c.Set("claims", claims)
		}

		// 如果认证通过，继续处理请求
		c.Next()
	}
}
