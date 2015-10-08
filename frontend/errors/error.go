package errors

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("item not found")

const (
	ErrCodeOK                  = 0
	ErrCodeInternalServerError = 10000 // 内部服务器出错
	ErrCodeAuthTypeMissing     = 20000 // auth_type 缺失
	ErrCodeAuthTypeUnknown     = 20001 // auth_type 无法识别
	ErrCodeSessionTokenEncode  = 20002 // sessiontoken 编码出错
)

var ErrOK = &Error{
	ErrCode: ErrCodeOK,
	ErrMsg:  "ok",
}
var ErrInternalServerError = &Error{
	ErrCode: ErrCodeInternalServerError,
	ErrMsg:  "internal server error ",
}
var ErrAuthTypeMissing = &Error{
	ErrCode: ErrCodeAuthTypeMissing,
	ErrMsg:  "auth_type missing",
}
var ErrAuthTypeUnknown = &Error{
	ErrCode: ErrCodeAuthTypeUnknown,
	ErrMsg:  "auth_type unknown",
}
var ErrSessionTokenEncode = &Error{
	ErrCode: ErrCodeSessionTokenEncode,
	ErrMsg:  "sessiontoken encoding failure",
}

type Error struct {
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

func NewError(ErrCode int, ErrMsg string) *Error {
	return &Error{
		ErrCode: ErrCode,
		ErrMsg:  ErrMsg,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("err_code: %d, err_msg: %s", e.ErrCode, e.ErrMsg)
}