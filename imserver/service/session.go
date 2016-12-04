package service

import (
	// "encoding/json"
	"github.com/golang/protobuf/proto"
	imServerBean "im/imserver/bean"
	dao "im/imserver/dao"
	imServerError "im/imserver/error"
	// imServerResponse "im/imserver/response"
	protocolBean "im/protocol/bean"
	"log"
	// "net/http"
	// "strconv"
	// "time"
)

func HandleCreateSession(request *protocolBean.CreateSessionRequest) (proto.Message, error) {

	log.Println(imServerBean.StructToJsonString(request))

	if request.Count > 10 ||
		request.Count <= 0 {
		err := imServerError.ErrorNotFound
		return nil, err
	}

	sessions := make([]*imServerBean.Session, request.Count)
	for i := 0; i < (int)(request.Count); i++ {
		sessions[i] = &imServerBean.Session{
			AppId:        request.AppId,
			CreateUserId: request.CreateUserId,
		}
		count, err := dao.NewDao().Insert(sessions[i])
		if err != nil || count != 1 {
			err = imServerError.ErrorInternalServerError
			return nil, err
		}
		sessionMap := &imServerBean.SessionMap{
			SessionId: sessions[i].Id,
			UserId:    request.CreateUserId,
			ReadIndex: 0,
		}
		count, err = dao.NewDao().Insert(sessionMap)
		if err != nil || count != 1 {
			err = imServerError.ErrorInternalServerError
			return nil, err
		}
	}

	sessionIds := make([]*protocolBean.Session, request.Count)
	for i := 0; i < (int)(request.Count); i++ {
		log.Print(sessions[i].Id)
		sessionIds[i] = &protocolBean.Session{
			SessionId: (uint64)(sessions[i].Id),
		}
	}

	response := &protocolBean.CreateSessionResponse{
		Rid:        (uint64)(request.Rid),
		Code:       imServerError.CommonSuccess,
		Desc:       imServerError.ErrorCodeToText(imServerError.CommonSuccess),
		SessionIds: sessionIds,
	}

	return response, nil
}
