package server

import (
	"github.com/golang/protobuf/proto"
	bean "im/protocal/bean"
	coder "im/protocal/coder"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

var linkCount int = 0

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

	buffer := make([]byte, 2048)
	for true {

		count, udpAddr, err := listen.ReadFromUDP(buffer)
		if err != nil {
			log.Println(err.Error())
		}
		count = count
		udpAddr = udpAddr
	}
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

	log.Println("net.ListenTCP", addr)

	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			log.Println("accept tcp fail", err.Error())
			continue
		}
		go handleConnection(conn)
	}
}

var mutex sync.Mutex

func handleConnection(conn *net.TCPConn) {

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
			handleMessage(conn, message)
		}
	}
}

func handleMessage(conn *net.TCPConn, message *coder.Message) {

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
}
