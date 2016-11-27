package server

import (
	"im/protocal/bean"
	"im/protocal/coder"
	"log"
	"net"
	"os"
)

func ServerUDPAddr() string {
	return "localhost:6001"
}

func ListenOnPort() {
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

		decoder.Decode(buffer[0:count])
		if err != nil {
			log.Println(err.Error())
			break
		}
		//处理
		for _, message := range messages {
			go handleMessage(listen, udpAddr, message)
		}
	}
}

func handleMessage(listen *net.UDPConn, addr *net.UDPAddr, message *coder.Message) {
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
	case MessageTypeDeviceRegisteRequest:
		{
			response = &bean.DeviceLoginResponse{}
		}
	case MessageTypeDeviceLoginRequest:
		{

		}
	}
}
