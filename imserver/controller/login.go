package controller

import (
	"im/imserver/server"
	"im/imserver/service"
	protocolClient "im/protocol/client"
	"log"
)

func HandleLogin(c server.Context) error {

	request, ok := c.ProtoMessage().(*protocolClient.DeviceLoginRequest)

	if !ok {
		return nil
	}

	tokenBean, err := service.HandleLogin(c, request)

	if err != nil {
		log.Println(err)
		return c.SendProtoMessage(protocolClient.MessageTypeDeviceRegisteResponse, server.NewCommonResponseWithError(err, request.Rid))
	}

	return c.SendProtoMessage(protocolClient.MessageTypeDeviceRegisteResponse, tokenBean)
}
