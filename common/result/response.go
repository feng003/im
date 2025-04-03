package result

import (
	"fmt"
)

type ApiError struct {
	StatusCode int
	Code       int32
	Message    string
}

type ServerError struct {
	ApiError
	Detail error
}

func (e ApiError) Error() string {
	return fmt.Sprintf("(%d) %s", e.Code, e.Message)
}

func (e ApiError) Response() *ErrorRes {
	return &ErrorRes{
		Code: e.Code,
		Msg:  e.Message,
	}
}

type ErrorRes struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}
