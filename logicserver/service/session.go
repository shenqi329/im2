package service

import (
	//"github.com/golang/protobuf/proto"
	logicserverBean "im/logicserver/bean"
	dao "im/logicserver/dao"
	logicserverError "im/logicserver/error"
	grpcPb "im/logicserver/grpc/pb"
	"log"
)

func HandleCreateSession(request *grpcPb.CreateSessionRequest) (*grpcPb.CreateSessionResponse, error) {

	log.Println(request.String())

	if request.Count > 10 ||
		request.Count <= 0 {
		err := logicserverError.ErrorNotFound
		return nil, err
	}

	sessions := make([]*logicserverBean.Session, request.Count)
	for i := 0; i < (int)(request.Count); i++ {
		sessions[i] = &logicserverBean.Session{
			AppId:        request.AppId,
			CreateUserId: request.UserId,
		}
		count, err := dao.NewDao().Insert(sessions[i])
		if err != nil || count != 1 {
			err = logicserverError.ErrorInternalServerError
			return nil, err
		}
		sessionMap := &logicserverBean.SessionMap{
			SessionId: sessions[i].Id,
			UserId:    request.UserId,
		}
		count, err = dao.NewDao().Insert(sessionMap)
		if err != nil || count != 1 {
			err = logicserverError.ErrorInternalServerError
			return nil, err
		}
	}

	sessionIds := make([]*grpcPb.SessionInfo, request.Count)
	for i := 0; i < (int)(request.Count); i++ {
		log.Print(sessions[i].Id)
		sessionIds[i] = &grpcPb.SessionInfo{
			SessionId: (uint64)(sessions[i].Id),
		}
	}

	response := &grpcPb.CreateSessionResponse{
		Rid:        (uint64)(request.Rid),
		Code:       logicserverError.CommonSuccess,
		Desc:       logicserverError.ErrorCodeToText(logicserverError.CommonSuccess),
		SessionIds: sessionIds,
	}

	return response, nil
}
