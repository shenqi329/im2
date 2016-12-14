package server

import (
	"github.com/golang/protobuf/proto"
	netContext "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	grpcPb "im/logicserver/grpc/pb"
	coder "im/protocol/coder"
	"log"
	"net"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

type Request struct {
	message      *coder.Message
	protoMessage proto.Message
	messageType  grpcPb.MessageType
	connId       uint32
	conn         *net.TCPConn
}

type ConnectInfo struct {
	conn    *net.TCPConn
	isLogin bool
	token   string
	userId  string
}

type Server struct {
	rid       uint64 //请求流水号
	ridMutex  sync.Mutex
	connCount int32
	connId    uint32 //请求的id

	localTcpAddr string
	proxyUdpAddr string

	grpcEasynoteClientConn *grpc.ClientConn
	grpcLogicClientConn    *grpc.ClientConn
	grpcServer             *grpc.Server

	handle func(context Context) error
}

func (s *Server) createRID() uint64 {
	s.ridMutex.Lock()
	s.rid++
	s.ridMutex.Unlock()
	return s.rid
}

func (s *Server) createConnId() uint32 {
	return atomic.AddUint32(&s.connId, 1)
}

func NEWServer(localTcpAddr string, proxyUdpAddr string) (s *Server) {

	return &Server{
		localTcpAddr: localTcpAddr,
		proxyUdpAddr: proxyUdpAddr,
	}
}

func (s *Server) GrpcServer() *grpc.Server {
	if s.grpcServer == nil {
		s.grpcServer = s.newGrpcServer()
	}
	return s.grpcServer
}

func (s *Server) Run(grpcServerAddr string) {
	if s.handle == nil {
		s.handle = Handle
	}

	//grpcPb.RegisterRpcServer(s.GrpcServer(), &serverGrpc.Rpc{})

	go s.grpcServerServe(grpcServerAddr)
	s.grpcEasynoteClientConn = s.grpcConnectServer("localhost:6006")
	s.grpcLogicClientConn = s.grpcConnectServer("localhost:6005")
	s.ListenOnTcpPort(s.localTcpAddr)
}

func (s *Server) grpcConnectServer(tcpAddr string) *grpc.ClientConn {
	//Set up a connection to the server.
	conn, err := grpc.Dial(tcpAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func (s *Server) grpcServerServe(addr string) {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	reflection.Register(s.GrpcServer())
	if err := s.GrpcServer().Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) newGrpcServer() *grpc.Server {

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(func(ctx netContext.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		//log.Println("设置环境变量")
		//ctx = netContext.WithValue(ctx, "xxxxx", v)
		return handler(ctx, req)
	}))
	return grpcServer
}

func (s *Server) ListenOnTcpPort(localTcpAddr string) {

	addr, err := net.ResolveTCPAddr("tcp", localTcpAddr)

	if err != nil {
		log.Println("net.ResolveTCPAddr fail.", err)
		os.Exit(1)
	}

	listen, err := net.ListenTCP("tcp", addr)
	defer listen.Close()

	if err != nil {
		log.Println("net.ListenTCP fail.", err)
		os.Exit(1)
	}
	log.Println("net.ListenTCP", addr)

	//
	reqChan := make(chan *Request, 1000)
	closeChan := make(chan uint32, 1000)

	go s.connectIMServer(reqChan, closeChan)

	//
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			log.Println("accept tcp fail", err.Error())
			continue
		}

		go s.handleTcpConnection(conn, reqChan, closeChan)
	}
}

func (s *Server) transToLogicServer(rpcRequest *grpcPb.RpcRequest, protocolRespChan chan<- *ProtocolResp) {
	//log.Println("转发给逻辑服务器")
	rpcClient := grpcPb.NewRpcClient(s.grpcLogicClientConn)
	response, err := rpcClient.Rpc(netContext.Background(), rpcRequest)
	if err != nil {
		//直接返回错误给调用者[刘俊仕]
		log.Println(err.Error())
		return
	}

	protocolBuf, err := coder.EncoderMessage((int)(response.MessageType), response.ProtoBuf)
	if err != nil {
		log.Println(err.Error())
		return
	}
	protocolRespChan <- &ProtocolResp{
		protocolBuf: protocolBuf,
		connId:      (uint32)(rpcRequest.ConnId),
	}
}

func (s *Server) transToBusinessServer(rpcRequest *grpcPb.RpcRequest, rpcRespChan chan<- *grpcPb.RpcResponse) {
	//easynote业务id
	if rpcRequest.AppId == "89897" {
		//log.Println("转发给业务服务器")
		rpcClient := grpcPb.NewRpcClient(s.grpcEasynoteClientConn)
		response, err := rpcClient.Rpc(netContext.Background(), rpcRequest)
		if err != nil {
			//直接返回错误给调用者[刘俊仕]
			log.Println(err.Error())
			return
		}
		log.Println(response.String())
		rpcRespChan <- response
	}
}

type ProtocolResp struct {
	protocolBuf []byte
	connId      uint32
}

//连接到逻辑服务器
func (s *Server) connectIMServer(reqChan <-chan *Request, closeChan <-chan uint32) {

	rpcRespChan := make(chan *grpcPb.RpcResponse, 1000)
	protocolRespChan := make(chan *ProtocolResp, 1000)
	connMap := make(map[uint32]*ConnectInfo)

	for {
		select {
		case connId := <-closeChan:
			{
				if connMap[connId] != nil {
					delete(connMap, connId)
				}
			}

		case req := <-reqChan:
			{
				connInfo := connMap[req.connId]
				if connInfo == nil {
					connInfo = &ConnectInfo{
						conn:    req.conn,
						isLogin: false,
					}
					log.Println(req.messageType)
					if req.messageType == grpcPb.MessageTypeDeviceLoginRequest {
						loginRequest, ok := req.protoMessage.(*grpcPb.DeviceLoginRequest)
						log.Println(ok)
						if ok {
							log.Println(loginRequest.String())
							connInfo.token = loginRequest.Token
							connInfo.userId = loginRequest.UserId
						}
					}
					connMap[req.connId] = connInfo
				}
				if req.message.Type == grpcPb.MessageTypeRPCRequest {
					//转发给具体的业务服务器
					if !connInfo.isLogin {
						log.Println("没有登录,不转发消息")
						connInfo.conn.Close()
						delete(connMap, req.connId)
						break
					}
					rpcRequest, ok := req.protoMessage.(*grpcPb.RpcRequest)
					if !ok {
						break
					}
					rpcRequest.ConnId = (uint64)(req.connId)
					rpcRequest.UserId = connInfo.userId
					rpcRequest.Token = connInfo.token
					go s.transToBusinessServer(rpcRequest, rpcRespChan)
				} else {
					//转发给im逻辑服务器
					protoBuf, err := proto.Marshal(req.protoMessage)
					if err == nil {
						rpcRequest := &grpcPb.RpcRequest{}
						rpcRequest.ConnId = (uint64)(req.connId)
						rpcRequest.UserId = connInfo.userId
						rpcRequest.Token = connInfo.token
						rpcRequest.MessageType = (int32)(req.messageType)
						rpcRequest.ProtoBuf = protoBuf
						go s.transToLogicServer(rpcRequest, protocolRespChan)
					}
				}
			}
		case rpcResp := <-rpcRespChan:
			{
				connInfo := connMap[(uint32)(rpcResp.ConnId)]
				if connInfo == nil {
					break
				}
				buffer, err := coder.EncoderProtoMessage(grpcPb.MessageTypeRPCResponse, rpcResp)
				if err != nil {
					log.Println(err.Error())
				}
				connInfo.conn.Write(buffer)
			}
		case protocolBufChan := <-protocolRespChan:
			{
				connInfo := connMap[(uint32)(protocolBufChan.connId)]
				connInfo.conn.Write(protocolBufChan.protocolBuf)
			}
		}
	}
}

func (s *Server) handleTcpConnection(conn *net.TCPConn, reqChan chan<- *Request, closeChan chan<- uint32) {

	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(10 * time.Second)

	decoder := coder.NEWDecoder()

	atomic.AddInt32(&s.connCount, 1)
	connId := atomic.AddUint32(&s.connId, 1) //生成连接的id
	log.Println("connCount=", s.connCount)

	defer func() {
		atomic.AddInt32(&s.connCount, -1)
		log.Println("connCount=", s.connCount)
		conn.Close()
		closeChan <- connId
	}()

	for true {
		buf := make([]byte, 1024)
		count, err := conn.Read(buf)
		if err != nil {
			log.Println(err.Error())
			break
		}
		messages, err := decoder.Decode(buf[0:count])
		if err != nil {
			log.Println(err.Error())
			break
		}
		for _, message := range messages {
			request := &Request{
				connId: connId,
				conn:   conn,
			}
			//log.Println(message.Type)
			context := &context{
				reqChan:   reqChan,
				message:   message,
				closeChan: closeChan,
				conn:      conn,
				server:    s,
				request:   request,
			}
			if s.handle != nil {
				s.handle(context)
			}
		}
	}
}
