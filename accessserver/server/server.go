package server

import (
	proto "github.com/golang/protobuf/proto"
	pb "im/accessserver/bean"
	"im/accessserver/coder"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

var linkCount int = 0

func ListenOnPort() {

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:6000")

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
			//log.Println(err.Error())
			break
		}
		messages, err := decoder.Decode(buf[0:count])
		if err != nil {
			//log.Println(err.Error())
			break
		}
		for _, v := range messages {
			person := &pb.Person{}
			if err := proto.Unmarshal(v.MessageBuf, person); err != nil {
				log.Println(err.Error())
				return
			}
		}
	}

}
