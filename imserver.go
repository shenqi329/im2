package main

import (
	"im/imserver"
	controller "im/imserver/controller"
	protocolBean "im/protocol/bean"
	"log"
	"net"
	"runtime"

	//"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "im/grpc/message"
	imserverGrpc "im/imserver/grpc"
)

func DeviceLogin(c imserver.Context) error {
	return nil
}

func grpcRegister() {
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

	grpcRegister()

	var localUdpAddr string
	if runtime.GOOS == "windows" {
		localUdpAddr = "localhost:6001"
	} else {
		localUdpAddr = "localhost:6001"
	}

	s := imserver.NEWServer(localUdpAddr)

	s.Handle(protocolBean.MessageTypeDeviceRegisteRequest, controller.HandleRegiste)
	s.Handle(protocolBean.MessageTypeDeviceLoginRequest, controller.HandleLogin)
	s.Handle(protocolBean.MessageTypeCreateSessionRequest, controller.HandleCreateSession)
	s.Handle(protocolBean.MessageTypeCreateMessageRequest, controller.HandleCreateMessage)

	s.Run()
}
