package grpc

import (
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	imserverError "im/logicserver/error"
	grpcPb "im/logicserver/grpc/pb"
	"im/logicserver/service"
	"im/logicserver/util/key"
	"log"
)

type Message struct{}

func (m *Message) CreateMessage(context context.Context, request *grpcPb.CreateMessageRequest) (*grpcPb.CreateMessageResponse, error) {

	message, err := CreateMessage(context, request)
	response := message.(*grpcPb.CreateMessageResponse)

	return response, err
}

func CreateMessage(context context.Context, message proto.Message) (proto.Message, error) {

	log.Println("CreateMessage")

	request := message.(*grpcPb.CreateMessageRequest)
	userId := context.Value(key.UserId).(string)

	protoMessage, err := service.HandleCreateMessage(request, userId)

	if err != nil {
		log.Println(err.Error())
		reply := &grpcPb.CreateMessageResponse{
			Rid:  request.GetRid(),
			Code: imserverError.CommonInternalServerError,
			Desc: imserverError.ErrorCodeToText(imserverError.CommonInternalServerError),
		}
		return reply, nil
	}
	return protoMessage, nil
}
