package server

// import (
// 	"github.com/golang/protobuf/proto"
// 	bean "im/protocal/bean"
// 	coder "im/protocal/coder"
// 	"log"
// 	"net"
// 	"os"
// )

// func ListenUDPOnPort(laddr string) {

// 	addr, err := net.ResolveUDPAddr("udp", laddr)

// 	if err != nil {
// 		log.Println("net.ResolveUDPAddr fail.", err)
// 		os.Exit(1)
// 	}

// 	listen, err := net.ListenUDP("udp", addr)
// 	defer listen.Close()

// 	if err != nil {
// 		log.Println("net.ListenUDP fail.", err)
// 		os.Exit(1)
// 	}
// 	log.Println("net.ListenUDP", addr)

// 	buffer := make([]byte, 4096)
// 	decoder := coder.NEWDecoder()
// 	for true {
// 		decoder.Reset()

// 		count, udpAddr, err := listen.ReadFromUDP(buffer)
// 		if err != nil {
// 			log.Println("读取数据失败!", err.Error())
// 			continue
// 		}

// 		messages, err := decoder.Decode(buffer[0:count])
// 		if err != nil {
// 			log.Println(err.Error())
// 			continue
// 		}
// 		//处理
// 		for _, message := range messages {
// 			go handleUDPMessage(listen, udpAddr, message)
// 		}
// 	}
// }

// func handleUDPMessage(listen *net.UDPConn, addr *net.UDPAddr, message *coder.Message) {
// 	protoMessage := bean.Factory((bean.MessageType)(message.MessageType))

// 	if protoMessage == nil {
// 		log.Println("未识别的消息")
// 		return
// 	}

// 	if err := proto.Unmarshal(message.MessageBuf, protoMessage); err != nil {
// 		log.Println(err.Error())
// 		log.Println("消息格式错误")
// 		return
// 	}
// 	//log.Println("1")
// 	//log.Println(proto.CompactTextString(protoMessage))

// 	switch message.MessageType {
// 	case bean.MessageTypeDeviceRegisteResponse:
// 		{
// 			handleRegisterResponse(listen, addr, protoMessage.(*bean.DeviceRegisteResponse))
// 		}
// 	case bean.MessageTypeDeviceLoginRequest:
// 		{
// 			handleLoginResponse(listen, addr, protoMessage.(*bean.DeviceLoginResponse))
// 		}
// 	}
// }

// func handleRegisterResponse(listen *net.UDPConn, addr *net.UDPAddr, request *bean.DeviceRegisteResponse) {

// 	// response := &bean.DeviceRegisteResponse{
// 	// 	Rid:   request.Rid,
// 	// 	Code:  "00000001",
// 	// 	Desc:  "success",
// 	// 	Token: "a token from proxyserver",
// 	// }

// 	// b, err := coder.EncoderProtoMessage(bean.MessageTypeDeviceRegisteResponse, response)
// 	// if err != nil {
// 	// 	log.Println(err.Error())
// 	// 	return
// 	// }
// 	// listen.WriteTo(b, addr)
// }

// func handleLoginResponse(listen *net.UDPConn, addr *net.UDPAddr, request *bean.DeviceLoginResponse) {
// 	// response := &bean.DeviceLoginResponse{
// 	// 	Rid:  request.Rid,
// 	// 	Code: "00000001",
// 	// 	Desc: "success",
// 	// }

// 	// b, err := coder.EncoderProtoMessage(bean.MessageTypeDeviceLoginResponse, response)
// 	// if err != nil {
// 	// 	log.Println(err.Error())
// 	// 	return
// 	// }
// 	// listen.WriteTo(b, addr)
// }
