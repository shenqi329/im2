package response

import ()

const (
	//通用状态模块码 [000]
	CommonSuccess             = "00000001"
	CommonIllegalParams       = "00000002"
	CommonResourceNoExist     = "00000003"
	CommonResourceExist       = "00000004"
	CommonInternalServerError = "00000005"
	CommonTokenInvalidated    = "00000006"
)

type (
	Response struct {
		Code string      `json:"code"`
		Desc string      `json:"desc"`
		Data interface{} `json:"data"`
	}
)
