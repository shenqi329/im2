package response

import (
	imError "im/imserver/error"
	"strings"
)

type (
	Response struct {
		Code string      `json:"code"`
		Desc string      `json:"desc"`
		Data interface{} `json:"data"`
	}
)

func (r *Response) IsSuccesss() bool {
	return strings.EqualFold(r.Code, imError.CommonSuccess)
}

func ResponseToError(response *Response) *imError.IMError {

	if response.IsSuccesss() {
		return nil
	}

	err := &imError.IMError{
		Code: response.Code,
		Desc: response.Desc,
	}

	return err
}
