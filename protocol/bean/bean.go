package bean

import (
	proto "github.com/golang/protobuf/proto"
)

type MessageType int32

//系统保留类型 1-10
const (
	MessageTypeWraper = 1
)

//使用类型 11-
const (
	MessageTypeDeviceRegisteRequest  = 11
	MessageTypeDeviceRegisteResponse = 12
	MessageTypeDeviceLoginRequest    = 13
	MessageTypeDeviceLoginResponse   = 14
)

var kinds = map[MessageType]func() proto.Message{
	MessageTypeWraper:                func() proto.Message { return &WraperMessage{} },
	MessageTypeDeviceRegisteRequest:  func() proto.Message { return &DeviceRegisteRequest{} },
	MessageTypeDeviceRegisteResponse: func() proto.Message { return &DeviceRegisteResponse{} },
	MessageTypeDeviceLoginRequest:    func() proto.Message { return &DeviceLoginRequest{} },
	MessageTypeDeviceLoginResponse:   func() proto.Message { return &DeviceLoginResponse{} },
}

func Factory(messageType MessageType) proto.Message {

	createFunc := kinds[messageType]
	if createFunc != nil {
		return createFunc()
	}
	return nil
}