package grpc

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	grpcPb "im/grpc/pb"
	"log"
)

type Session struct {
	ClientConn *grpc.ClientConn
}

func (s *Session) CreateSession(c context.Context, request *grpcPb.CreateSessionRequest) (*grpcPb.CreateSessionReply, error) {

	log.Println("CreateMessage")

	reply := &grpcPb.CreateSessionReply{
		Rid:  request.GetRid(),
		Code: "200000001",
		Desc: "success",
	}

	return reply, nil
}
