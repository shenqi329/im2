package controller

import (
	"im/imserver"
	"im/imserver/service"
	protocolClient "im/protocol/client"
	"log"
)

func HandleLogin(c imserver.Context) error {

	request, ok := c.ProtoMessage().(*protocolClient.DeviceLoginRequest)

	if !ok {
		return nil
	}

	tokenBean, err := service.HandleLogin(c, request)

	if err != nil {
		log.Println(err)
		return c.SendProtoMessage(protocolClient.MessageTypeDeviceRegisteResponse, imserver.NewCommonResponseWithError(err, request.Rid))
	}

	return c.SendProtoMessage(protocolClient.MessageTypeDeviceRegisteResponse, tokenBean)
}
