package service

import (
	"encoding/json"
	imServerBean "im/imserver/bean"
	imServerResponse "im/imserver/response"
	protocalBean "im/protocal/bean"
	"log"
	"net/http"
	"strings"
)

func HandleRegisteRequest(deviceRegisteRequest *protocalBean.DeviceRegisteRequest) (tokenBean *imServerBean.Token, err error) {

	client := &http.Client{}
	httpRequest, err := http.NewRequest(http.MethodGet, "http://172.17.0.3:8081/user/info", nil)
	httpRequest.Header.Set("token", deviceRegisteRequest.SsoToken)
	resp, err := client.Do(httpRequest)
	defer resp.Body.Close()

	responseBean := &imServerResponse.Response{}
	json.NewDecoder(resp.Body).Decode(responseBean)

	if !strings.EqualFold(responseBean.Code, imServerResponse.CommonSuccess) {
		log.Println(responseBean.Desc)
		return
	}

	tokenBean = &imServerBean.Token{
		Token: "123456",
	}

	return
}
