package error

import (
	//"errors"
	"fmt"
)

var (
	//通用错误
	ErrorIllegalParams       = NEWError(CommonIllegalParams)
	ErrorResourceExist       = NEWError(CommonResourceExist)
	ErrorNotFound            = NEWError(CommonResourceNoExist)
	ErrorInternalServerError = NEWError(CommonInternalServerError)
	ErrorTokenInvalidated    = NEWError(CommonTokenInvalidated)
)

const (
	//通用状态模块码 [000]
	CommonSuccess             = "00000001"
	CommonIllegalParams       = "00000002"
	CommonResourceNoExist     = "00000003"
	CommonResourceExist       = "00000004"
	CommonInternalServerError = "00000005"
	CommonTokenInvalidated    = "00000006"
)

var codeText = map[string]string{
	//通用状态
	CommonSuccess:             "success",
	CommonIllegalParams:       "illegal parameter",
	CommonResourceNoExist:     "resource doesn't exist",
	CommonResourceExist:       "resource already exists",
	CommonInternalServerError: "internal server wrong",
	CommonTokenInvalidated:    "token invalidated",
}

func ErrorCodeToText(code string) string {
	return codeText[code]
}

type (
	SSOError struct {
		Code string
		Desc string
	}
)

func NEWError(code string) *SSOError {
	return &SSOError{Code: code, Desc: ErrorCodeToText(code)}
}

func (err *SSOError) Error() string {
	errString := fmt.Sprintf("code = %s,desc = %s", err.Code, err.Desc)
	return errString
}