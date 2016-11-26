package server

import (
	proto "github.com/golang/protobuf/proto"
	pb "im/accessserver/bean"
	"log"
	"net"
	"os"
	"time"
)

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

func handleConnection(conn *net.TCPConn) {

	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(10 * time.Second)

	buf := make([]byte, 1024)

	count, err := conn.Read(buf)

	if err != nil {
		log.Println(err.Error())
	}
	log.Println(count)
	person := &pb.Person{}

	log.Println(proto.MessageName(person))

	if err := proto.Unmarshal(buf[:count], person); err != nil {
		log.Println(err.Error())
	}
	log.Println(proto.MessageName(person))

	log.Println(person.String())
}
