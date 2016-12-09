package server

import (
	"github.com/golang/protobuf/proto"
	protocolClient "im/protocol/client"
	"log"
)

func Handle(context Context) error {
	protoMessage := protocolClient.Factory((protocolClient.MessageType)(context.Message().Type))
	requestBean := &protocolClient.DeviceRegisteRequest{}
	proto.Unmarshal(context.Message().Body, requestBean)

	if protoMessage == nil {
		log.Println("未识别的消息")
		context.Close()
		return nil
	}
	if err := proto.Unmarshal(context.Message().Body, protoMessage); err != nil {
		log.Println(err.Error())
		context.Close()
		return nil
	}

	//只检查消息的合法性,然后将消息转发出去
	context.Request().message = context.Message()
	context.Request().protoMessage = protoMessage

	context.ReqChan() <- context.Request()

	return nil
}
