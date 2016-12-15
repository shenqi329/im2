package main

import (
	logicserverGrpc "im/logicserver/grpc"
	grpcPb "im/logicserver/grpc/pb"
	"im/logicserver/server"
	"log"
	//"runtime"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	s := server.NEWServer()

	grpcPb.RegisterRegisteServer(s.GrpcServer(), &logicserverGrpc.Registe{})
	grpcPb.RegisterLoginServer(s.GrpcServer(), &logicserverGrpc.Login{})
	grpcPb.RegisterSessionServer(s.GrpcServer(), &logicserverGrpc.Session{})
	grpcPb.RegisterMessageServer(s.GrpcServer(), &logicserverGrpc.Message{})
	grpcPb.RegisterRpcServer(s.GrpcServer(), &logicserverGrpc.Rpc{})

	s.Run("localhost:6005")
}
