package controller

import (
	"im/imserver"
	"im/imserver/service"
	protocolClient "im/protocol/client"
	"log"
)

func HandleCreateMessage(c imserver.Context) error {

	request, ok := c.ProtoMessage().(*protocolClient.CreateMessageRequest)

	if !ok {
		return nil
	}

	tokenBean, err := service.HandleCreateMessage(c, request)

	if err != nil {
		log.Println(err)
		return c.SendProtoMessage(protocolClient.MessageTypeCreateMessageResponse, imserver.NewCommonResponseWithError(err, request.Rid))
	}

	return c.SendProtoMessage(protocolClient.MessageTypeCreateMessageResponse, tokenBean)
}
