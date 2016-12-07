package bean

import (
	proto "github.com/golang/protobuf/proto"
	//"log"
	//"reflect"
)

type MessageType int32

const (
	MessageTypeWraper = 1
)

var kinds = map[MessageType]func() proto.Message{
	MessageTypeWraper: func() proto.Message { return &WraperMessage{} },
}

func Factory(messageType MessageType) proto.Message {

	createFunc := kinds[messageType]
	if createFunc != nil {
		return createFunc()
	}
	return nil
}
