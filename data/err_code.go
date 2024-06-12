package data

import (
	"errors"
	"fmt"
)

type ErrCode struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	error `json:"err"`
}

func (e ErrCode) Error() string {
	if nil != e.error {
		return fmt.Sprintf("code: %d, msg: %s, error: %v", e.Code, e.Msg, e.error)
	} else {
		return fmt.Sprintf("code: %d, msg: %s", e.Code, e.Msg)
	}

}

func NewErrCode(code int, msg string) *ErrCode {
	return &ErrCode{Code: code, Msg: msg}
}

func IsErrCode(err error) bool {
	var result bool
	if nil != err {
		var codeError *ErrCode
		switch {
		case errors.As(err, &codeError):
			result = true
		default:
			result = false
		}
	} else {
		result = false
	}
	return result
}

func ToErrCode(err error) *ErrCode {
	if !IsErrCode(err) {
		return nil
	}
	var result *ErrCode
	switch {
	case errors.As(err, &result):
		return result
	default:
		return nil
	}
}
