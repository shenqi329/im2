package response

import ()

type (
	Response struct {
		Code string      `json:"code"`
		Desc string      `json:"desc"`
		Data interface{} `json:"data"`
	}
)
