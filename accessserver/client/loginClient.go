package main

import (
	"github.com/golang/protobuf/proto"
	"im/protocal/bean"
	"im/protocal/coder"
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

	raddr, err := net.ResolveTCPAddr("tcp", "172.17.0.4:6000")
	if runtime.GOOS == "windows" {
		raddr, err = net.ResolveTCPAddr("tcp", "localhost:6000")
	}

	if err != nil {
		log.Println("net.ResolveTCPAddr fail.", err)
		return
	}
	connect, err := net.DialTCP("tcp", nil, raddr)

	if err != nil {
		log.Println("net.ListenTCP fail.", err.Error())
		return
	}

	connect.SetKeepAlive(true)
	connect.SetKeepAlivePeriod(10 * time.Second)
	go handleConnection(connect)

	for i := 0; i < 1; i++ {
		{
			registerRequest := &bean.DeviceRegisteRequest{
				Rid:      getRid(),
				SsoToken: "7f3ce0ecde6a4ad984126bd76b34f7a6",
				AppId:    "89897",
				DeviceId: "024b36dc22425556bc01605d438f4d0c",
				Platform: "windows",
			}
			buffer, err := coder.EncoderProtoMessage(bean.MessageTypeDeviceRegisteRequest, registerRequest)
			if err != nil {
				log.Println(err.Error())
			}
			connect.Write(buffer)
			time.Sleep(1 * time.Millisecond)
		}
		// {
		// 	loginRequest := &bean.DeviceLoginRequest{
		// 		Rid:      getRid(),
		// 		Token:    "123456dc22425556bc01605d438f4d0c",
		// 		AppId:    "89897",
		// 		DeviceId: "024b36dc22425556bc01605d438f4d0c",
		// 		Platform: "windows",
		// 	}
		// 	buffer, err := coder.EncoderProtoMessage(bean.MessageTypeDeviceLoginRequest, loginRequest)
		// 	if err != nil {
		// 		log.Println(err.Error())
		// 	}
		// 	connect.Write(buffer)
		// }
	}
	// time.Sleep(60 * time.Minute)
}

func handleConnection(conn *net.TCPConn) {

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

func handleMessage(conn *net.TCPConn, message *coder.Message) {

	protoMessage := bean.Factory((bean.MessageType)(message.Type))

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
	log.Println("recvMsg count = ", gRecvCount, "context:", proto.CompactTextString(protoMessage))

}
