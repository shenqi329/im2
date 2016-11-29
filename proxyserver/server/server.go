package server

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"im/protocal/bean"
	"im/protocal/coder"
	"log"
	"net"
	"os"
)

type Request struct {
	isCancel bool
	reqPkg   []byte
	rspChan  chan<- []byte
}

type Server struct {
	localUdpAddr string
}

func NEWServer(localUdpAddr string) *Server {
	return &Server{
		localUdpAddr: localUdpAddr,
	}
}

func (s *Server) Run() {
	s.listenOnUdpPort(s.localUdpAddr)
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

	conn.SetReadBuffer(1024 * 1024 * 100)
	conn.SetWriteBuffer(1024 * 1024 * 100)
	log.Println("net.ListenUDP", addr)
	//
	reqChan := make(chan *Request, 1000)
	var recvAndSendCount uint32 = 0

	for true {
		buf := make([]byte, 1024)
		rlen, remote, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("读取数据失败!", err.Error())
			continue
		}
		recvAndSendCount++
		//log.Println("recvAndSendCount:", recvAndSendCount, " rlen:", rlen)
		go s.processHandler(conn, remote, buf[:rlen], reqChan)
	}
}

func (s *Server) processHandler(conn *net.UDPConn, remote *net.UDPAddr, msg []byte, reqChan chan<- *Request) {

	decoder := coder.NEWDecoder()
	beanWraperMessages, err := decoder.Decode(msg)
	if err != nil {
		log.Println(err.Error())
		return
	}
	//处理
	for _, beanWraperMessage := range beanWraperMessages {
		if beanWraperMessage.Type != bean.MessageTypeWraper {
			continue
		}

		wraperMessage := &bean.WraperMessage{}
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
			s.handleMessage(conn, remote, beanMessage, reqChan, wraperMessage.Rid)
		}
	}
}

func (s *Server) handleMessage(conn *net.UDPConn, addr *net.UDPAddr, message *coder.Message, reqChan chan<- *Request, rid uint64) {

	protoMessage := bean.Factory((bean.MessageType)(message.Type))
	if err := proto.Unmarshal(message.Body, protoMessage); err != nil {
		log.Println(err.Error())
		return
	}
	switch message.Type {
	case bean.MessageTypeDeviceRegisteRequest:
		{
			s.handleRegisterRequest(conn, addr, protoMessage.(*bean.DeviceRegisteRequest), reqChan, rid)
		}
	case bean.MessageTypeDeviceLoginRequest:
		{
			s.handleLoginRequest(conn, addr, protoMessage.(*bean.DeviceLoginRequest), reqChan, rid)
		}
	}
}

func (s *Server) handleRegisterRequest(conn *net.UDPConn, addr *net.UDPAddr, request *bean.DeviceRegisteRequest, reqChan chan<- *Request, rid uint64) {

	response := &bean.DeviceRegisteResponse{
		Rid:   request.Rid,
		Code:  "00000001",
		Desc:  "success",
		Token: fmt.Sprintf("%d", rid),
	}
	buffer, err := coder.EncoderProtoMessage(bean.MessageTypeDeviceRegisteResponse, response)
	if err != nil {
		log.Println(err)
		return
	}
	s.wraperMessageAndSendBack(conn, addr, reqChan, rid, buffer)
}

func (s *Server) handleLoginRequest(conn *net.UDPConn, addr *net.UDPAddr, request *bean.DeviceLoginRequest, reqChan chan<- *Request, rid uint64) {

	response := &bean.DeviceLoginResponse{
		Rid:  request.Rid,
		Code: "00000001",
		Desc: "success",
	}

	buffer, err := coder.EncoderProtoMessage(bean.MessageTypeDeviceLoginResponse, response)
	if err != nil {
		log.Println(err)
		return
	}
	s.wraperMessageAndSendBack(conn, addr, reqChan, rid, buffer)
}

func (s *Server) wraperMessageAndSendBack(conn *net.UDPConn, addr *net.UDPAddr, reqChan chan<- *Request, rid uint64, buffer []byte) {
	wraperMessage := &bean.WraperMessage{
		Rid:     rid,
		Message: buffer,
	}
	buffer, err := coder.EncoderProtoMessage(bean.MessageTypeWraper, wraperMessage)
	if err != nil {
		log.Println(err)
		return
	}

	conn.WriteTo(buffer, addr)
}
