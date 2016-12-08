package grpc

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	grpcPb "im/grpc/pb"
	imserverError "im/imserver/error"
	serverGrpc "im/imserver/server/grpc"
	"im/imserver/service"
	"log"
)

type Session struct{}

func (s *Session) CreateSession(ctx context.Context, request *grpcPb.CreateSessionRequest) (*grpcPb.CreateSessionReply, error) {

	log.Println("CreateSession")

	clientConn, ok := ctx.Value(serverGrpc.KeyClientConn).(*grpc.ClientConn)

	if !ok {
		reply := &grpcPb.CreateSessionReply{
			Rid:  request.GetRid(),
			Code: imserverError.CommonInternalServerError,
			Desc: imserverError.ErrorCodeToText(imserverError.CommonInternalServerError),
		}
		return reply, nil
	}

	log.Println(clientConn)

	protoMessage, err := service.HandleCreateSession(request)

	if err != nil {
		log.Println(err.Error())
		reply := &grpcPb.CreateSessionReply{
			Rid:  request.GetRid(),
			Code: imserverError.CommonInternalServerError,
			Desc: imserverError.ErrorCodeToText(imserverError.CommonInternalServerError),
		}
		return reply, nil
	}
	protoMessage = protoMessage
	reply := &grpcPb.CreateSessionReply{
		Rid:  request.GetRid(),
		Code: "200000001",
		Desc: "success",
	}
	return reply, nil
}
