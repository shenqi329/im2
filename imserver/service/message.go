package service

import (
	//"github.com/golang/protobuf/proto"
	grpcPb "im/grpc/pb"
	imServerBean "im/imserver/bean"
	dao "im/imserver/dao"
	imServerError "im/imserver/error"
	"log"
	"time"
)

func HandleCreateMessage(request *grpcPb.CreateMessageRequest, tokenConnInfoChan chan<- int64) (*grpcPb.CreateMessageReply, error) {

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

	response := &grpcPb.CreateMessageReply{
		Rid:  (uint64)(request.Rid),
		Code: imServerError.CommonSuccess,
		Desc: imServerError.ErrorCodeToText(imServerError.CommonSuccess),
	}
	go xxxxxxxxxxxxxxxxxxx(tokenConnInfoChan, request.SessionId)

	return response, nil
}

func xxxxxxxxxxxxxxxxxxx(tokenConnInfoChan chan<- int64, sessionId int64) {

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
		xxx(tokenConnInfoChan, sessionMap)
	}
}

func xxx(tokenConnInfoChan chan<- int64, sessionMap *imServerBean.SessionMap) {

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
		tokenConnInfoChan <- token.Id
		//token.Id 根据登录的id去发送
	}
}
