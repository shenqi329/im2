package controller

import (
	"im/logicserver/server"
	"im/logicserver/service"
	protocolClient "im/protocol/client"
	"log"
)

func HandleLogin(c server.Context) error {

	request, ok := c.ProtoMessage().(*protocolClient.DeviceLoginRequest)

	if !ok {
		return nil
	}

	tokenBean, err := service.HandleLogin(request)

	if err != nil {
		log.Println(err)
		return c.SendProtoMessage(protocolClient.MessageTypeDeviceLoginResponse, server.NewCommonResponseWithError(err, request.Rid))
	}
	c.SetIsLogin(true)

	return c.SendProtoMessage(protocolClient.MessageTypeDeviceLoginResponse, tokenBean)
}
