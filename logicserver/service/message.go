package service

import (
	//"github.com/golang/protobuf/proto"
	grpcPb "im/grpc/pb"
	logicserverBean "im/logicserver/bean"
	dao "im/logicserver/dao"
	logicserverError "im/logicserver/error"
	"log"
	"time"
)

func HandleCreateMessage(request *grpcPb.CreateMessageRequest, tokenConnInfoChan chan<- int64) (*grpcPb.CreateMessageReply, error) {

	index, err := dao.MessageMaxIndex(request.SessionId)
	if err != nil {
		log.Println(err)
		return nil, logicserverError.ErrorInternalServerError
	}

	timeNow := time.Now()
	message := &logicserverBean.Message{
		SessionId:  request.SessionId,
		Type:       (int)(request.Type),
		Content:    request.Content,
		Index:      index + 1,
		CreateTime: &timeNow,
	}

	_, err = dao.NewDao().Insert(message)
	if err != nil {
		log.Println(err)
		return nil, logicserverError.ErrorInternalServerError
	}

	response := &grpcPb.CreateMessageReply{
		Rid:  (uint64)(request.Rid),
		Code: logicserverError.CommonSuccess,
		Desc: logicserverError.ErrorCodeToText(logicserverError.CommonSuccess),
	}
	//go xxxxxxxxxxxxxxxxxxx(tokenConnInfoChan, request.SessionId)

	return response, nil
}

func xxxxxxxxxxxxxxxxxxx(tokenConnInfoChan chan<- int64, sessionId int64) {

	var sessionMaps []*logicserverBean.SessionMap

	err := dao.NewDao().Find(&sessionMaps,
		&logicserverBean.SessionMap{
			SessionId: sessionId,
		})
	if err != nil {
		log.Println(err)
		return
	}

	for _, sessionMap := range sessionMaps {
		xxx(tokenConnInfoChan, sessionMap)
	}
}

func xxx(tokenConnInfoChan chan<- int64, sessionMap *logicserverBean.SessionMap) {

	var tokens []*logicserverBean.Token

	err := dao.NewDao().Find(&tokens,
		&logicserverBean.Token{
			UserId: sessionMap.UserId,
		})
	if err != nil {
		log.Println(err)
		return
	}

	for _, token := range tokens {
		//log.Println(token.Id)
		tokenConnInfoChan <- token.Id
		//token.Id 根据登录的id去发送
	}
}
