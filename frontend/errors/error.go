package errors

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("item not found")

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
