package grpc

import (
	"golang.org/x/net/context"
	//"google.golang.org/grpc"
	grpcPb "im/grpc/pb"
	//"im/imserver/service"
	//"log"
)

type Message struct{}

func (m *Message) CreateMessage(context context.Context, request *grpcPb.CreateMessageRequest) (*grpcPb.CreateMessageReply, error) {

	// protoMessage, err := service.HandleCreateSession(request)

	// if err != nil {

	// }

	return nil, nil
	// log.Println("CreateMessage")

	// // messageClient := grpcPb.NewMessageClient(m.ClientConn)

	// // reply, err := messageClient.CreateMessage(context, request)

	// return reply, err
}
