package imserver

import (
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	grpcMessage "im/grpc/message"
	imserverGrpc "im/imserver/grpc"
	protocolClient "im/protocol/client"
	coder "im/protocol/coder"
	protocolServer "im/protocol/server"
	"log"
	"net"
	"os"
)

type Request struct {
	isCancel bool
	reqPkg   []byte
	rspChan  chan<- []byte
}

type ConnInfo struct {
	IsLogin bool
	UdpConn *net.UDPConn
	UdpAddr *net.UDPAddr
	ConnId  uint64
	Token   int64
	UserId  string
}

type Server struct {
	localUdpAddr string
	handleFuncs  map[protocolClient.MessageType]func(c Context) error
	connInfos    map[uint64]*ConnInfo //connId
	tokenInfos   map[int64]*ConnInfo
}

func NEWServer(localUdpAddr string) *Server {
	return &Server{
		localUdpAddr: localUdpAddr,
		handleFuncs:  make(map[protocolClient.MessageType]func(c Context) error),
		connInfos:    make(map[uint64]*ConnInfo),
		tokenInfos:   make(map[int64]*ConnInfo),
	}
}

func (s *Server) Handle(messageType protocolClient.MessageType, handle func(c Context) error) {
	s.handleFuncs[messageType] = handle
}

func (s *Server) Run() {

	go s.grpcServerRegister(":6005")

	s.listenOnUdpPort(s.localUdpAddr)
}

func (s *Server) GetConnInfo(connId uint64) *ConnInfo {
	connInfo := s.connInfos[connId]
	if connInfo == nil {
		connInfo = &ConnInfo{
			IsLogin: false,
		}
		s.connInfos[connId] = connInfo
	}
	return connInfo
}

func (s *Server) grpcServerRegister(tcpAddr string) {

	clientConn := s.grpcConnToEasyNoteAddr("localhost:6006")

	lis, err := net.Listen("tcp", tcpAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	grpcMessage.RegisterMessageServer(grpcServer, &imserverGrpc.Message{ClientConn: clientConn})
	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *Server) grpcConnToEasyNoteAddr(grpcAddr string) *grpc.ClientConn {
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func (s *Server) listenOnUdpPort(localUdpAddr string) {

	addr, err := net.ResolveUDPAddr("udp", localUdpAddr)

	if err != nil {
		log.Println("net.ResolveUDPAddr fail.", err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", addr)
	defer conn.Close()

	if err != nil {
		log.Println("net.ListenUDP fail.", err)
		os.Exit(1)
	}

	log.Println("net.ListenUDP", addr)

	var recvAndSendCount uint32 = 0
	connInfoChan := make(chan *ConnInfo, 1000)
	tokenConnInfoChan := make(chan int64, 1000)
	go s.syncData(tokenConnInfoChan, connInfoChan)
	for true {
		buf := make([]byte, 1024)
		rlen, remote, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("读取数据失败!", err.Error())
			continue
		}
		recvAndSendCount++
		log.Println("recvAndSendCount:", recvAndSendCount, " rlen:", rlen)
		go s.processHandler(tokenConnInfoChan, connInfoChan, conn, remote, buf[:rlen])
	}
}

func (s *Server) syncData(tokenConnInfoChan <-chan int64, connInfoChan <-chan *ConnInfo) {
	for {
		select {
		case connInfo := <-connInfoChan:
			{
				log.Println("UserId", connInfo.UserId, "token", connInfo.Token)
				s.connInfos[connInfo.ConnId] = connInfo
				s.tokenInfos[connInfo.Token] = connInfo
			}
		case tokenId := <-tokenConnInfoChan:
			{
				log.Println(tokenId)
				connInfo := s.tokenInfos[tokenId]
				if connInfo != nil {
					SendSyncInform(connInfo.UdpAddr, connInfo.UdpConn, connInfo.ConnId, connInfo.UserId)
					log.Println("UserId", connInfo.UserId, "token", connInfo.Token)
				}
			}

		}
	}

}

func (s *Server) processHandler(tokenConnInfoChan chan<- int64, connInfoChan chan<- *ConnInfo, conn *net.UDPConn, remote *net.UDPAddr, msg []byte) {

	decoder := coder.NEWDecoder()
	messages, err := decoder.Decode(msg)
	if err != nil {
		log.Println(err.Error())
		return
	}
	//处理
	for _, message := range messages {
		if message.Type == protocolServer.MessageTypeWraper {
			wraperMessage := &protocolServer.WraperMessage{}
			if err := proto.Unmarshal(message.Body, wraperMessage); err != nil {
				log.Println(err)
				continue
			}

			decoder.Reset()
			beanMessages, err := decoder.Decode(wraperMessage.Message)
			if err != nil {
				log.Println(err.Error())
				continue
			}
			for _, beanMessage := range beanMessages {
				s.handleMessage(tokenConnInfoChan, connInfoChan, conn, remote, beanMessage, wraperMessage.ConnId, true)
			}
		} else {
			s.handleMessage(tokenConnInfoChan, connInfoChan, conn, remote, message, 0, false)
		}
	}
}

func (s *Server) handleMessage(tokenConnInfoChan chan<- int64, connInfoChan chan<- *ConnInfo, conn *net.UDPConn, addr *net.UDPAddr, message *coder.Message, connId uint64, needWraper bool) {

	protoMessage := protocolClient.Factory((protocolClient.MessageType)(message.Type))
	if err := proto.Unmarshal(message.Body, protoMessage); err != nil {
		log.Println(err.Error())
		return
	}

	c := &context{
		imServer:          s,
		udpConn:           conn,
		udpAddr:           addr,
		protoMessage:      protoMessage,
		connId:            connId,
		needWraper:        needWraper,
		connInfoChan:      connInfoChan,
		tokenConnInfoChan: tokenConnInfoChan,
	}

	f := s.handleFuncs[(protocolClient.MessageType)(message.Type)]
	if f == nil {
		log.Println("不处理")
		return
	}

	//处理收到的数据
	if err := f(c); err != nil {
		log.Println(err)
		return
	}
}
