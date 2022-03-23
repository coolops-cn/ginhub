package errors

import (
	"fmt"
	"net/http"
)

type Error struct {
	code    int
	message string
	details []string
}

var codes = map[int]string{}

// NewError 初始化Error
func NewError(code int, message string) *Error {
	// 判断 code 是否已经存在
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码%d已经存在,请更换一个", code))
	}
	codes[code] = message
	return &Error{code: code, message: message}
}

// Error 格式化错误信息
func (e *Error) Error() string {
	return fmt.Sprintf("错误码: %d, 错误信息: %s", e.Code(), e.Msg())
}

// Code 获取错误状态码
func (e *Error) Code() int {
	return e.code
}

// Msg 获取错误信息
func (e *Error) Msg() string {
	return e.message
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.Msg(), args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	newError.details = append(newError.details, details...)
	return &newError
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case NotFound.Code():
		return http.StatusNotFound
	case ErrHeaderMalformed.Code():
		fallthrough
	case ErrTokenExpired.Code():
		fallthrough
	case ErrTokenExpiredMaxRefresh.Code():
		fallthrough
	case ErrTokenMalformed.Code():
		fallthrough
	case ErrHeaderEmpty.Code():
		fallthrough
	case ErrTokenInvalid.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
