package main

import (
	"im/accessserver/coder"
	"im/accessserver/server"
	bean "im/protocal/bean"
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
	time.Sleep(60 * time.Minute)
}

func handleConnection(conn *net.TCPConn) {

	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(10 * time.Second)

}
