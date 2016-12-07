package service

import (
	"github.com/golang/protobuf/proto"
	"im/imserver/bean"
	protocolClient "im/protocol/client"
	"log"
)

func HandleRpc(request *protocolClient.RpcRequest) (proto.Message, error) {

	log.Println(bean.StructToJsonString(request))

	return nil, nil
}
