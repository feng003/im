package xerr

import (
	"errors"
	"fmt"
	pkgErrors "github.com/pkg/errors"
)

/**
常用通用固定错误
*/

type CodeError struct {
	errCode uint32
	errMsg  string
}

// 返回给前端的错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

// 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

func NewErrCode(errCode uint32) *CodeError {
	return &CodeError{errCode: errCode, errMsg: MapErrMsg(errCode)}
}

func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{errCode: SERVER_COMMON_ERROR, errMsg: errMsg}
}

// JudgeErrCodeIs 判断错误码是否一致
func JudgeErrCodeIs(errCode uint32, err error) bool {
	causeErr := pkgErrors.Cause(err) // err类型
	var e *CodeError
	if errors.As(causeErr, &e) { //自定义错误类型
		// 判断自定义CodeError是否一致
		return errCode == e.GetErrCode()
	}
	return false
}
