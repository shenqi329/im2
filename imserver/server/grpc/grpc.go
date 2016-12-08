package grpc

import (
	//"github.com/golang/protobuf/proto"
	netContext "golang.org/x/net/context"
	//protocolClient "im/protocol/client"
	//coder "im/protocol/coder"
	//protocolServer "im/protocol/server"
	"google.golang.org/grpc"
	//grpcPb "im/grpc/pb"
	"google.golang.org/grpc/reflection"
	imserverGrpc "im/imserver/grpc"
	"log"
	"net"
)

const (
	KeyClientConn = "clientConn"
)

func GrpcServerRegister(tcpAddr string) {

	clientConn := GrpcConnToEasyNoteAddr("localhost:6006")

	lis, err := net.Listen("tcp", tcpAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(func(ctx netContext.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.Println("设置环境变量")
		ctx = netContext.WithValue(ctx, KeyClientConn, clientConn)
		return handler(ctx, req)
	}))

	//grpcPb.RegisterMessageServer(grpcServer, &imserverGrpc.Message{})
	//grpcPb.RegisterSessionServer(grpcServer, &imserverGrpc.Session{})
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func GrpcConnToEasyNoteAddr(grpcAddr string) *grpc.ClientConn {
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}
