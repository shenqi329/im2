package grpc

import (
	"golang.org/x/net/context"
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

	return protoMessage, err
}
