package main

import (
	"github.com/golang/protobuf/proto"
	imserverBean "im/imserver/bean"
	protocolClient "im/protocol/client"
	"im/protocol/coder"
	"log"
	"net"
	"runtime"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	for i := 0; i < 1; i++ {
		go connectToPort()
		//time.Sleep(20 * time.Millisecond)
	}
	time.Sleep(60 * time.Minute)
}

var gRid uint64 = 0
var gRecvCount uint32 = 0

func getRid() uint64 {
	gRid++
	//log.Println(gRid)
	return gRid
}

func connectToPort() {

	raddr, err := net.ResolveUDPAddr("udp", "localhost:6001")
	if runtime.GOOS == "windows" {
		raddr, err = net.ResolveUDPAddr("udp", "localhost:6001")
	}

	if err != nil {
		log.Println("net.ResolveTCPAddr fail.", err)
		return
	}
	connect, err := net.DialUDP("udp", nil, raddr)

	if err != nil {
		log.Println("net.ListenTCP fail.", err.Error())
		return
	}

	go handleConnection(connect)

	for i := 0; i < 1; i++ {
		{
			request := &protocolClient.CreateMessageRequest{
				Rid:       getRid(),
				SessionId: 32,
				Type:      1,
				Content:   "a message from push",
			}
			buffer, err := coder.EncoderProtoMessage(protocolClient.MessageTypeCreateMessageRequest, request)
			if err != nil {
				log.Println(err.Error())
			}

			connect.Write(buffer)
			time.Sleep(1 * time.Millisecond)
		}
	}
}

func handleConnection(conn *net.UDPConn) {

	decoder := coder.NEWDecoder()
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

func handleMessage(conn *net.UDPConn, message *coder.Message) {

	protoMessage := protocolClient.Factory((protocolClient.MessageType)(message.Type))

	if protoMessage == nil {
		log.Println("未识别的消息")
		conn.Close()
		return
	}

	if err := proto.Unmarshal(message.Body, protoMessage); err != nil {
		log.Println(err.Error())
		log.Println("消息格式错误")
		conn.Close()
		return
	}
	gRecvCount++
	//log.Println("recvMsg count = ", gRecvCount, "context:", proto.CompactTextString(protoMessage))
	log.Println("recvMsg count = ", gRecvCount, "context:", imserverBean.StructToJsonString(protoMessage))
}
