package controller

import (
	"im/imserver"
	"im/imserver/service"
	protocolClient "im/protocol/client"
	"log"
)

func HandleRpc(c imserver.Context) error {

	request, ok := c.ProtoMessage().(*protocolClient.RpcRequest)

	if !ok {
		return nil
	}

	tokenBean, err := service.HandleRpc(request)

	if err != nil {
		log.Println(err)
		return c.SendProtoMessage(protocolClient.MessageTypeRPCResponse, imserver.NewCommonResponseWithError(err, request.Rid))
	}

	return c.SendProtoMessage(protocolClient.MessageTypeRPCResponse, tokenBean)
}
