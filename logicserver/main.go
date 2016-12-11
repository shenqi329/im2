package main

import (
	grpcPb "im/grpc/pb"
	"im/logicserver/controller"
	logicserverGrpc "im/logicserver/grpc"
	"im/logicserver/server"
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

	grpcPb.RegisterSessionServer(s.GrpcServer(), &logicserverGrpc.Session{})
	grpcPb.RegisterMessageServer(s.GrpcServer(), &logicserverGrpc.Message{})
	protocolClient.RegisterRpcServer(s.GrpcServer(), &logicserverGrpc.Rpc{})

	s.Run(localUdpAddr, "localhost:6005")
}
