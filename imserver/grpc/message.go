package grpc

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	grpcPb "im/grpc/pb"
	//"log"
)

type Message struct {
	ClientConn *grpc.ClientConn
}

func (m *Message) CreateMessage(context context.Context, request *grpcPb.CreateMessageRequest) (*grpcPb.CreateMessageReply, error) {
	return nil, nil
	// log.Println("CreateMessage")

	// // messageClient := grpcPb.NewMessageClient(m.ClientConn)

	// // reply, err := messageClient.CreateMessage(context, request)

	// return reply, err
}
