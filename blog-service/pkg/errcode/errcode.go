package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	ECode    int      `json:"code"`
	EMsg     string   `json:"msg"`
	EDetails []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已存在，请换一个", code))
	}

	codes[code] = msg
	return &Error{ECode: code, EMsg: msg}
}

func (e *Error) Code() int {
	return e.ECode
}

func (e *Error) Msg() string {
	return e.EMsg
}

func (e *Error) Details() []string {
	return e.EDetails
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.EMsg, args...)
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码:%d,错误信息:%s", e.Code(), e.Msg())
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.EDetails = []string{}
	for _, d := range details {
		newError.EDetails = append(newError.EDetails, d)
	}
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
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
