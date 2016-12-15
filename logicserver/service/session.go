package service

import (
	//"github.com/golang/protobuf/proto"
	logicserverBean "im/logicserver/bean"
	dao "im/logicserver/dao"
	logicserverError "im/logicserver/error"
	grpcPb "im/logicserver/grpc/pb"
	"log"
)

func CreateSession(request *grpcPb.CreateSessionRequest, userId string) (*grpcPb.CreateSessionResponse, error) {

	log.Println(request.String())

	session := &logicserverBean.Session{
		AppId:        request.AppId,
		CreateUserId: userId,
	}

	count, err := dao.NewDao().Insert(session)
	if err != nil {
		err = logicserverError.ErrorInternalServerError
		return nil, err
	}

	sessionMap := &logicserverBean.SessionMap{
		SessionId: session.Id,
		UserId:    userId,
	}
	count, err = dao.NewDao().Insert(sessionMap)
	if err != nil || count != 1 {
		err = logicserverError.ErrorInternalServerError
		return nil, err
	}

	for i := 0; i < len(request.UserIds); i++ {
		sessionMap := &logicserverBean.SessionMap{
			SessionId: session.Id,
			UserId:    request.UserIds[i],
		}
		count, err = dao.NewDao().Insert(sessionMap)
		if err != nil || count != 1 {
			err = logicserverError.ErrorInternalServerError
			return nil, err
		}
	}

	response := &grpcPb.CreateSessionResponse{
		Rid:       (uint64)(request.Rid),
		Code:      logicserverError.CommonSuccess,
		Desc:      logicserverError.ErrorCodeToText(logicserverError.CommonSuccess),
		SessionId: (uint64)(session.Id),
	}

	return response, nil
}

func DeleteSessionUsers(request *grpcPb.DeleteSessionUsersRequest) (*grpcPb.Response, error) {

	log.Println(request.String())

	return nil, nil
}

func AddSessionUsers(request *grpcPb.AddSessionUsersRequest) (*grpcPb.Response, error) {

	return nil, nil
}
