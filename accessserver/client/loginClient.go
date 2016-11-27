package main

import (
	"github.com/golang/protobuf/proto"
	"im/accessserver/server"
	"im/protocal/bean"
	"im/protocal/coder"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	connectToPort()
}

var rid int64 = 0

func getRid() int64 {
	rid++
	return rid
}

func connectToPort() {

	//raddr, err := net.ResolveTCPAddr("tcp", "172.17.0.2:6000")
	raddr, err := net.ResolveTCPAddr("tcp", server.ServerAddr())

	if err != nil {
		log.Println("net.ResolveTCPAddr fail.", err)
		os.Exit(1)
	}
	connect, err := net.DialTCP("tcp", nil, raddr)

	if err != nil {
		log.Println("net.ListenTCP fail.", err.Error())
		os.Exit(1)
	}

	connect.SetKeepAlive(true)
	connect.SetKeepAlivePeriod(10 * time.Second)
	go handleConnection(connect)

	for i := 0; i < 100000; i++ {
		{
			registerRequest := &bean.DeviceRegisteRequest{
				Rid:      getRid(),
				SsoToken: "123456dc22425556dd01605d438f4d0c",
				AppId:    "89897",
				DeviceId: "024b36dc22425556bc01605d438f4d0c",
				Platform: "windows",
			}
			buffer, err := coder.EncoderProtoMessage(bean.MessageTypeDeviceRegisteRequest, registerRequest)
			if err != nil {
				log.Println(err.Error())
			}
			connect.Write(buffer)
		}
		{
			loginRequest := &bean.DeviceLoginRequest{
				Rid:      getRid(),
				Token:    "123456dc22425556bc01605d438f4d0c",
				AppId:    "89897",
				DeviceId: "024b36dc22425556bc01605d438f4d0c",
				Platform: "windows",
			}
			buffer, err := coder.EncoderProtoMessage(bean.MessageTypeDeviceLoginRequest, loginRequest)
			if err != nil {
				log.Println(err.Error())
			}
			connect.Write(buffer)
		}
		log.Println(i)
		time.Sleep(1 * time.Millisecond)
	}

	time.Sleep(60 * time.Minute)
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

}