package controller

// import (
// 	grpcPb "im/grpc/pb"
// 	"im/imserver"
// 	"im/imserver/service"
// 	"log"
// )

// func HandleCreateSession(c imserver.Context) error {

// 	request, ok := c.ProtoMessage().(*grpcPb.CreateSessionRequest)

// 	if !ok {
// 		return nil
// 	}

// 	tokenBean, err := service.HandleCreateSession(request)

// 	if err != nil {
// 		log.Println(err)
// 		return c.SendProtoMessage(grpcPb.MessageTypeCreateSessionReply, imserver.NewCommonResponseWithError(err, request.Rid))
// 	}

// 	return c.SendProtoMessage(grpcPb.MessageTypeCreateSessionReply, tokenBean)
// }
