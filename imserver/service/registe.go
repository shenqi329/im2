package service

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	imServerBean "im/imserver/bean"
	dao "im/imserver/dao"
	imServerError "im/imserver/error"
	imServerResponse "im/imserver/response"
	protocolClient "im/protocol/client"
	"log"
	"net/http"
	"strconv"
	"time"
)

func CheckDeviceRegistReqeust(deviceRegisteRequest *protocolClient.DeviceRegisteRequest) error {

	if err := CheckDeviceId(deviceRegisteRequest.DeviceId); err != nil {
		return err
	}

	if err := CheckToken(deviceRegisteRequest.SsoToken); err != nil {
		return err
	}

	return nil
}

func HandleRegiste(deviceRegisteRequest *protocolClient.DeviceRegisteRequest) (protoMessage proto.Message, err error) {

	if err = CheckDeviceRegistReqeust(deviceRegisteRequest); err != nil {
		return
	}

	client := &http.Client{}
	httpRequest, err := http.NewRequest(http.MethodGet, "http://localhost:8081/user/info", nil)
	httpRequest.Header.Set("token", deviceRegisteRequest.SsoToken)
	resp, err := client.Do(httpRequest)
	defer resp.Body.Close()

	if err != nil {
		log.Println(err)
		err = imServerError.ErrorInternalServerError
		return
	}

	responseBean := &imServerResponse.Response{}
	json.NewDecoder(resp.Body).Decode(responseBean)

	if responseBean.IsFail() {
		err = responseBean.ResponseToError()
		return
	}

	maps, ok := responseBean.Data.(map[string]interface{})
	if !ok {
		err = imServerError.ErrorInternalServerError
		return
	}

	log.Println(imServerBean.StructToJsonString(responseBean))
	log.Println(deviceRegisteRequest)

	createTime := time.Now()
	tokenBean := &imServerBean.Token{
		UserId:     maps["id"].(string),
		DeviceId:   deviceRegisteRequest.DeviceId,
		AppId:      deviceRegisteRequest.AppId,
		Platform:   deviceRegisteRequest.Platform,
		CreateTime: &createTime,
	}

	_, err = dao.NewDao().Insert(tokenBean)
	if err != nil {
		err = imServerError.ErrorInternalServerError
		return
	}

	protoMessage = &protocolClient.DeviceRegisteResponse{
		Rid:   deviceRegisteRequest.Rid,
		Code:  imServerError.CommonSuccess,
		Desc:  imServerError.ErrorCodeToText(imServerError.CommonSuccess),
		Token: strconv.FormatInt(tokenBean.Id, 10),
	}

	return
}
