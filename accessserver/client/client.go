package client

import (
	proto "github.com/golang/protobuf/proto"
	pb "im/accessserver/bean"
	"log"
	"net"
	"os"
	"time"
)

func ConnectToPort() {

	raddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:6000")

	if err != nil {
		log.Println("net.ResolveTCPAddr fail.", err)
		os.Exit(1)
	}
	connect, err := net.DialTCP("tcp", nil, raddr)

	if err != nil {
		log.Println("net.ListenTCP fail.", err.Error())
		os.Exit(1)
	}

	person := &pb.Person{
		Id:    -24,
		Name:  "wujingchao",
		Email: "wujingchao92@gmail.com",
	}

	b, err := proto.Marshal(person)
	if err != nil {
		log.Println("proto.Marshal fail.", err.Error())
		os.Exit(1)
	}

	count, err := connect.Write(b)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	log.Println(count)

	// listen, err := net.ListenTCP("tcp", addr)
	// defer listen.Close()

	// if err != nil {
	// 	log.Println("net.ListenTCP fail.", err)
	// 	os.Exit(1)
	// }

	// log.Println("net.ListenTCP", addr)

	// for {
	// 	conn, err := listen.AcceptTCP()
	// 	if err != nil {
	// 		log.Println("accept tcp fail", err.Error())
	// 		continue
	// 	}
	// 	go handleConnection(conn)
	// }
}

func handleConnection(conn *net.TCPConn) {

	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(10 * time.Second)

}
