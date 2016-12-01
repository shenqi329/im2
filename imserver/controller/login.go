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

	tokenBean, err := service.HandleLogin(request)

	if err != nil {
		log.Println(err)
		return err
	}

	return c.WraperProtoMessage(protocolBean.MessageTypeDeviceRegisteResponse, tokenBean)
}
