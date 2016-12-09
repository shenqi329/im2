package grpc

import (
	"golang.org/x/net/context"
	//"google.golang.org/grpc"
	grpcPb "im/grpc/pb"
	//imserverError "im/imserver/error"
	//"im/imserver/service"
	"log"
)

type Message struct{}

func (m *Message) CreateMessage(context context.Context, request *grpcPb.CreateMessageRequest) (*grpcPb.CreateMessageReply, error) {

	log.Println("CreateSession")
	// clientConn, ok := ctx.Value("clientConn").(*grpc.ClientConn)

	// if !ok {
	// 	reply := &grpcPb.CreateMessageReply{
	// 		Rid:  request.GetRid(),
	// 		Code: imserverError.CommonInternalServerError,
	// 		Desc: imserverError.ErrorCodeToText(imserverError.CommonInternalServerError),
	// 	}
	// 	return reply, nil
	// }
	// clientConn = clientConn

	// protoMessage, err := service.HandleCreateMessage(nil, request)

	// if err != nil {
	// 	log.Println(err.Error())
	// 	reply := &grpcPb.CreateSessionReply{
	// 		Rid:  request.GetRid(),
	// 		Code: imserverError.CommonInternalServerError,
	// 		Desc: imserverError.ErrorCodeToText(imserverError.CommonInternalServerError),
	// 	}
	// 	return reply, nil
	// }

	// return protoMessage, nil

	// protoMessage, err := service.HandleCreateSession(request)

	// if err != nil {

	// }

	return nil, nil
	// log.Println("CreateMessage")

	// // messageClient := grpcPb.NewMessageClient(m.ClientConn)

	// // reply, err := messageClient.CreateMessage(context, request)

	// return reply, err
}
