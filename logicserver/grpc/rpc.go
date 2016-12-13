package grpc

import (
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	grpcPb "im/grpc/pb"
	imserverError "im/logicserver/error"
	protocolClient "im/protocol/client"
	"log"
)

type Rpc struct{}

func (m *Rpc) Rpc(ctx context.Context, request *protocolClient.RpcRequest) (*protocolClient.RpcResponse, error) {

	rpcResponse := &protocolClient.RpcResponse{
		Rid:    request.GetRid(),
		Code:   imserverError.CommonInternalServerError,
		Desc:   imserverError.ErrorCodeToText(imserverError.CommonInternalServerError),
		ConnId: request.ConnId,
	}

	ctx = context.WithValue(ctx, "UserId", request.UserId)
	ctx = context.WithValue(ctx, "Token", request.Token)
	ctx = context.WithValue(ctx, "ConnId", request.ConnId)

	protoMessage := grpcPb.Factory((grpcPb.MessageType)(request.MessageType))
	err := proto.Unmarshal(request.ProtoBuf, protoMessage)
	if err != nil {
		log.Println(err.Error())
		return rpcResponse, nil
	}

	if request.MessageType == grpcPb.MessageTypeCreateMessageRequest {

		reply, err := CreateMessage(ctx, protoMessage.(*grpcPb.CreateMessageRequest))
		if err != nil {
			log.Println(err.Error())
			return rpcResponse, nil
		}
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
	} else if request.MessageType == grpcPb.MessageTypeCreateSessionRequest {

	}

	return rpcResponse, nil
}
