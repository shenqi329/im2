package main

import (
	client "im/accessserver/client/client"
	grpcPb "im/logicserver/grpc/pb"
	"im/logicserver/uuid"
	"im/protocol/coder"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	c := client.Client{}

	c.SetAfterLogin(func(c *client.Client) {

		log.Println("登陆成功")
		for i := 0; i < 1; i++ {
			request := &grpcPb.CreateMessageRequest{
				Rid:       c.GetRid(),
				SessionId: 32,
				Type:      1,
				Id:        uuid.Rand().Hex(),
				Content:   "a message from push",
			}

			buffer, err := coder.EncoderProtoMessage(grpcPb.MessageTypeCreateMessageRequest, request)
			if err != nil {
				log.Println(err.Error())
			}
			c.Conn.Write(buffer)
		}
	})

	c.LoginToAccessServer()
	time.Sleep(60 * time.Minute)
}
