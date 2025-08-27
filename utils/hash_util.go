package utils
import (
    "golang.org/x/crypto/bcrypt"
)

//Go 官方推荐使用 "golang.org/x/crypto/bcrypt" 它会自动生成随机salt并加密
func HashPassword(password string) (string, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hash), nil
}

//check password
func CheckPassword(hash, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}