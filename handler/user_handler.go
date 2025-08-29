package handler

import (
	"errors"
	"hulio-user-service/handler/request"
	"hulio-user-service/handler/response"
	"hulio-user-service/service"
	"hulio-user-service/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	r.POST("/api/v1/user/login", Login)
	r.POST("/api/v1/user/register", Register)
	r.GET("/api/v1/user", GetUserInfo)
	r.DELETE("/api/v1/user", DeleteUser)
	r.PUT("/api/v1/user", UpdateUser)
	r.GET("/api/v1/mock/panic", MockPanic) // 用于测试中间件的 panic 处理
	r.GET("/health", HealthCheck)          // 健康检查端点
}

func Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}
	if err := service.Register(&req); err != nil {
		c.Error(err)
		return
	}
	response.Success(c, nil)
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
	if res, err := service.GetUserInfo(nil, uc); err != nil {
		c.Error(err)
		return
	} else {
		response.Success(c, res)
	}

}

func DeleteUser(c *gin.Context) {
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

	if err := service.DeleteUser(uc); err != nil {
		c.Error(err)
		return
	}
	response.Success(c, nil)
}

func UpdateUser(c *gin.Context) {
	var req request.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		return
	}
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
	if res, err := service.UpdateUser(&req, uc); err != nil {
		c.Error(err)
		return
	} else {
		response.Success(c, res)
	}
}

func MockPanic(c *gin.Context) {
	// 模拟一个 panic
	panic("this is a test panic")
}

func HealthCheck(c *gin.Context) {
	// 健康检查端点，返回简单的状态信息
	c.JSON(200, gin.H{
		"status":    "ok",
		"service":   "hulio-user-service",
		"timestamp": time.Now().Unix(),
	})
}
