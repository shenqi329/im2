package controller

import (
	"im/imserver"
	"im/imserver/service"
	protocolBean "im/protocol/bean"
	"log"
)

func HandleCreateSession(c imserver.Context) error {

	request, ok := c.ProtoMessage().(*protocolBean.CreateSessionRequest)

	if !ok {
		return nil
	}

	tokenBean, err := service.HandleCreateSession(request)

	if err != nil {
		log.Println(err)
		return c.SendProtoMessage(protocolBean.MessageTypeCreateSessionResponse, imserver.NewCommonResponseWithError(err, request.Rid))
	}

	return c.SendProtoMessage(protocolBean.MessageTypeCreateSessionResponse, tokenBean)
}
