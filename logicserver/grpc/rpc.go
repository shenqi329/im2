package grpc

import (
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	imserverError "im/logicserver/error"
	grpcPb "im/logicserver/grpc/pb"
	"log"
)

type Rpc struct{}

func (m *Rpc) Rpc(ctx context.Context, request *grpcPb.RpcRequest) (*grpcPb.RpcResponse, error) {

	rpcResponse := &grpcPb.RpcResponse{
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

	if request.MessageType == grpcPb.MessageTypeDeviceLoginRequest {

		response, err := HandleLogin(ctx, protoMessage.(*grpcPb.DeviceLoginRequest))
		if err != nil {
			log.Println(err.Error())
			return rpcResponse, nil
		}
		protoBuf, err := proto.Marshal(response)
		if err != nil {
			log.Println(err.Error())
			return rpcResponse, nil
		}
		rpcResponse = &grpcPb.RpcResponse{
			Rid:         request.GetRid(),
			Code:        response.Code,
			Desc:        response.Desc,
			MessageType: grpcPb.MessageTypeDeviceLoginResponse,
			ProtoBuf:    protoBuf,
			ConnId:      request.ConnId,
		}

	} else if request.MessageType == grpcPb.MessageTypeDeviceLoginRequest {

	} else if request.MessageType == grpcPb.MessageTypeCreateMessageRequest {

		response, err := CreateMessage(ctx, protoMessage.(*grpcPb.CreateMessageRequest))
		if err != nil {
			log.Println(err.Error())
			return rpcResponse, nil
		}
		protoBuf, err := proto.Marshal(response)
		if err != nil {
			log.Println(err.Error())
			return rpcResponse, nil
		}
		rpcResponse = &grpcPb.RpcResponse{
			Rid:         request.GetRid(),
			Code:        response.Code,
			Desc:        response.Desc,
			MessageType: grpcPb.MessageTypeCreateMessageResponse,
			ProtoBuf:    protoBuf,
			ConnId:      request.ConnId,
		}

	} else if request.MessageType == grpcPb.MessageTypeCreateSessionRequest {

	}

	return rpcResponse, nil
}
