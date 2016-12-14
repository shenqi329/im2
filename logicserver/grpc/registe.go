package grpc

import (
	"golang.org/x/net/context"
	imserverError "im/logicserver/error"
	grpcPb "im/logicserver/grpc/pb"
	"im/logicserver/service"
	"log"
)

type Registe struct{}

func (r *Registe) Registe(ctx context.Context, request *grpcPb.DeviceRegisteRequest) (*grpcPb.DeviceRegisteResponse, error) {

	return HandleRegiste(ctx, request)

}

func HandleRegiste(ctx context.Context, request *grpcPb.DeviceRegisteRequest) (*grpcPb.DeviceRegisteResponse, error) {
	log.Println("Login")

	protoMessage, err := service.HandleRegiste(request)

	if err != nil {
		log.Println(err.Error())
		reply := &grpcPb.DeviceRegisteResponse{
			Rid:  request.GetRid(),
			Code: imserverError.CommonInternalServerError,
			Desc: imserverError.ErrorCodeToText(imserverError.CommonInternalServerError),
		}
		return reply, nil
	}
	return protoMessage, nil
}
