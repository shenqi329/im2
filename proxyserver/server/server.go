package server

import (
	"github.com/golang/protobuf/proto"
	"im/protocal/bean"
	"im/protocal/coder"
	"log"
	"net"
	"os"
)

func ServerUDPAddr() string {
	return "localhost:6001"
}

type Request struct {
	isCancel bool
	reqPkg   []byte
	rspChan  chan<- []byte
}

func ListenOnPort() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)
	addr, err := net.ResolveUDPAddr("udp", ServerUDPAddr())

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
	//
	reqChan := make(chan *Request, 1000)

	for true {
		buf := make([]byte, 4096)
		rlen, remote, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("读取数据失败!", err.Error())
			continue
		}
		go processHandler(conn, remote, buf[:rlen], reqChan)
	}
}

func processHandler(conn *net.UDPConn, remote *net.UDPAddr, msg []byte, reqChan chan<- *Request) {

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
			handleMessage(conn, remote, beanMessage, reqChan, wraperMessage.Rid)
		}
	}
}

func handleMessage(listen *net.UDPConn, addr *net.UDPAddr, message *coder.Message, reqChan chan<- *Request, rid int64) {

	protoMessage := bean.Factory((bean.MessageType)(message.Type))
	if err := proto.Unmarshal(message.Body, protoMessage); err != nil {
		log.Println(err.Error())
		return
	}
	switch message.Type {
	case bean.MessageTypeDeviceRegisteRequest:
		{
			handleRegisterRequest(listen, addr, protoMessage.(*bean.DeviceRegisteRequest), reqChan, rid)
		}
	case bean.MessageTypeDeviceLoginRequest:
		{
			handleLoginRequest(listen, addr, protoMessage.(*bean.DeviceLoginRequest), reqChan, rid)
		}
	}
}

func handleRegisterRequest(listen *net.UDPConn, addr *net.UDPAddr, request *bean.DeviceRegisteRequest, reqChan chan<- *Request, rid int64) {

	response := &bean.DeviceRegisteResponse{
		Rid:   request.Rid,
		Code:  "00000001",
		Desc:  "success",
		Token: "a token from proxyserver",
	}
	buffer, err := coder.EncoderProtoMessage(bean.MessageTypeDeviceLoginResponse, response)
	if err != nil {
		log.Println(err)
		return
	}
	sendBack(listen, addr, reqChan, rid, buffer)
}

func sendBack(listen *net.UDPConn, addr *net.UDPAddr, reqChan chan<- *Request, rid int64, buffer []byte) {
	wraperMessage := &bean.WraperMessage{
		Rid:     rid,
		Message: buffer,
	}
	buffer, err := coder.EncoderProtoMessage(bean.MessageTypeWraper, wraperMessage)
	if err != nil {
		log.Println(err)
		return
	}
	listen.WriteTo(buffer, addr)
}

func handleLoginRequest(listen *net.UDPConn, addr *net.UDPAddr, request *bean.DeviceLoginRequest, reqChan chan<- *Request, rid int64) {

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
	sendBack(listen, addr, reqChan, rid, buffer)
}
