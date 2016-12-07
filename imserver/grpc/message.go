package grpc

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	grpcMessage "im/grpc/message"
	"log"
)

type Message struct {
	ClientConn *grpc.ClientConn
}

func (m *Message) CreateMessage(context context.Context, message *grpcMessage.CreateMessageRequest) (*grpcMessage.CreateMessageReply, error) {
	log.Println("CreateMessage")

	messageClient := grpcMessage.NewMessageClient(m.ClientConn)

	reply, err := messageClient.CreateMessage(context, message)

	return reply, err
}
