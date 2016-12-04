package bean

import (
	proto "github.com/golang/protobuf/proto"
	//"log"
	//"reflect"
)

type MessageType int32

//系统保留类型 1-100
const (
	MessageTypeWraper                = 1
	MessageTypeCreateSessionRequest  = 2
	MessageTypeCreateSessionResponse = 3
	MessageTypeCreateMessageRequest  = 4
	MessageTypeCreateMessageResponse = 5
)

//使用类型 11-
const (
	MessageTypeDeviceRegisteRequest  = 101
	MessageTypeDeviceRegisteResponse = 102
	MessageTypeDeviceLoginRequest    = 103
	MessageTypeDeviceLoginResponse   = 104
	MessageTypeSyncInform            = 105
)

var kinds = map[MessageType]func() proto.Message{
	MessageTypeWraper:                func() proto.Message { return &WraperMessage{} },
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

// func MessageTypeFromMessage(message proto.Message) MessageType {

// 	typ := reflect.ValueOf(message).Type()
// 	log.Println(typ.PkgPath())
// 	switch typ.Kind() {
// 	case reflect.ValueOf(WraperMessage{}).Type().Kind():
// 		{
// 			return MessageTypeWraper
// 		}
// 	case reflect.ValueOf(DeviceRegisteRequest{}).Type().Kind():
// 		{
// 			return MessageTypeDeviceRegisteRequest
// 		}
// 	case reflect.ValueOf(DeviceRegisteResponse{}).Type().Kind():
// 		{
// 			return MessageTypeDeviceRegisteResponse
// 		}
// 	case reflect.ValueOf(DeviceLoginRequest{}).Type().Kind():
// 		{
// 			return MessageTypeDeviceLoginRequest
// 		}
// 	case reflect.ValueOf(DeviceLoginResponse{}).Type().Kind():
// 		{
// 			return MessageTypeDeviceLoginResponse
// 		}
// 	case reflect.ValueOf(SyncInform{}).Type().Kind():
// 		{
// 			return MessageTypeSyncInform
// 		}
// 	}

// 	return 0
// }
