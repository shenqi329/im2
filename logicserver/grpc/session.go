package grpc

import (
	"golang.org/x/net/context"
	grpcPb "im/logicserver/grpc/pb"
	"im/logicserver/service"
	"im/logicserver/util/key"
	"log"
)

type Session struct{}

func (s *Session) CreateSession(ctx context.Context, request *grpcPb.CreateSessionRequest) (*grpcPb.CreateSessionResponse, error) {
	return CreateSession(ctx, request)
}

func (s *Session) DeleteUsers(ctx context.Context, request *grpcPb.DeleteSessionUsersRequest) (*grpcPb.Response, error) {
	return DeleteSessionUsers(ctx, request)
}

func (s *Session) AddUsers(ctx context.Context, request *grpcPb.AddSessionUsersRequest) (*grpcPb.Response, error) {
	return nil, nil
}

func CreateSession(ctx context.Context, request *grpcPb.CreateSessionRequest) (*grpcPb.CreateSessionResponse, error) {

	log.Println("CreateSession")
	userId := ctx.Value(key.UserId).(string)

	protoMessage, err := service.CreateSession(request, userId)

	return protoMessage, err
}

func DeleteSessionUsers(ctx context.Context, request *grpcPb.DeleteSessionUsersRequest) (*grpcPb.Response, error) {

	log.Println("DeleteSessionUsers")
	protoMessage, err := service.DeleteSessionUsers(request)

	return protoMessage, err
}

func AddSessionUsers(ctx context.Context, request *grpcPb.AddSessionUsersRequest) (*grpcPb.Response, error) {

	log.Println("AddSessionUsers")
	protoMessage, err := service.AddSessionUsers(request)

	return protoMessage, err
}
