package controller

import (
	"im/imserver"
	"im/imserver/service"
	protocolClient "im/protocol/client"
	"log"
)

func HandleCreateSession(c imserver.Context) error {

	request, ok := c.ProtoMessage().(*protocolClient.CreateSessionRequest)

	if !ok {
		return nil
	}

	tokenBean, err := service.HandleCreateSession(request)

	if err != nil {
		log.Println(err)
		return c.SendProtoMessage(protocolClient.MessageTypeCreateSessionResponse, imserver.NewCommonResponseWithError(err, request.Rid))
	}

	return c.SendProtoMessage(protocolClient.MessageTypeCreateSessionResponse, tokenBean)
}
