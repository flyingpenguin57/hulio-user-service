package bizerror

import (
	"fmt"
)

// 定义一个结构体来表示自定义错误
type BizError struct {
	Code    int
	Message string
}

// 实现 error 接口（必须实现 Error() string 方法）
func (e *BizError) Error() string {
	return fmt.Sprintf("biz error %d: %s", e.Code, e.Message)
}

var NoAuthHeader = &BizError{
	Code:    001,
	Message: "miss athorization header",
}

var UserNotExist = &BizError{Code: 001, Message: "user not exist"}

var PasswordError = &BizError{Code: 002, Message: "password error"}

var UsernameExisted = &BizError{Code: 003, Message: "username existed"}
