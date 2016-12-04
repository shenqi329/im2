package service

import (
	"github.com/golang/protobuf/proto"
	imserver "im/imserver"
	imServerBean "im/imserver/bean"
	dao "im/imserver/dao"
	imServerError "im/imserver/error"
	protocolBean "im/protocol/bean"
	"log"
	"strconv"
	"time"
)

func HandleLogin(c imserver.Context, deviceLoginRequest *protocolBean.DeviceLoginRequest) (protoMessage proto.Message, err error) {

	id, _ := strconv.ParseInt(deviceLoginRequest.Token, 10, 64)

	tokenBean := &imServerBean.Token{
		Id:       id,
		AppId:    deviceLoginRequest.AppId,
		DeviceId: deviceLoginRequest.DeviceId,
		Platform: deviceLoginRequest.Platform,
	}
	has, err := dao.NewDao().Get(tokenBean)

	if err != nil {
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
	if tokenBean.LoginTime == nil {
		timeNow := time.Now()
		tokenBean.LoginTime = &timeNow
	}

	//将连接设置为登录状态
	connInfo := c.IMServer().GetConnInfo(c.ConnId())
	if connInfo != nil {
		if !connInfo.IsLogin {
			connInfo.IsLogin = true
		}
	}

	sendNofity(c, deviceLoginRequest)

	protoMessage = &protocolBean.DeviceLoginResponse{
		Rid:  deviceLoginRequest.Rid,
		Code: imServerError.CommonSuccess,
		Desc: imServerError.ErrorCodeToText(imServerError.CommonSuccess),
	}

	return
}

func sendNofity(c imserver.Context, deviceLoginRequest *protocolBean.DeviceLoginRequest) {

	var sessionMaps []*imServerBean.SessionMap

	err := dao.NewDao().Find(&sessionMaps, &imServerBean.SessionMap{
		UserId: "1",
	})
	if err != nil {
		log.Println(err)
		return
	}

	for _, sessionMap := range sessionMaps {
		sendNofityWithSessionMap(sessionMap)
	}
}

func sendNofityWithSessionMap(sessionMap *imServerBean.SessionMap) {

	var messages []*imServerBean.Message

	err := dao.NewDao().Find(&messages, &imServerBean.Message{
		SessionId: sessionMap.SessionId,
	})

	index, err := dao.MessageMaxIndex(sessionMap.SessionId)

	if err != nil {
		log.Println(err)
		return
	}
	log.Println(index)
}
