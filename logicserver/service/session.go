package service

import (
	//"github.com/golang/protobuf/proto"
	logicserverBean "im/logicserver/bean"
	dao "im/logicserver/dao"
	logicserverError "im/logicserver/error"
	grpcPb "im/logicserver/grpc/pb"
	"log"
	"strings"
)

func CreateSession(request *grpcPb.CreateSessionRequest, userId string) (*grpcPb.CreateSessionResponse, error) {
	log.Println(request.String())

	session := &logicserverBean.Session{
		AppId:        request.RpcInfo.AppId,
		CreateUserId: userId,
	}

	_, err := dao.NewDao().Insert(session)
	if err != nil {
		log.Println(err.Error())
		err = logicserverError.ErrorInternalServerError
		return nil, err
	}

	sessionMap := &logicserverBean.SessionMap{
		SessionId: session.Id,
		UserId:    userId,
	}
	sessionMaps := []interface{}{sessionMap}

	for i := 0; i < len(request.UserIds); i++ {
		if strings.EqualFold(request.UserIds[i], userId) {
			continue
		}
		sessionMap = &logicserverBean.SessionMap{
			SessionId: session.Id,
			UserId:    request.UserIds[i],
		}
		sessionMaps = append(sessionMaps, sessionMap)
	}
	_, err = dao.NewDao().Insert(sessionMaps...)

	if err != nil {
		log.Println(err.Error())
		err = logicserverError.ErrorInternalServerError
		return nil, err
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
