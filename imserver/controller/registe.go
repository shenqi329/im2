package controller

import (
	"im/imserver"
	"im/imserver/service"
	protocolClient "im/protocol/client"
	"log"
)

func HandleRegiste(c imserver.Context) error {

	request, ok := c.ProtoMessage().(*protocolClient.DeviceRegisteRequest)

	if !ok {
		return nil
	}

	tokenBean, err := service.HandleRegiste(request)

	if err != nil {
		log.Println(err)
		return c.SendProtoMessage(protocolClient.MessageTypeDeviceRegisteResponse, imserver.NewCommonResponseWithError(err, request.Rid))
	}

	return c.SendProtoMessage(protocolClient.MessageTypeDeviceRegisteResponse, tokenBean)
}
