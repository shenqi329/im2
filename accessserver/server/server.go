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
	rid      int64
	conn     *net.TCPConn
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

func Start() {
	//go ListenUDPOnPort(ServerUDPAddr())
	ListenOnTcpPort(ServerAddr())
}

var gRid int64 = 0

func createRID() int64 {
	mutex.Lock()
	gRid++
	mutex.Unlock()
	return gRid
}

func ListenOnTcpPort(laddr string) {

	addr, err := net.ResolveTCPAddr("tcp", laddr)

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
	go connectProxyServer(reqChan)

	//
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			log.Println("accept tcp fail", err.Error())
			continue
		}
		go handleTcpConnection(conn, reqChan)
	}
}

func connectProxyServer(reqChan <-chan *Request) {
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
	log.Println("net.DialUDP", addr)
	defer conn.Close()

	sendChan := make(chan []byte, 1000)
	go sendHandler(conn, sendChan)

	recvChan := make(chan []byte, 1000)
	go recvHandler(conn, recvChan)

	reqMap := make(map[int64]*Request)
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
				req.conn.Write(protoWraperMessage.Message)
			}
		}
	}
}

func sendHandler(conn *net.UDPConn, sendChan <-chan []byte) {
	for data := range sendChan {
		//log.Println("处理转发请求")
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
			go handleMessage(conn, message, reqChan)
		}
	}
}

func handleMessage(conn *net.TCPConn, message *coder.Message, reqChan chan<- *Request) {

	protoMessage := bean.Factory((bean.MessageType)(message.Type))

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
	transformMessage(conn, message, reqChan)
}

func transformMessage(conn *net.TCPConn, message *coder.Message, reqChan chan<- *Request) {

	resChan := make(chan []byte, 1)
	//发送打包后的数据,数据中包含流水号
	rid := createRID()
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
