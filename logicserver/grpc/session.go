package grpc

import (
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	grpcPb "im/logicserver/grpc/pb"
	"im/logicserver/service"
	//"im/logicserver/util/key"
	logicserverError "im/logicserver/error"
	"log"
)

type Session struct{}

func (s *Session) CreateSession(ctx context.Context, request *grpcPb.CreateSessionRequest) (*grpcPb.CreateSessionResponse, error) {

	message, err := CreateSession(ctx, request)
	response := message.(*grpcPb.CreateSessionResponse)
	return response, err
}

func (s *Session) DeleteUsers(ctx context.Context, request *grpcPb.DeleteSessionUsersRequest) (*grpcPb.Response, error) {

	message, err := DeleteSessionUsers(ctx, request)
	response := message.(*grpcPb.Response)

	return response, err
}

func (s *Session) AddUsers(ctx context.Context, request *grpcPb.AddSessionUsersRequest) (*grpcPb.Response, error) {

	message, err := AddSessionUsers(ctx, request)
	response := message.(*grpcPb.Response)

	return response, err
}

func CreateSession(ctx context.Context, message proto.Message) (proto.Message, error) {

	log.Println("CreateSession")
	request := message.(*grpcPb.CreateSessionRequest)
	userId := request.RpcInfo.UserId

	protoResponse := &grpcPb.CreateSessionResponse{
		Rid:  request.Rid,
		Code: logicserverError.CommonInternalServerError,
		Desc: logicserverError.ErrorCodeToText(logicserverError.CommonInternalServerError),
	}
	response, err := service.CreateSession(request, userId)
	if err != nil {
		log.Println(err.Error())
		return protoResponse, err
	}
	return response, nil
}

func DeleteSessionUsers(ctx context.Context, message proto.Message) (proto.Message, error) {

	log.Println("DeleteSessionUsers")

	request := message.(*grpcPb.DeleteSessionUsersRequest)
	return service.DeleteSessionUsers(request)
}

func AddSessionUsers(ctx context.Context, message proto.Message) (proto.Message, error) {

	log.Println("AddSessionUsers")
	request := message.(*grpcPb.AddSessionUsersRequest)

	return service.AddSessionUsers(request)
}
