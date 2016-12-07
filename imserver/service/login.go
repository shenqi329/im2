package service

import (
	"github.com/golang/protobuf/proto"
	imserver "im/imserver"
	imServerBean "im/imserver/bean"
	dao "im/imserver/dao"
	imServerError "im/imserver/error"
	protocolClient "im/protocol/client"
	"log"
	"strconv"
	"time"
)

func HandleLogin(c imserver.Context, deviceLoginRequest *protocolClient.DeviceLoginRequest) (protoMessage proto.Message, err error) {

	id, _ := strconv.ParseInt(deviceLoginRequest.Token, 10, 64)

	tokenBean := &imServerBean.Token{
		Id:       id,
		AppId:    deviceLoginRequest.AppId,
		DeviceId: deviceLoginRequest.DeviceId,
		Platform: deviceLoginRequest.Platform,
	}
	has, err := dao.NewDao().Get(tokenBean)

	if err != nil {
		protoMessage = &protocolClient.DeviceLoginResponse{
			Rid:  deviceLoginRequest.Rid,
			Code: imServerError.CommonInternalServerError,
			Desc: imServerError.ErrorCodeToText(imServerError.CommonInternalServerError),
		}
		return
	}
	if !has {
		protoMessage = &protocolClient.DeviceLoginResponse{
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

	connInfo := &imserver.ConnInfo{
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
		Code: imServerError.CommonSuccess,
		Desc: imServerError.ErrorCodeToText(imServerError.CommonSuccess),
	}

	return
}

//发送同步通知
func sendSyncInform(c imserver.Context, deviceLoginRequest *protocolClient.DeviceLoginRequest, userId string) {

	var sessionMaps []*imServerBean.SessionMap

	err := dao.NewDao().Find(&sessionMaps, &imServerBean.SessionMap{
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

func sendSyncInformWithSessionMap(c imserver.Context, sessionMap *imServerBean.SessionMap) {

	var messages []*imServerBean.Message

	err := dao.NewDao().Find(&messages, &imServerBean.Message{
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
