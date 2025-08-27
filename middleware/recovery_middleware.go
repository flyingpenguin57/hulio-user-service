package middleware

import (
	"fmt"
	"hulio-user-service/handler/response"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware 统一异常处理
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//defer 定义的函数会在当前函数返回时执行，即使发生Panic也会执行
		//先定义后执行
		defer func() {
			if err := recover(); err != nil {
				// 打印错误日志（你也可以换成 zap / logrus 等）
				fmt.Printf("[PANIC] %v\n%s\n", err, string(debug.Stack()))

				// 返回统一格式
				response.Fail(c, 500, err.(string))
				c.Abort()
			}
		}()

		c.Next()

		// 捕获业务层返回的错误（比如 c.Error(err)）
		if len(c.Errors) > 0 {
			// 取第一个错误
			err := c.Errors[0].Err
			response.Fail(c, 400, err.Error())
			c.Abort()
			return
		}
	}
}
