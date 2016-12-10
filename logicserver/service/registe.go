package service

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	logicserverBean "im/logicserver/bean"
	dao "im/logicserver/dao"
	logicserverError "im/logicserver/error"
	logicserverResponse "im/logicserver/response"
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
		err = logicserverError.ErrorInternalServerError
		return
	}

	responseBean := &logicserverResponse.Response{}
	json.NewDecoder(resp.Body).Decode(responseBean)

	if responseBean.IsFail() {
		err = responseBean.ResponseToError()
		return
	}

	maps, ok := responseBean.Data.(map[string]interface{})
	if !ok {
		err = logicserverError.ErrorInternalServerError
		return
	}

	log.Println(logicserverBean.StructToJsonString(responseBean))
	log.Println(deviceRegisteRequest)

	createTime := time.Now()
	tokenBean := &logicserverBean.Token{
		UserId:     maps["id"].(string),
		DeviceId:   deviceRegisteRequest.DeviceId,
		AppId:      deviceRegisteRequest.AppId,
		Platform:   deviceRegisteRequest.Platform,
		CreateTime: &createTime,
	}

	_, err = dao.NewDao().Insert(tokenBean)
	if err != nil {
		err = logicserverError.ErrorInternalServerError
		return
	}

	protoMessage = &protocolClient.DeviceRegisteResponse{
		Rid:   deviceRegisteRequest.Rid,
		Code:  logicserverError.CommonSuccess,
		Desc:  logicserverError.ErrorCodeToText(logicserverError.CommonSuccess),
		Token: strconv.FormatInt(tokenBean.Id, 10),
	}

	return
}
