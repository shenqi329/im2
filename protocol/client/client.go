package client

import (
	proto "github.com/golang/protobuf/proto"
)

type MessageType int32

const (
	MessageTypeDeviceRegisteRequest  = 1
	MessageTypeDeviceRegisteResponse = 2
	MessageTypeDeviceLoginRequest    = 3
	MessageTypeDeviceLoginResponse   = 4
	MessageTypeSyncInform            = 5
	MessageTypeRPCRequest            = 6
	MessageTypeRPCResponse           = 7

	MessageTypeCreateSessionRequest  = 8
	MessageTypeCreateSessionResponse = 9
	MessageTypeCreateMessageRequest  = 10
	MessageTypeCreateMessageResponse = 11
)

var kinds = map[MessageType]func() proto.Message{
	MessageTypeCreateSessionRequest:  func() proto.Message { return &CreateSessionRequest{} },
	MessageTypeCreateSessionResponse: func() proto.Message { return &CreateSessionResponse{} },
	MessageTypeCreateMessageRequest:  func() proto.Message { return &CreateMessageRequest{} },
	MessageTypeCreateMessageResponse: func() proto.Message { return &CreateMessageResponse{} },

	MessageTypeDeviceRegisteRequest:  func() proto.Message { return &DeviceRegisteRequest{} },
	MessageTypeDeviceRegisteResponse: func() proto.Message { return &DeviceRegisteResponse{} },
	MessageTypeDeviceLoginRequest:    func() proto.Message { return &DeviceLoginRequest{} },
	MessageTypeDeviceLoginResponse:   func() proto.Message { return &DeviceLoginResponse{} },
	MessageTypeSyncInform:            func() proto.Message { return &SyncInform{} },
}

func Factory(messageType MessageType) proto.Message {

	createFunc := kinds[messageType]
	if createFunc != nil {
		return createFunc()
	}
	return nil
}
