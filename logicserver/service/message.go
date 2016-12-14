package service

import (
	//"github.com/golang/protobuf/proto"
	logicserverBean "im/logicserver/bean"
	dao "im/logicserver/dao"
	logicserverError "im/logicserver/error"
	grpcPb "im/logicserver/grpc/pb"
	"log"
	"strings"
	"time"
)

func HandleCreateMessage(request *grpcPb.CreateMessageRequest, userId string) (*grpcPb.CreateMessageResponse, error) {

	// index, err := dao.MessageMaxIndex(request.SessionId)
	// if err != nil {
	// 	log.Println(err)
	// 	return nil, logicserverError.ErrorInternalServerError
	// }

	timeNow := time.Now()
	message := &logicserverBean.Message{
		Id:         request.Id,
		SessionId:  request.SessionId,
		UserId:     userId,
		Type:       (int)(request.Type),
		Content:    request.Content,
		CreateTime: &timeNow,
	}

	_, err := dao.MessageInsert(message)
	if err != nil {
		log.Println(err)
		return nil, logicserverError.ErrorInternalServerError
	}

	if request.SessionId > 0 {
		var sessionMaps []*logicserverBean.SessionMap
		dao.NewDao().Find(&sessionMaps, &logicserverBean.SessionMap{
			SessionId: request.SessionId,
		})

		for _, sessionMap := range sessionMaps {
			if strings.EqualFold(sessionMap.UserId, userId) {
				continue
			}
			message := &logicserverBean.Message{
				Id:         request.Id,
				SessionId:  request.SessionId,
				UserId:     sessionMap.UserId,
				Type:       (int)(request.Type),
				Content:    request.Content,
				CreateTime: &timeNow,
			}
			dao.MessageInsert(message)
		}
	}

	response := &grpcPb.CreateMessageResponse{
		Rid:  (uint64)(request.Rid),
		Code: logicserverError.CommonSuccess,
		Desc: logicserverError.ErrorCodeToText(logicserverError.CommonSuccess),
	}

	//go xxxxxxxxxxxxxxxxxxx(tokenConnInfoChan, request.SessionId)

	return response, nil
}

// func xxxxxxxxxxxxxxxxxxx(tokenConnInfoChan chan<- int64, sessionId int64) {

// 	var sessionMaps []*logicserverBean.SessionMap

// 	err := dao.NewDao().Find(&sessionMaps,
// 		&logicserverBean.SessionMap{
// 			SessionId: sessionId,
// 		})
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	for _, sessionMap := range sessionMaps {
// 		xxx(tokenConnInfoChan, sessionMap)
// 	}
// }

// func xxx(tokenConnInfoChan chan<- int64, sessionMap *logicserverBean.SessionMap) {

// 	var tokens []*logicserverBean.Token

// 	err := dao.NewDao().Find(&tokens,
// 		&logicserverBean.Token{
// 			UserId: sessionMap.UserId,
// 		})
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	for _, token := range tokens {
// 		//log.Println(token.Id)
// 		tokenConnInfoChan <- token.Id
// 		//token.Id 根据登录的id去发送
// 	}
// }
