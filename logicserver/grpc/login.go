package grpc

// import (
// 	"golang.org/x/net/context"
// 	grpcPb "im/grpc/pb"
// 	imserverError "im/logicserver/error"
// 	"im/logicserver/service"
// 	"im/logicserver/util/key"
// 	"log"
// )

// type Login struct{}

// func (m *Message) Login(ctx context.Context, request *grpcPb.DeviceLoginRequest) (*grpcPb.DeviceLoginResponse, error) {

// 	return LoginIn(ctx, request)

// }

// func LoginIn(ctx context.Context, request *grpcPb.DeviceLoginRequest) (*grpcPb.DeviceLoginResponse, error) {
// 	log.Println("Login")

// 	tokenConnInfoChan := ctx.Value(key.TokenConnInfoChan).(chan int64)

// 	protoMessage, err := service.HandleLogin(request)

// 	if err != nil {
// 		log.Println(err.Error())
// 		reply := &grpcPb.CreateMessageReply{
// 			Rid:  request.GetRid(),
// 			Code: imserverError.CommonInternalServerError,
// 			Desc: imserverError.ErrorCodeToText(imserverError.CommonInternalServerError),
// 		}
// 		return reply, nil
// 	}
// 	return protoMessage, nil
// }
