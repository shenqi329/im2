package accessserver

import (
	"github.com/golang/protobuf/proto"
	coder "im/protocol/coder"
	protocolServer "im/protocol/server"
	"log"
	"net"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

type Request struct {
	reqPkg  []byte
	connId  uint32
	conn    *net.TCPConn
	isLogin bool
}

type ConnectInfo struct {
	conn    *net.TCPConn
	isLogin bool
}

type Server struct {
	rid       uint64 //请求流水号
	ridMutex  sync.Mutex
	connCount int32
	connId    uint32 //请求的id

	localTcpAddr string
	proxyUdpAddr string

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

func (s *Server) Run() {
	if s.handle == nil {
		s.handle = Handle
	}
	s.ListenOnTcpPort(s.localTcpAddr)
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

	go s.connectProxyServer(reqChan, closeChan)

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

func (s *Server) connectProxyServer(reqChan <-chan *Request, closeChan <-chan uint32) {
	addr, err := net.ResolveUDPAddr("udp", s.proxyUdpAddr)

	if err != nil {
		log.Println("net.ResolveUDPAddr fail.", err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Println("net.DialUDP fail.", err)
		os.Exit(1)
	}
	log.Println("net.DialUDP", addr)
	defer conn.Close()

	sendChan := make(chan []byte, 1000)
	go s.sendHandler(conn, sendChan)

	recvChan := make(chan []byte, 1000)
	go s.recvHandler(conn, recvChan)

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
					connMap[req.connId] = connInfo
				}
				wraperMessage := &protocolServer.WraperMessage{
					ConnId:    (uint64)(req.connId),
					Message:   req.reqPkg,
					IsLoginIn: connInfo.isLogin,
				}
				reqPkg, err := coder.EncoderProtoMessage(protocolServer.MessageTypeWraper, wraperMessage)
				if err != nil {
					log.Println(err)
					req.conn.Close()
					if connMap[req.connId] != nil {
						delete(connMap, req.connId)
					}
				}
				sendChan <- reqPkg
			}
		case rsp := <-recvChan:
			decoder := coder.NEWDecoder()
			beanWraperMessages, err := decoder.Decode(rsp)
			if err != nil {
				log.Println(err.Error())
				return
			}
			for _, beanWraperMessage := range beanWraperMessages {
				if beanWraperMessage.Type != protocolServer.MessageTypeWraper {
					continue
				}
				protoWraperMessage := &protocolServer.WraperMessage{}
				err := proto.Unmarshal(beanWraperMessage.Body, protoWraperMessage)
				if err != nil {
					log.Println(err)
					continue
				}
				connInfo := connMap[(uint32)(protoWraperMessage.ConnId)]
				if protoWraperMessage.IsLoginIn {
					connInfo.isLogin = true
				} else if protoWraperMessage.IsLoginOut {
					connInfo.isLogin = false
				}
				if connInfo != nil {
					connInfo.conn.Write(protoWraperMessage.Message)
				}
			}
		}
	}
}

func (s *Server) sendHandler(conn *net.UDPConn, sendChan <-chan []byte) {
	for data := range sendChan {
		//log.Println("处理转发请求")
		wlen, err := conn.Write(data)
		if err != nil || wlen != len(data) {
			log.Println("conn.Write fail.", err)
			continue
		}
	}
}

func (s *Server) recvHandler(conn *net.UDPConn, recvChan chan<- []byte) {
	for {
		buf := make([]byte, 4096)
		rlen, err := conn.Read(buf)
		if err != nil || rlen <= 0 {
			log.Println(err)
			continue
		}
		recvChan <- buf[:rlen]
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
