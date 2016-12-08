package service

import (
	//"github.com/golang/protobuf/proto"
	grpcPb "im/grpc/pb"
	imServerBean "im/imserver/bean"
	dao "im/imserver/dao"
	imServerError "im/imserver/error"
	"log"
)

func HandleCreateSession(request *grpcPb.CreateSessionRequest) (*grpcPb.CreateSessionReply, error) {

	log.Println(request.String())

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

	sessionIds := make([]*grpcPb.SessionInfo, request.Count)
	for i := 0; i < (int)(request.Count); i++ {
		log.Print(sessions[i].Id)
		sessionIds[i] = &grpcPb.SessionInfo{
			SessionId: (uint64)(sessions[i].Id),
		}
	}

	response := &grpcPb.CreateSessionReply{
		Rid:        (uint64)(request.Rid),
		Code:       imServerError.CommonSuccess,
		Desc:       imServerError.ErrorCodeToText(imServerError.CommonSuccess),
		SessionIds: sessionIds,
	}

	return response, nil
}
