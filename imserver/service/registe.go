package service

import (
	//"encoding/json"
	"github.com/golang/protobuf/proto"
	imServerError "im/imserver/error"
	//imServerResponse "im/imserver/response"
	protocolBean "im/protocol/bean"
	//"log"
	//"net/http"
	//"strings"
)

func HandleRegiste(deviceRegisteRequest *protocolBean.DeviceRegisteRequest) (protoMessage proto.Message, err error) {

	// client := &http.Client{}
	// httpRequest, err := http.NewRequest(http.MethodGet, "http://172.17.0.3:8081/user/info", nil)
	// httpRequest.Header.Set("token", deviceRegisteRequest.SsoToken)
	// resp, err := client.Do(httpRequest)
	// defer resp.Body.Close()

	// responseBean := &imServerResponse.Response{}
	// json.NewDecoder(resp.Body).Decode(responseBean)

	// if !strings.EqualFold(responseBean.Code, imServerError.CommonSuccess) {
	// 	log.Println(responseBean.Desc)
	// 	return
	// }

	protoMessage = &protocolBean.DeviceRegisteResponse{
		Rid:   deviceRegisteRequest.Rid,
		Code:  imServerError.CommonSuccess,
		Desc:  imServerError.ErrorCodeToText(imServerError.CommonSuccess),
		Token: "a token from handle registe",
	}

	return
}
