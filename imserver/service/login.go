package service

import (
	// "encoding/json"
	"github.com/golang/protobuf/proto"
	imServerBean "im/imserver/bean"
	dao "im/imserver/dao"
	imServerError "im/imserver/error"
	// imServerResponse "im/imserver/response"
	protocolBean "im/protocol/bean"
	//"log"
	// "net/http"
	"strconv"
	// "strings"
	// "time"
)

func HandleLogin(deviceLoginRequest *protocolBean.DeviceLoginRequest) (protoMessage proto.Message, err error) {

	id, _ := strconv.ParseInt(deviceLoginRequest.Token, 10, 64)

	tokenBean := &imServerBean.Token{
		Id:       id,
		AppId:    deviceLoginRequest.AppId,
		DeviceId: deviceLoginRequest.DeviceId,
		Platform: deviceLoginRequest.Platform,
	}
	//log.Println(imServerBean.StructToJsonString(tokenBean))

	has, err := dao.GetToken(tokenBean)

	if err != nil {
		//err = imServerError.ErrorInternalServerError
		protoMessage = &protocolBean.DeviceLoginResponse{
			Rid:  deviceLoginRequest.Rid,
			Code: imServerError.CommonInternalServerError,
			Desc: imServerError.ErrorCodeToText(imServerError.CommonInternalServerError),
		}
		return
	}
	if !has {
		protoMessage = &protocolBean.DeviceLoginResponse{
			Rid:  deviceLoginRequest.Rid,
			Code: imServerError.CommonResourceNoExist,
			Desc: imServerError.ErrorCodeToText(imServerError.CommonResourceNoExist),
		}
		//err = imServerError.ErrorNotFound
		return
	}

	protoMessage = &protocolBean.DeviceLoginResponse{
		Rid:  deviceLoginRequest.Rid,
		Code: imServerError.CommonSuccess,
		Desc: imServerError.ErrorCodeToText(imServerError.CommonSuccess),
	}

	return
}
