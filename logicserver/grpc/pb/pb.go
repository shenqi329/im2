package pb

import (
	proto "github.com/golang/protobuf/proto"
)

type MessageType int32

const (
	MessageTypeDeviceRegisteRequest  = 1
	MessageTypeDeviceRegisteResponse = 2
	MessageTypeDeviceLoginRequest    = 3
	MessageTypeDeviceLoginResponse   = 4
	MessageTypeRPCRequest            = 5
	MessageTypeRPCResponse           = 6
	MessageTypeCreateMessageRequest  = 7
	MessageTypeCreateMessageResponse = 8

	MessageTypeCreateSessionRequest       = 10001
	MessageTypeCreateSessionResponse      = 10002
	MessageTypeAddSessionUsersRequest     = 10003
	MessageTypeAddSessionUsersResponse    = 10004
	MessageTypeDeleteSessionUsersRequest  = 10005
	MessageTypeDeleteSessionUsersResponse = 10006
)

var kinds = map[MessageType]func() proto.Message{
	MessageTypeDeviceRegisteRequest:  func() proto.Message { return &DeviceLoginRequest{} },
	MessageTypeDeviceRegisteResponse: func() proto.Message { return &DeviceLoginResponse{} },
	MessageTypeDeviceLoginRequest:    func() proto.Message { return &DeviceLoginRequest{} },
	MessageTypeDeviceLoginResponse:   func() proto.Message { return &DeviceLoginResponse{} },
	MessageTypeCreateMessageRequest:  func() proto.Message { return &CreateMessageRequest{} },
	MessageTypeCreateMessageResponse: func() proto.Message { return &CreateMessageResponse{} },
	MessageTypeCreateSessionRequest:  func() proto.Message { return &CreateSessionRequest{} },
	MessageTypeCreateSessionResponse: func() proto.Message { return &CreateSessionResponse{} },
	MessageTypeRPCRequest:            func() proto.Message { return &RpcRequest{} },
	MessageTypeRPCResponse:           func() proto.Message { return &RpcResponse{} },
}

func Factory(messageType MessageType) proto.Message {

	createFunc := kinds[messageType]
	if createFunc != nil {
		return createFunc()
	}
	return nil
}
