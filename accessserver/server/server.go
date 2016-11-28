package server

import (
	"github.com/golang/protobuf/proto"
	bean "im/protocal/bean"
	coder "im/protocal/coder"
	"log"
	"net"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

type Request struct {
	isCancel bool
	reqPkg   []byte
	rspChan  chan<- []byte
	rid      uint64
	conn     *net.TCPConn
}

type Server struct {
	rid       uint64 //请求流水号
	ridMutex  sync.Mutex
	linkCount int32

	localTcpAddr string
	proxyUdpAddr string
}

func (s *Server) createRID() uint64 {
	s.ridMutex.Lock()
	s.rid++
	s.ridMutex.Unlock()
	return s.rid
}

func NEWServer(localTcpAddr string, proxyUdpAddr string) (s *Server) {

	return &Server{
		localTcpAddr: localTcpAddr,
		proxyUdpAddr: proxyUdpAddr,
	}
}

func (s *Server) Run() {

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
	go s.connectProxyServer(reqChan)

	//
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			log.Println("accept tcp fail", err.Error())
			continue
		}
		go s.handleTcpConnection(conn, reqChan)
	}
}

func (s *Server) connectProxyServer(reqChan <-chan *Request) {
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

	reqMap := make(map[uint64]*Request)
	for {
		select {
		case req := <-reqChan:
			//log.Println("收到转发请求")
			sendChan <- req.reqPkg
			reqMap[req.rid] = req

		case rsp := <-recvChan:
			decoder := coder.NEWDecoder()
			beanWraperMessages, err := decoder.Decode(rsp)
			if err != nil {
				log.Println(err.Error())
				return
			}
			for _, beanWraperMessage := range beanWraperMessages {
				if beanWraperMessage.Type != bean.MessageTypeWraper {
					continue
				}
				protoWraperMessage := &bean.WraperMessage{}
				err := proto.Unmarshal(beanWraperMessage.Body, protoWraperMessage)
				if err != nil {
					log.Println(err)
					continue
				}
				req := reqMap[protoWraperMessage.Rid]
				if req != nil {
					delete(reqMap, protoWraperMessage.Rid)
					req.conn.Write(protoWraperMessage.Message)
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

func (s *Server) handleTcpConnection(conn *net.TCPConn, reqChan chan<- *Request) {

	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(10 * time.Second)

	decoder := coder.NEWDecoder()

	atomic.AddInt32(&s.linkCount, 1)
	log.Println("linkCount=", s.linkCount)

	defer func() {
		atomic.AddInt32(&s.linkCount, -1)
		log.Println("linkCount=", s.linkCount)
		conn.Close()
	}()

	buf := make([]byte, 1024)
	for true {
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
			go s.handleMessage(conn, message, reqChan)
		}
	}
}

func (s *Server) handleMessage(conn *net.TCPConn, message *coder.Message, reqChan chan<- *Request) {

	protoMessage := bean.Factory((bean.MessageType)(message.Type))

	request := &bean.DeviceRegisteRequest{}
	proto.Unmarshal(message.Body, request)

	if protoMessage == nil {
		log.Println("未识别的消息")
		conn.Close()
		return
	}

	if err := proto.Unmarshal(message.Body, protoMessage); err != nil {
		log.Println(err.Error())
		conn.Close()
		return
	}

	//只检查消息的合法性,然后将消息转发出去
	s.transformMessage(conn, message, reqChan)
}

func (s *Server) transformMessage(conn *net.TCPConn, message *coder.Message, reqChan chan<- *Request) {

	resChan := make(chan []byte, 1)
	//发送打包后的数据,数据中包含流水号
	rid := s.createRID()
	wraperMessage := &bean.WraperMessage{
		Rid:     rid,
		Message: message.Encode(),
	}
	buffer, err := coder.EncoderProtoMessage(bean.MessageTypeWraper, wraperMessage)
	if err != nil {
		log.Println(err)
	}
	reqChan <- &Request{
		isCancel: false,
		reqPkg:   buffer,
		rspChan:  resChan,
		rid:      rid,
		conn:     conn,
	}
}
