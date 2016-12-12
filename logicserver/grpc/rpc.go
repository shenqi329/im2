package grpc

import (
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	grpcPb "im/grpc/pb"
	imserverError "im/logicserver/error"
	"im/logicserver/service"
	protocolClient "im/protocol/client"
	"log"
)

type Rpc struct{}

func (m *Rpc) Rpc(context context.Context, request *protocolClient.RpcRequest) (*protocolClient.RpcResponse, error) {

	//log.Println("Rpc")

	v := context.Value("tokenConnInfoChan")

	rpcResponse := &protocolClient.RpcResponse{
		Rid:    request.GetRid(),
		Code:   imserverError.CommonInternalServerError,
		Desc:   imserverError.ErrorCodeToText(imserverError.CommonInternalServerError),
		ConnId: request.ConnId,
	}

	tokenConnInfoChan, ok := v.(chan int64)
	if !ok {
		return rpcResponse, nil
	}

	//设置缓冲策略

	if request.MessageType == grpcPb.MessageTypeCreateMessageRequest {

		protoMessage := grpcPb.Factory((grpcPb.MessageType)(request.MessageType))
		err := proto.Unmarshal(request.ProtoBuf, protoMessage)

		if err != nil {
			log.Println(err.Error())
			return rpcResponse, nil
		}
		//log.Println(protoMessage.String())

		reply, err := service.HandleCreateMessage(protoMessage.(*grpcPb.CreateMessageRequest), tokenConnInfoChan)
		if err != nil {
			log.Println(err.Error())
			return rpcResponse, nil
		}

		//log.Println(reply.String())
		//log.Println(request.ConnId)
		protoBuf, err := proto.Marshal(reply)
		if err != nil {
			log.Println(err.Error())
			return rpcResponse, nil
		}
		rpcResponse = &protocolClient.RpcResponse{
			Rid:         request.GetRid(),
			Code:        reply.Code,
			Desc:        reply.Desc,
			MessageType: grpcPb.MessageTypeCreateMessageReply,
			ProtoBuf:    protoBuf,
			ConnId:      request.ConnId,
		}
		return rpcResponse, nil
	}

	return rpcResponse, nil
}
