package service

import (
	"github.com/golang/protobuf/proto"
	logicserverBean "im/logicserver/bean"
	dao "im/logicserver/dao"
	logicserverError "im/logicserver/error"
	server "im/logicserver/server"
	protocolClient "im/protocol/client"
	"log"
	"strconv"
	"time"
)

func HandleLogin(c server.Context, deviceLoginRequest *protocolClient.DeviceLoginRequest) (protoMessage proto.Message, err error) {

	id, _ := strconv.ParseInt(deviceLoginRequest.Token, 10, 64)

	tokenBean := &logicserverBean.Token{
		Id:       id,
		AppId:    deviceLoginRequest.AppId,
		DeviceId: deviceLoginRequest.DeviceId,
		Platform: deviceLoginRequest.Platform,
	}
	has, err := dao.NewDao().Get(tokenBean)

	if err != nil {
		protoMessage = &protocolClient.DeviceLoginResponse{
			Rid:  deviceLoginRequest.Rid,
			Code: logicserverError.CommonInternalServerError,
			Desc: logicserverError.ErrorCodeToText(logicserverError.CommonInternalServerError),
		}
		return
	}
	if !has {
		protoMessage = &protocolClient.DeviceLoginResponse{
			Rid:  deviceLoginRequest.Rid,
			Code: logicserverError.CommonResourceNoExist,
			Desc: logicserverError.ErrorCodeToText(logicserverError.CommonResourceNoExist),
		}
		//err = logicserverError.ErrorNotFound
		return
	}
	if tokenBean.LoginTime == nil {
		timeNow := time.Now()
		tokenBean.LoginTime = &timeNow
	}

	//将连接设置为登录状态
	// connInfo := c.IMServer().GetConnInfo(c.ConnId())
	// if connInfo != nil {
	// 	if !connInfo.IsLogin {
	// 		connInfo.IsLogin = true
	// 	}
	// 	connInfo.UdpAddr = c.UDPAddr()
	// 	connInfo.UdpConn = c.UDPConn()
	// 	connInfo.ConnId = c.ConnId()
	// 	connInfo.Token = tokenBean.Id
	// 	connInfo.UserId = tokenBean.UserId
	// }

	connInfo := &server.ConnInfo{
		IsLogin: true,
		UdpAddr: c.UDPAddr(),
		UdpConn: c.UDPConn(),
		ConnId:  c.ConnId(),
		Token:   tokenBean.Id,
		UserId:  tokenBean.UserId,
	}

	c.ConnInfoChan() <- connInfo

	go sendSyncInform(c, deviceLoginRequest, tokenBean.UserId)

	protoMessage = &protocolClient.DeviceLoginResponse{
		Rid:  deviceLoginRequest.Rid,
		Code: logicserverError.CommonSuccess,
		Desc: logicserverError.ErrorCodeToText(logicserverError.CommonSuccess),
	}

	return
}

//发送同步通知
func sendSyncInform(c server.Context, deviceLoginRequest *protocolClient.DeviceLoginRequest, userId string) {

	var sessionMaps []*logicserverBean.SessionMap

	err := dao.NewDao().Find(&sessionMaps, &logicserverBean.SessionMap{
		UserId: userId,
	})
	if err != nil {
		log.Println(err)
		return
	}

	for _, sessionMap := range sessionMaps {
		sendSyncInformWithSessionMap(c, sessionMap)
	}
}

func sendSyncInformWithSessionMap(c server.Context, sessionMap *logicserverBean.SessionMap) {

	var messages []*logicserverBean.Message

	err := dao.NewDao().Find(&messages, &logicserverBean.Message{
		SessionId: sessionMap.SessionId,
	})

	latestIndex, err := dao.MessageMaxIndex(sessionMap.SessionId)

	if sessionMap.ReadIndex >= latestIndex {
		//log.Println("不需发送同步通知")
		return
	}

	if err != nil {
		log.Println(err)
		return
	}
	log.Println(latestIndex)

	syncInfo := &protocolClient.SyncInform{
		SessionId:   sessionMap.SessionId,
		LatestIndex: latestIndex,
		ReadIndex:   sessionMap.ReadIndex,
	}

	c.SendProtoMessage(protocolClient.MessageTypeSyncInform, syncInfo)
}
