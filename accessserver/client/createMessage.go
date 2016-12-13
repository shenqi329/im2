package main

import (
	"github.com/golang/protobuf/proto"
	grpcPb "im/grpc/pb"
	"im/logicserver/uuid"
	client "im/protocol/client"
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

	raddr, err := net.ResolveTCPAddr("tcp", "localhost:6000")
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

	go handleConnection(connect)

	for i := 0; i < 100; i++ {
		{
			pb := &grpcPb.CreateMessageRequest{
				Rid:       getRid(),
				SessionId: 32,
				Type:      1,
				Id:        uuid.Rand().Hex(),
				Content:   "a message from push",
			}
			protoBuf, err := proto.Marshal(pb)

			request := &client.RpcRequest{
				Rid:         getRid(),
				AppId:       "89897",
				Type:        client.RpcRequest_LogicServer,
				MessageType: grpcPb.MessageTypeCreateMessageRequest,
				ProtoBuf:    protoBuf,
			}
			buffer, err := coder.EncoderProtoMessage(client.MessageTypeRPCRequest, request)
			if err != nil {
				log.Println(err.Error())
			}
			connect.Write(buffer)
			//time.Sleep(40 * time.Millisecond)
		}
	}
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

	log.Println("recvMsg count = ", gRecvCount, "context:", protoMessage.String())

	if message.Type == client.MessageTypeRPCResponse {
		rpcResponse, ok := protoMessage.(*client.RpcResponse)

		if ok && rpcResponse.ProtoBuf != nil {
			protoMessage = grpcPb.Factory((grpcPb.MessageType)(rpcResponse.MessageType))
			if err := proto.Unmarshal(rpcResponse.ProtoBuf, protoMessage); err != nil {
				log.Println(err.Error())
				log.Println("消息格式错误")
				conn.Close()
				return
			}
			log.Println(protoMessage.String())
		}
	}
	//log.Println("recvMsg count = ", gRecvCount, "context:", proto.CompactTextString(protoMessage))
}
