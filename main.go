package main

import (
	"hulio-user-service/config"
	"hulio-user-service/handler"
	"hulio-user-service/handler/response"
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
		AllowOrigins: []string{
			// 本地开发环境
			"http://localhost:3000", // Next.js 默认端口
			"http://localhost:5173", // Vite 默认端口
			"http://localhost:8080", // 其他常用端口
			"http://localhost:3001", // 其他常用端口
			"http://localhost:4200", // Angular 默认端口
			// Vercel 托管域名
			"https://hulio-user-service.vercel.app",
			"https://hulio-user-service-git-main.vercel.app",
			"https://hulio-user-service-git-dev.vercel.app",
			"https://hulio-user-service-git-feature.vercel.app",
			// 生产环境域名
			"https://www.hulio88.xyz",
			"https://hulio88.xyz",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(middleware.RecoveryMiddleware()) // 使用自定义的异常处理 middleware
	r.Use(middleware.RequestRecorder())    // 使用自定义的日志记录 middleware
	r.Use(middleware.AuthMiddleware())     // 使用自定义的认证 middleware

	//register user routes
	handler.RegisterUserRoutes(r)

	// standardize 404/405 responses to JSON to keep client decoding stable
	r.NoRoute(func(c *gin.Context) {
		response.Fail(c, 404, "route not found")
	})
	r.NoMethod(func(c *gin.Context) {
		response.Fail(c, 405, "method not allowed")
	})

	r.Run(":8080")
}
