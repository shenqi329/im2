package server

import (
	"github.com/golang/protobuf/proto"
	bean "im/protocal/bean"
	coder "im/protocal/coder"
	//proxyServer "im/proxyserver/server"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

type Request struct {
	isCancel bool
	reqPkg   []byte
	rspChan  chan<- []byte
}

var linkCount int = 0
var mutex sync.Mutex

func ServerAddr() string {
	//return "192.168.0.107:6000"
	//return "172.17.0.2:6000"
	return "localhost:6000"
}

func ServerUDPAddr() string {
	return "localhost:6002"
}

func ListenUDPOnPort() {

	addr, err := net.ResolveUDPAddr("udp", ServerUDPAddr())

	if err != nil {
		log.Println("net.ResolveUDPAddr fail.", err)
		os.Exit(1)
	}

	listen, err := net.ListenUDP("udp", addr)
	defer listen.Close()

	if err != nil {
		log.Println("net.ListenUDP fail.", err)
		os.Exit(1)
	}

	buffer := make([]byte, 4096)
	decoder := coder.NEWDecoder()
	for true {
		decoder.Reset()

		count, udpAddr, err := listen.ReadFromUDP(buffer)
		if err != nil {
			log.Println("读取数据失败!", err.Error())
			continue
		}

		messages, err := decoder.Decode(buffer[0:count])
		if err != nil {
			log.Println(err.Error())
			continue
		}
		//处理
		for _, message := range messages {
			go handleUDPMessage(listen, udpAddr, message)
		}
	}
}

func handleUDPMessage(listen *net.UDPConn, addr *net.UDPAddr, message *coder.Message) {
	protoMessage := bean.Factory((bean.MessageType)(message.MessageType))

	if protoMessage == nil {
		log.Println("未识别的消息")
		return
	}

	if err := proto.Unmarshal(message.MessageBuf, protoMessage); err != nil {
		log.Println(err.Error())
		log.Println("消息格式错误")
		return
	}

	log.Println(proto.CompactTextString(protoMessage))

	switch message.MessageType {
	case bean.MessageTypeDeviceRegisteResponse:
		{
			handleRegisterResponse(listen, addr, protoMessage.(*bean.DeviceRegisteResponse))
		}
	case bean.MessageTypeDeviceLoginRequest:
		{
			handleLoginResponse(listen, addr, protoMessage.(*bean.DeviceLoginResponse))
		}
	}
}

func handleRegisterResponse(listen *net.UDPConn, addr *net.UDPAddr, request *bean.DeviceRegisteResponse) {

	// response := &bean.DeviceRegisteResponse{
	// 	Rid:   request.Rid,
	// 	Code:  "00000001",
	// 	Desc:  "success",
	// 	Token: "a token from proxyserver",
	// }

	// b, err := coder.EncoderProtoMessage(bean.MessageTypeDeviceRegisteResponse, response)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// listen.WriteTo(b, addr)
}

func handleLoginResponse(listen *net.UDPConn, addr *net.UDPAddr, request *bean.DeviceLoginResponse) {
	// response := &bean.DeviceLoginResponse{
	// 	Rid:  request.Rid,
	// 	Code: "00000001",
	// 	Desc: "success",
	// }

	// b, err := coder.EncoderProtoMessage(bean.MessageTypeDeviceLoginResponse, response)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// listen.WriteTo(b, addr)
}

func ListenOnPort() {

	//addr, err := net.ResolveTCPAddr("tcp", "172.17.0.2:6000")
	addr, err := net.ResolveTCPAddr("tcp", ServerAddr())

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

	reqChan := make(chan *Request, 1000)
	go connectUdpHandler(reqChan)

	go ListenUDPOnPort()

	log.Println("net.ListenTCP", addr)

	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			log.Println("accept tcp fail", err.Error())
			continue
		}
		go handleTcpConnection(conn, reqChan)
	}
}

func connectUdpHandler(reqChan <-chan *Request) {
	addr, err := net.ResolveUDPAddr("udp", "localhost:6001")

	if err != nil {
		log.Println("net.ResolveUDPAddr fail.", err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Println("net.DialUDP fail.", err)
		os.Exit(1)
	}
	defer conn.Close()

	sendChan := make(chan []byte, 1000)
	go sendHandler(conn, sendChan)

	recvChan := make(chan []byte, 1000)
	go recvHandler(conn, recvChan)

	//reqMap := make(map[int]*Request)

	for {
		select {
		case req := <-reqChan:
			sendChan <- req.reqPkg

		case rsp := <-recvChan:
			decoder := coder.NEWDecoder()
			messages, err := decoder.Decode(rsp)
			if err != nil {
				log.Println(err.Error())
				conn.Close()
				return
			}
			for _, message := range messages {
				protoMessage := bean.Factory((bean.MessageType)(message.MessageType))
				log.Println(proto.CompactTextString(protoMessage))
			}
		}
	}
}

func sendHandler(conn *net.UDPConn, sendChan <-chan []byte) {
	for data := range sendChan {
		wlen, err := conn.Write(data)
		if err != nil || wlen != len(data) {
			log.Println("conn.Write fail.", err)
			continue
		}
	}
}

func recvHandler(conn *net.UDPConn, recvChan chan<- []byte) {
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

func handleTcpConnection(conn *net.TCPConn, reqChan chan<- *Request) {

	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(10 * time.Second)

	decoder := coder.NEWDecoder()

	mutex.Lock()
	linkCount++
	log.Println("linkCount=", linkCount)
	mutex.Unlock()

	defer func() {
		mutex.Lock()
		linkCount--
		log.Println("linkCount=", linkCount)
		mutex.Unlock()
		conn.Close()
	}()

	buf := make([]byte, 512)
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
			handleMessage(conn, message, reqChan)
		}
	}
}

func handleMessage(conn *net.TCPConn, message *coder.Message, reqChan chan<- *Request) {

	protoMessage := bean.Factory((bean.MessageType)(message.MessageType))

	if protoMessage == nil {
		log.Println("未识别的消息")
		conn.Close()
		return
	}

	if err := proto.Unmarshal(message.MessageBuf, protoMessage); err != nil {
		log.Println(err.Error())
		log.Println("消息格式错误")
		conn.Close()
		return
	}
	log.Println(proto.CompactTextString(protoMessage))
	//只检查消息的合法性,然后将消息转发出去
	go transformMessage(conn, message, reqChan)
}

func transformMessage(conn *net.TCPConn, message *coder.Message, reqChan chan<- *Request) {

	resChan := make(chan []byte, 1)
	buffer, err := coder.EncoderMessage(message.MessageType, message.MessageBuf)
	if err != nil {
		log.Println(err)
		conn.Close()
	}
	reqChan <- &Request{
		isCancel: false,
		reqPkg:   buffer,
		rspChan:  resChan,
	}

	// raddr, err := net.ResolveUDPAddr("udp", proxyServer.ServerUDPAddr())

	// raddr, err := net.ResolveUDPAddr("udp", "localhost:6001")

	// if err != nil {
	// 	log.Println("net.ResolveUDPAddr fail.", err)
	// 	return
	// }

	// socker, err := net.DialUDP("udp", nil, raddr)
	// if err != nil {
	// 	log.Println("net.DialUDP fail.", err)
	// 	conn.Close()
	// 	return
	// }
	// //defer socker.Close()

	// b, err := coder.EncoderMessage(message.MessageType, message.MessageBuf)
	// if err != nil {
	// 	log.Println(err)
	// 	conn.Close()
	// }
	// socker.Write(b)

	// //接收数据
	// buffer := make([]byte, 4096)
	// count, _, err := socker.ReadFromUDP(buffer)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	conn.Close()
	// 	return
	// }
	// decoder := coder.NEWDecoder()
	// _, err = decoder.Decode(buffer[0:count])
	// if err != nil {
	// 	log.Println(err.Error())
	// 	conn.Close()
	// 	return
	// }

	// conn.Write(buffer[0:count])
}
