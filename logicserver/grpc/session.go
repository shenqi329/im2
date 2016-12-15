package grpc

import (
	"golang.org/x/net/context"
	imserverError "im/logicserver/error"
	grpcPb "im/logicserver/grpc/pb"
	"im/logicserver/service"
	"im/logicserver/util/key"
	"log"
)

type Session struct{}

func (s *Session) CreateSession(ctx context.Context, request *grpcPb.CreateSessionRequest) (*grpcPb.CreateSessionResponse, error) {

	return CreateSession(ctx, request)

}

func CreateSession(ctx context.Context, request *grpcPb.CreateSessionRequest) (*grpcPb.CreateSessionResponse, error) {

	log.Println("CreateSession")
	userId := ctx.Value(key.UserId).(string)

	protoMessage, err := service.HandleCreateSession(request, userId)

	if err != nil {
		log.Println(err.Error())
		reply := &grpcPb.CreateSessionResponse{
			Rid:  request.GetRid(),
			Code: imserverError.CommonInternalServerError,
			Desc: imserverError.ErrorCodeToText(imserverError.CommonInternalServerError),
		}
		return reply, nil
	}

	return protoMessage, nil
}
