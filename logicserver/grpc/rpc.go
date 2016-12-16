package grpc

import (
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	logicserverError "im/logicserver/error"
	grpcPb "im/logicserver/grpc/pb"
	"log"
)

//protoc --go_out=plugins=grpc:. *.proto

type HandleFunc func(ctx context.Context, message proto.Message) (proto.Message, error)

type HandleFuncInfo struct {
	handle       HandleFunc
	responseType grpcPb.MessageType
}

type Rpc struct {
	handleFuncMap map[grpcPb.MessageType]*HandleFuncInfo
}

func (r *Rpc) AddHandleFunc(messageType grpcPb.MessageType, responseType grpcPb.MessageType, handle HandleFunc) {
	if r.handleFuncMap == nil {
		r.handleFuncMap = make(map[grpcPb.MessageType]*HandleFuncInfo)
	}
	r.handleFuncMap[messageType] = &HandleFuncInfo{
		handle:       handle,
		responseType: responseType,
	}
}

func (r *Rpc) Rpc(ctx context.Context, request *grpcPb.RpcRequest) (*grpcPb.RpcResponse, error) {

	rpcResponse := &grpcPb.RpcResponse{
		Rid:    request.GetRid(),
		Code:   logicserverError.CommonInternalServerError,
		Desc:   logicserverError.ErrorCodeToText(logicserverError.CommonInternalServerError),
		ConnId: request.ConnId,
	}

	ctx = context.WithValue(ctx, "UserId", request.UserId)
	ctx = context.WithValue(ctx, "Token", request.Token)
	ctx = context.WithValue(ctx, "ConnId", request.ConnId)

	handleFuncInfo := r.handleFuncMap[(grpcPb.MessageType)(request.MessageType)]
	if handleFuncInfo == nil {
		log.Println("不支持的类型")
		return nil, nil
	}

	protoMessage := grpcPb.Factory((grpcPb.MessageType)(request.MessageType))
	err := proto.Unmarshal(request.ProtoBuf, protoMessage)
	if err != nil {
		log.Println(err.Error())
		return rpcResponse, nil
	}

	response, err := handleFuncInfo.handle(ctx, protoMessage)
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
		Code:        logicserverError.CommonSuccess,
		Desc:        logicserverError.ErrorCodeToText(logicserverError.CommonSuccess),
		MessageType: (int32)(handleFuncInfo.responseType),
		ProtoBuf:    protoBuf,
		ConnId:      request.ConnId,
	}

	// if request.MessageType == grpcPb.MessageTypeDeviceLoginRequest {

	// 	response, err := HandleLogin(ctx, protoMessage.(*grpcPb.DeviceLoginRequest))
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		return rpcResponse, nil
	// 	}
	// 	protoBuf, err := proto.Marshal(response)
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		return rpcResponse, nil
	// 	}
	// 	rpcResponse = &grpcPb.RpcResponse{
	// 		Rid:         request.GetRid(),
	// 		Code:        response.Code,
	// 		Desc:        response.Desc,
	// 		MessageType: grpcPb.MessageTypeDeviceLoginResponse,
	// 		ProtoBuf:    protoBuf,
	// 		ConnId:      request.ConnId,
	// 	}

	// } else if request.MessageType == grpcPb.MessageTypeDeviceLoginRequest {

	// } else if request.MessageType == grpcPb.MessageTypeCreateMessageRequest {

	// 	response, err := CreateMessage(ctx, protoMessage.(*grpcPb.CreateMessageRequest))
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		return rpcResponse, nil
	// 	}
	// 	protoBuf, err := proto.Marshal(response)
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 		return rpcResponse, nil
	// 	}
	// 	rpcResponse = &grpcPb.RpcResponse{
	// 		Rid:         request.GetRid(),
	// 		Code:        response.Code,
	// 		Desc:        response.Desc,
	// 		MessageType: grpcPb.MessageTypeCreateMessageResponse,
	// 		ProtoBuf:    protoBuf,
	// 		ConnId:      request.ConnId,
	// 	}

	// } else if request.MessageType == grpcPb.MessageTypeCreateSessionRequest {

	// }

	return rpcResponse, nil
}
