package server

import (
	netContext "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	logicserverError "im/logicserver/error"
	logicserverGrpc "im/logicserver/grpc"
	grpcPb "im/logicserver/grpc/pb"
	"log"
	"net"
	"reflect"
)

type Server struct {
	grpcServer *grpc.Server
}

func NEWServer() *Server {
	s := &Server{}
	return s
}

func (s *Server) GrpcServer() *grpc.Server {
	if s.grpcServer == nil {
		s.grpcServer = s.newServer()
	}
	return s.grpcServer
}

func (s *Server) Run(grpcTcpPort string) {
	s.grpcServerServe(grpcTcpPort)
}

func (s *Server) newServer() *grpc.Server {

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(func(ctx netContext.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		//log.Println("设置环境变量")
		response, err := handler(ctx, req)
		v := reflect.ValueOf(response)
		if !v.IsValid() || v.IsNil() {
			request, ok := req.(logicserverGrpc.Request)
			if ok {
				response = &grpcPb.Response{
					Rid:  request.GetRid(),
					Code: logicserverError.CommonInternalServerError,
					Desc: logicserverError.ErrorCodeToText(logicserverError.CommonInternalServerError),
				}
			} else {
				response = &grpcPb.Response{
					Code: logicserverError.CommonInternalServerError,
					Desc: logicserverError.ErrorCodeToText(logicserverError.CommonInternalServerError),
				}
			}
		}
		return response, nil
	}))
	return grpcServer
}

func (s *Server) grpcServerServe(addr string) {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("grpcServerServe addr:", addr)

	reflection.Register(s.grpcServer)
	if err := s.grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
