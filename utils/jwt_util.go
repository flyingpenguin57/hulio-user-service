package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 自定义声明结构体
type UserClaims struct {
	Username string `json:"username"`
	UserId   int64  `json:"userId"`
	jwt.RegisteredClaims
}

// 生成 JWT
func GenerateToken(username string, id int64) (string, error) {
	// 加载 RSA 私钥（用于 RS256 签名）
	privateKey, err := loadRSAPrivateKey()
	if err != nil {
		return "", fmt.Errorf("加载私钥失败: %v", err)
	}
	claims := UserClaims{
		Username: username,
		UserId:   id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)), //expire after 2 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "hulio-user-service",
			Subject:   "user-token",
		},
	}

	// 使用 RS256 算法生成 token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privateKey)
}

// 公钥 PEM 字符串
const RSAPublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA0rS0nswzzdvGy6/IJQHw
9u5RE756m7WnY/AQwR9+2JcQNgIb/CrVEEPDheiN4zte8v1R0lZNmuaZjKsoW+aB
P3l9o0IuFDHyjo2alGcK/5UoPl/hhgTS2ID2mCIfd9s2j3mP75kiCAP2Sgp/VzM6
Umlp20yMVM4qkcoy8wbjoAJdmkUYehT4XCuUiU5MhzZ+OQdo7oubQMSReDimRKpk
oN3Oo81dOt95l+ccGbqZ/7x9VjECWqneAzr3lbLxtQYUEC5189Tl7/9kiI3mnXL2
1sHAtKLdq3wq9XoSTOUOG9xsLpndpuE2fdlOG7DlmREVU0qrMP2yWyM6QKpFAru0
WQIDAQAB
-----END PUBLIC KEY-----`

// 解析和验证 JWT
func ParseToken(tokenStr string) (*UserClaims, error) {
	// 加载 RSA 公钥（用于 RS256 验签）
	publicKey, err := loadRSAPublicKey()
	if err != nil {
		return nil, fmt.Errorf("加载公钥失败: %v", err)
	}

	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法是否一致
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	// 断言 claims 类型
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

func loadRSAPrivateKey() (*rsa.PrivateKey, error) {
	// 从环境变量 JWT_PRIVATE_KEY 读取 PEM 内容
	pemStr := os.Getenv("JWT_PRIVATE_KEY")
	if pemStr == "" {
		return nil, fmt.Errorf("没有找到私钥")
	}
	block, _ := pem.Decode([]byte(pemStr))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("无效的私钥环境变量")
	}
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("解析私钥失败: %v", err)
	}
	return key, nil
}

func loadRSAPublicKey() (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(RSAPublicKey))
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("无效的公钥")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("解析公钥失败: %v", err)
	}
	pubKey, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("公钥类型不是 RSA")
	}
	return pubKey, nil
}
