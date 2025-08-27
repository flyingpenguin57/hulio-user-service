package handler

import (
	"errors"
	"hulio-user-service/handler/request"
	"hulio-user-service/handler/response"
	"hulio-user-service/service"
	"hulio-user-service/utils"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	r.POST("/api/v1/user/login", Login)
	r.POST("/api/v1/user/register", Login)
	r.GET("/users/mock-panic", MockPanic) // 用于测试中间件的 panic 处理
}

func Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}
}

func Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}

	if loginRes, err := service.Login(&req); err != nil {
		c.Error(err)
		return
	} else {
		response.Success(c, loginRes)
	}
}

func GetUserInfo(c *gin.Context) {
    claims, exists := c.Get("claims")
    if !exists {
		c.Error(errors.New("claims not exist"))
        return
    }
    uc, ok := claims.(*utils.UserClaims)
	if !ok {
		c.Error(errors.New("claims type error"))
		return
	}
	service.GetUserInfo(nil, uc)
    
}

func MockPanic(c *gin.Context) {
	// 模拟一个 panic
	panic("this is a test panic")
}
