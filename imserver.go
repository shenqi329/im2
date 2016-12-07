package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "im/grpc/message"
	"im/imserver"
	controller "im/imserver/controller"
	imserverGrpc "im/imserver/grpc"
	protocolClient "im/protocol/client"
	"log"
	"net"
	"runtime"
)

func DeviceLogin(c imserver.Context) error {
	return nil
}

func grpcServerRegister() {
	lis, err := net.Listen("tcp", ":6005")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterMessageServer(s, &imserverGrpc.Message{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	go grpcServerRegister()

	var localUdpAddr string
	if runtime.GOOS == "windows" {
		localUdpAddr = "localhost:6001"
	} else {
		localUdpAddr = "localhost:6001"
	}

	s := imserver.NEWServer(localUdpAddr)

	s.Handle(protocolClient.MessageTypeDeviceRegisteRequest, controller.HandleRegiste)
	s.Handle(protocolClient.MessageTypeDeviceLoginRequest, controller.HandleLogin)
	s.Handle(protocolClient.MessageTypeCreateSessionRequest, controller.HandleCreateSession)
	s.Handle(protocolClient.MessageTypeCreateMessageRequest, controller.HandleCreateMessage)

	s.Run()
}
