package service

import (
	// "encoding/json"
	"github.com/golang/protobuf/proto"
	imserver "im/imserver"
	imServerBean "im/imserver/bean"
	dao "im/imserver/dao"
	imServerError "im/imserver/error"
	protocolClient "im/protocol/client"
	"log"
	// "net/http"
	// "strconv"
	"time"
)

func HandleCreateMessage(c imserver.Context, request *protocolClient.CreateMessageRequest) (proto.Message, error) {

	//log.Println(imServerBean.StructToJsonString(request))

	//log.Println(request.SessionId)
	index, err := dao.MessageMaxIndex(request.SessionId)
	if err != nil {
		log.Println(err)
		return nil, imServerError.ErrorInternalServerError
	}

	timeNow := time.Now()
	message := &imServerBean.Message{
		SessionId:  request.SessionId,
		Type:       (int)(request.Type),
		Content:    request.Content,
		Index:      index + 1,
		CreateTime: &timeNow,
	}

	_, err = dao.NewDao().Insert(message)
	if err != nil {
		log.Println(err)
		return nil, imServerError.ErrorInternalServerError
	}

	response := &protocolClient.CreateMessageResponse{
		Rid:  (uint64)(request.Rid),
		Code: imServerError.CommonSuccess,
		Desc: imServerError.ErrorCodeToText(imServerError.CommonSuccess),
	}
	go xxxxxxxxxxxxxxxxxxx(c, request.SessionId)

	return response, nil
}

func xxxxxxxxxxxxxxxxxxx(c imserver.Context, sessionId int64) {

	var sessionMaps []*imServerBean.SessionMap

	err := dao.NewDao().Find(&sessionMaps,
		&imServerBean.SessionMap{
			SessionId: sessionId,
		})
	if err != nil {
		log.Println(err)
		return
	}

	for _, sessionMap := range sessionMaps {
		xxx(c, sessionMap)
	}
}

func xxx(c imserver.Context, sessionMap *imServerBean.SessionMap) {

	var tokens []*imServerBean.Token

	err := dao.NewDao().Find(&tokens,
		&imServerBean.Token{
			UserId: sessionMap.UserId,
		})
	if err != nil {
		log.Println(err)
		return
	}

	for _, token := range tokens {
		log.Println(token.Id)
		log.Println(c.TokenConnInfoChan())
		c.TokenConnInfoChan() <- token.Id
		//token.Id 根据登录的id去发送
	}
}
