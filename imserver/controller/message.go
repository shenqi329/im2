package controller

// import (
// 	grpcPb "im/grpc/pb"
// 	"im/imserver"
// 	"im/imserver/service"
// 	"log"
// )

// func HandleCreateMessage(c imserver.Context) error {

// 	request, ok := c.ProtoMessage().(*grpcPb.CreateMessageRequest)

// 	if !ok {
// 		return nil
// 	}

// 	tokenBean, err := service.HandleCreateMessage(c, request)

// 	if err != nil {
// 		log.Println(err)
// 		return c.SendProtoMessage(grpcPb.MessageTypeCreateMessageReply, imserver.NewCommonResponseWithError(err, request.Rid))
// 	}

// 	return c.SendProtoMessage(grpcPb.MessageTypeCreateMessageReply, tokenBean)
// }
