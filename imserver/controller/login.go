package controller

import (
	"im/imserver"
	"im/imserver/service"
	protocolBean "im/protocol/bean"
	"log"
)

func HandleLogin(c imserver.Context) error {

	request, ok := c.ProtoMessage().(*protocolBean.DeviceLoginRequest)

	if !ok {
		return nil
	}

	tokenBean, err := service.HandleLogin(c, request)

	if err != nil {
		log.Println(err)
		return c.WraperProtoMessage(protocolBean.MessageTypeDeviceRegisteResponse, imserver.NewCommonResponseWithError(err, request.Rid))
	}

	return c.WraperProtoMessage(protocolBean.MessageTypeDeviceRegisteResponse, tokenBean)
}
