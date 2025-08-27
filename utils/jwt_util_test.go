package utils

import (
	"fmt"
	"testing"
)

func TestJWT(t *testing.T) {
	
	// 1. 生成 token
	token, err := GenerateToken("qinghao", 10000)
	if err != nil {
		panic(err)
	}
	fmt.Println("生成的 JWT Token:\n", token)

	// 2. 解析 token
	claims, err := ParseToken(token)
	if err != nil {
		fmt.Println("Token 验证失败:", err)
		return
	}
	fmt.Println("解析出的用户名:", claims.Username)
	fmt.Println("解析出的用户ID:", claims.UserId)
	fmt.Println("发行人:", claims.Issuer)
	fmt.Println("过期时间:", claims.ExpiresAt.Time)
}
