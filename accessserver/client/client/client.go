package client

import (
	"github.com/golang/protobuf/proto"
	grpcPb "im/logicserver/grpc/pb"
	"im/protocol/coder"
	"log"
	"net"
	"runtime"
	"sync/atomic"
	"time"
)

type Client struct {
	rid       uint64
	recvCount uint32
	Conn      *net.TCPConn

	loginState uint32

	afterLogin func(c *Client)
}

func (c *Client) GetRid() uint64 {
	atomic.AddUint64(&c.rid, 1)
	return c.rid
}

func (c *Client) SetAfterLogin(afterLogin func(c *Client)) {
	c.afterLogin = afterLogin
}

func (c *Client) LoginToAccessServer() {

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

	connect.SetKeepAlive(true)
	connect.SetKeepAlivePeriod(10 * time.Second)

	c.Conn = connect
	go c.handleConnection(connect)

	for i := 0; i < 1; i++ {
		if runtime.GOOS == "windows" {
			loginRequest := &grpcPb.DeviceLoginRequest{
				Rid:      c.GetRid(),
				Token:    "1",
				UserId:   "1",
				AppId:    "89897",
				DeviceId: "024b36dc22425556bc01605d438f4d0c",
				Platform: "windows",
			}
			buffer, err := coder.EncoderProtoMessage(grpcPb.MessageTypeDeviceLoginRequest, loginRequest)
			if err != nil {
				log.Println(err.Error())
			}
			connect.Write(buffer)
		} else {
			loginRequest := &grpcPb.DeviceLoginRequest{
				Rid:      c.GetRid(),
				Token:    "1",
				UserId:   "1",
				AppId:    "89897",
				DeviceId: "024b36dc22425556bc01605d438f4d0c",
				Platform: "windows",
			}

			buffer, err := coder.EncoderProtoMessage(grpcPb.MessageTypeDeviceLoginRequest, loginRequest)
			if err != nil {
				log.Println(err.Error())
			}
			connect.Write(buffer)
		}
	}
}

func (c *Client) handleConnection(conn *net.TCPConn) {

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
			go c.handleMessage(conn, message)
		}
	}
}

func (c *Client) handleMessage(conn *net.TCPConn, message *coder.Message) {

	protoMessage := grpcPb.Factory((grpcPb.MessageType)(message.Type))

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
	c.recvCount++
	log.Println("recvMsg count = ", c.recvCount, "context:", proto.CompactTextString(protoMessage))
	if (grpcPb.MessageType)(message.Type) == grpcPb.MessageTypeDeviceLoginResponse {
		c.loginState = 1
		go c.afterLogin(c)
	}
}
