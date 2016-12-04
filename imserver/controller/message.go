package controller

import (
	"im/imserver"
	"im/imserver/service"
	protocolBean "im/protocol/bean"
	"log"
)

func HandleCreateMessage(c imserver.Context) error {

	request, ok := c.ProtoMessage().(*protocolBean.CreateMessageRequest)

	if !ok {
		return nil
	}

	tokenBean, err := service.HandleCreateMessage(c, request)

	if err != nil {
		log.Println(err)
		return c.SendProtoMessage(protocolBean.MessageTypeCreateMessageResponse, imserver.NewCommonResponseWithError(err, request.Rid))
	}

	return c.SendProtoMessage(protocolBean.MessageTypeCreateMessageResponse, tokenBean)
}
