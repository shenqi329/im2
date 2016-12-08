package pb

import (
	proto "github.com/golang/protobuf/proto"
)

type MessageType int32

const (
	MessageTypeCreateMessageRequest = 1
	MessageTypeCreateMessageReply   = 2
	MessageTypeCreateSessionRequest = 3
	MessageTypeCreateSessionReply   = 4
)

var kinds = map[MessageType]func() proto.Message{
	MessageTypeCreateMessageRequest: func() proto.Message { return &CreateMessageRequest{} },
	MessageTypeCreateMessageReply:   func() proto.Message { return &CreateMessageReply{} },
	MessageTypeCreateSessionRequest: func() proto.Message { return &CreateSessionRequest{} },
	MessageTypeCreateSessionReply:   func() proto.Message { return &CreateSessionReply{} },
}

func Factory(messageType MessageType) proto.Message {

	createFunc := kinds[messageType]
	if createFunc != nil {
		return createFunc()
	}
	return nil
}
