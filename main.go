package main

import (
	"hulio-user-service/config"
	"hulio-user-service/handler"
	"hulio-user-service/middleware"
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	//init database
	config.InitDB()
	config.InitLogger()

	r := gin.Default()

	// 配置 CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // 允许访问的前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(middleware.RecoveryMiddleware()) // 使用自定义的异常处理 middleware
	r.Use(middleware.RequestRecorder())    // 使用自定义的日志记录 middleware
	r.Use(middleware.AuthMiddleware())     // 使用自定义的认证 middleware

	//register user routes
	handler.RegisterUserRoutes(r)

	r.Run(":8080")
}
