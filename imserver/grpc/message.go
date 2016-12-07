package grpc

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"im/grpc/message"
	"log"
)

type Message struct {
	ClientConn *grpc.ClientConn
}

func (m *Message) CreateMessage(context.Context, *message.CreateMessageRequest) (*message.CreateMessageReply, error) {
	log.Println("CreateMessage")
	return nil, nil
}
