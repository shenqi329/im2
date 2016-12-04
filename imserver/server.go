package imserver

import (
	//"errors"
	"github.com/golang/protobuf/proto"
	//imResponse "im/imserver/response"
	//imService "im/imserver/service"
	protocolBean "im/protocol/bean"
	"im/protocol/coder"
	"log"
	"net"
	"os"
	//"reflect"
)

type Request struct {
	isCancel bool
	reqPkg   []byte
	rspChan  chan<- []byte
}

type ConnInfo struct {
	IsLogin bool
}

type Server struct {
	localUdpAddr string
	handleFuncs  map[protocolBean.MessageType]func(c Context) error
	connInfos    map[uint64]*ConnInfo
}

func NEWServer(localUdpAddr string) *Server {
	return &Server{
		localUdpAddr: localUdpAddr,
		handleFuncs:  make(map[protocolBean.MessageType]func(c Context) error),
		connInfos:    make(map[uint64]*ConnInfo),
	}
}

func (s *Server) Handle(messageType protocolBean.MessageType, handle func(c Context) error) {
	s.handleFuncs[messageType] = handle
}

func (s *Server) Run() {
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

	for true {
		buf := make([]byte, 1024)
		rlen, remote, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("读取数据失败!", err.Error())
			continue
		}
		recvAndSendCount++
		log.Println("recvAndSendCount:", recvAndSendCount, " rlen:", rlen)
		go s.processHandler(conn, remote, buf[:rlen])
	}
}

func (s *Server) processHandler(conn *net.UDPConn, remote *net.UDPAddr, msg []byte) {

	decoder := coder.NEWDecoder()
	beanWraperMessages, err := decoder.Decode(msg)
	if err != nil {
		log.Println(err.Error())
		return
	}
	//处理
	for _, beanWraperMessage := range beanWraperMessages {
		if beanWraperMessage.Type != protocolBean.MessageTypeWraper {
			continue
		}

		wraperMessage := &protocolBean.WraperMessage{}
		if err := proto.Unmarshal(beanWraperMessage.Body, wraperMessage); err != nil {
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
			s.handleMessage(conn, remote, beanMessage, wraperMessage.ConnId)
		}
	}
}

func (s *Server) handleMessage(conn *net.UDPConn, addr *net.UDPAddr, message *coder.Message, connId uint64) {

	protoMessage := protocolBean.Factory((protocolBean.MessageType)(message.Type))
	if err := proto.Unmarshal(message.Body, protoMessage); err != nil {
		log.Println(err.Error())
		return
	}

	c := &context{
		imServer:     s,
		udpConn:      conn,
		udpAddr:      addr,
		protoMessage: protoMessage,
		connId:       connId,
	}

	f := s.handleFuncs[(protocolBean.MessageType)(message.Type)]
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
