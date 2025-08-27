package middleware

import (
	"time"

	"hulio-user-service/config"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RequestRecorder() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		config.Log.Info("HTTP Request - Start",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("client_ip", c.ClientIP()),
		)

		c.Next() // 处理请求

		duration := time.Since(start)

		config.Log.Info("HTTP Request - End",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", duration),
			zap.String("client_ip", c.ClientIP()),
		)
	}
}
