package main

import (
	"github.com/golang/protobuf/proto"
	client "im/accessserver/client/client"
	grpcPb "im/grpc/pb"
	"im/logicserver/uuid"
	protocolClient "im/protocol/client"
	"im/protocol/coder"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	c := client.Client{}

	c.SetAfterLogin(func(c *client.Client) {

		log.Println("登陆成功")
		pb := &grpcPb.CreateMessageRequest{
			Rid:       c.GetRid(),
			SessionId: 32,
			Type:      1,
			Id:        uuid.Rand().Hex(),
			Content:   "a message from push",
		}
		protoBuf, err := proto.Marshal(pb)

		request := &protocolClient.RpcRequest{
			Rid:         c.GetRid(),
			AppId:       "89897",
			Type:        protocolClient.RpcRequest_LogicServer,
			MessageType: grpcPb.MessageTypeCreateMessageRequest,
			ProtoBuf:    protoBuf,
		}
		buffer, err := coder.EncoderProtoMessage(protocolClient.MessageTypeRPCRequest, request)
		if err != nil {
			log.Println(err.Error())
		}
		c.Conn.Write(buffer)
	})

	c.LoginToAccessServer()
	time.Sleep(60 * time.Minute)
}
