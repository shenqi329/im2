package grpc

import (
	"golang.org/x/net/context"
	imserverError "im/logicserver/error"
	grpcPb "im/logicserver/grpc/pb"
	"im/logicserver/service"
	"im/logicserver/util/key"
	"log"
)

type Message struct{}

func (m *Message) CreateMessage(context context.Context, request *grpcPb.CreateMessageRequest) (*grpcPb.CreateMessageResponse, error) {

	return CreateMessage(context, request)

}

func CreateMessage(context context.Context, request *grpcPb.CreateMessageRequest) (*grpcPb.CreateMessageResponse, error) {
	log.Println("CreateMessage")

	//tokenConnInfoChan := context.Value(key.TokenConnInfoChan).(chan int64)
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
