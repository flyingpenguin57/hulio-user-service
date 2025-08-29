package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	// 生成 RSA 私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// 保存私钥到文件
	privateFile, err := os.Create("./test/private_key.pem")
	if err != nil {
		panic(err)
	}
	defer privateFile.Close()

	privateBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	if err := pem.Encode(privateFile, privateBlock); err != nil {
		panic(err)
	}

	// 生成并保存公钥到文件
	publicKey := &privateKey.PublicKey
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}

	publicFile, err := os.Create("./test/public_key.pem")
	if err != nil {
		panic(err)
	}
	defer publicFile.Close()

	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	if err := pem.Encode(publicFile, publicBlock); err != nil {
		panic(err)
	}

	fmt.Println("✅ 公钥和私钥已成功生成！")
}
