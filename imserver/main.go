package main

import (
	grpcPb "im/grpc/pb"
	"im/imserver/controller"
	imserverGrpc "im/imserver/grpc"
	"im/imserver/server"
	protocolClient "im/protocol/client"
	"log"
	"runtime"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	var localUdpAddr string
	if runtime.GOOS == "windows" {
		localUdpAddr = "localhost:6001"
	} else {
		localUdpAddr = "localhost:6001"
	}

	s := server.NEWServer()

	s.Handle(protocolClient.MessageTypeDeviceRegisteRequest, controller.HandleRegiste)
	s.Handle(protocolClient.MessageTypeDeviceLoginRequest, controller.HandleLogin)

	grpcPb.RegisterSessionServer(s.GrpcServer(), &imserverGrpc.Session{})
	grpcPb.RegisterMessageServer(s.GrpcServer(), &imserverGrpc.Message{})

	s.Run(localUdpAddr, "localhost:6005")
}
