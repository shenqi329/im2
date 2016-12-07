package main

import (
	"github.com/golang/protobuf/proto"
	grpcSession "im/grpc/session"
	imserverBean "im/imserver/bean"
	"im/protocol/client"
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
			pb := &grpcSession.CreateSessionRequest{
				Rid:          getRid(),
				AppId:        "89897",
				CreateUserId: "1",
				Count:        1,
			}
			protoBuf, err := proto.Marshal(pb)

			request := &client.RpcRequest{
				Rid:         getRid(),
				AppId:       "89897",
				MessageType: 1,
				ProtoBuf:    protoBuf,
			}
			buffer, err := coder.EncoderProtoMessage(client.MessageTypeRPCRequest, request)
			if err != nil {
				log.Println(err.Error())
			}

			connect.Write(buffer)
		}
	}

	// for i := 0; i < 33; i++ {
	// 	{
	// 		request := &bean.CreateSessionRequest{
	// 			Rid:          getRid(),
	// 			AppId:        "89897",
	// 			CreateUserId: "1",
	// 			Count:        1,
	// 		}
	// 		buffer, err := coder.EncoderProtoMessage(bean.MessageTypeCreateSessionRequest, request)
	// 		if err != nil {
	// 			log.Println(err.Error())
	// 		}

	// 		// wraper := &bean.WraperMessage{
	// 		// 	ConnId:  1,
	// 		// 	Message: buffer,
	// 		// }

	// 		// buffer, err = coder.EncoderProtoMessage(bean.MessageTypeWraper, wraper)
	// 		// if err != nil {
	// 		// 	log.Println(err.Error())
	// 		// }

	// 		connect.Write(buffer)
	// 	}
	// }
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

	protoMessage := client.Factory((client.MessageType)(message.Type))

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
