package controller

import (
	"im/imserver"
	"im/imserver/service"
	protocolBean "im/protocol/bean"
	"log"
)

func HandleRegiste(c imserver.Context) error {

	request, ok := c.ProtoMessage().(*protocolBean.DeviceRegisteRequest)

	if !ok {
		return nil
	}

	tokenBean, err := service.HandleRegiste(request)

	if err != nil {
		log.Println(err)
		return c.WraperProtoMessage(protocolBean.MessageTypeDeviceRegisteResponse, imserver.NewCommonResponseWithError(err, request.Rid))
	}

	return c.WraperProtoMessage(protocolBean.MessageTypeDeviceRegisteResponse, tokenBean)
}
