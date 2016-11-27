package client

import (
	pb "im/accessserver/bean"
	"im/accessserver/coder"
	"log"
	"net"
	"os"
	"time"
)

func ConnectToPort() {

	//raddr, err := net.ResolveTCPAddr("tcp", "172.17.0.2:6000")
	raddr, err := net.ResolveTCPAddr("tcp", coder.Addr())

	if err != nil {
		log.Println("net.ResolveTCPAddr fail.", err)
		os.Exit(1)
	}
	connect, err := net.DialTCP("tcp", nil, raddr)

	if err != nil {
		log.Println("net.ListenTCP fail.", err.Error())
		os.Exit(1)
	}

	persons := []*pb.Person{
		&pb.Person{
			Id:    -24,
			Name:  "wujingchao",
			Email: "1@gmail.com",
		},
		&pb.Person{
			Id:    -24,
			Name:  "wujingchao",
			Email: "wujingchao92wujingchao92wujingchao92@gmail.com",
		},
		&pb.Person{
			Id:    -24,
			Name:  "wujingchao",
			Email: "wujingchao92wujingchao92wujingchao92wujingchao92wujingchao92wujingchao92wujingchao92wujingchao92wujingchao92wujingchao92wujingchao92wujingchao92wujingchao92wujingchao92wujingchao92wujingchao92wujingchao92wujingchao92wujingchao92wujingchao92wujingchao92@gmail.com",
		},
		&pb.Person{
			Id:    -24,
			Name:  "wujingchao",
			Email: "@gmail.com",
		}, &pb.Person{
			Id:    -24,
			Name:  "wujingchao",
			Email: "@gmail.com",
		}, &pb.Person{
			Id:    -24,
			Name:  "wujingchao",
			Email: "1@gmail.com",
		}, &pb.Person{
			Id:    -24,
			Name:  "wujingchao",
			Email: "2@gmail.com",
		}, &pb.Person{
			Id:    -24,
			Name:  "wujingchao",
			Email: "3@gmail.com",
		}, &pb.Person{
			Id:    -24,
			Name:  "wujingchao",
			Email: "4@gmail.com",
		}, &pb.Person{
			Id:    -24,
			Name:  "wujingchao",
			Email: "5@gmail.com",
		}, &pb.Person{
			Id:    -24,
			Name:  "wujingchao",
			Email: "6@gmail.com",
		}, &pb.Person{
			Id:    -24,
			Name:  "wujingchao",
			Email: "7@gmail.com",
		}, &pb.Person{
			Id:    -24,
			Name:  "wujingchao",
			Email: "8@gmail.com",
		}, &pb.Person{
			Id:    -24,
			Name:  "wujingchao",
			Email: "9@gmail.com",
		}, &pb.Person{
			Id:    -24,
			Name:  "wujingchao",
			Email: "10@gmail.com",
		},
	}

	for i := 0; i < 1; i++ {
		for _, person := range persons {

			b, err := coder.EncoderProtoMessage(1, person)
			if err != nil {
				log.Println(err.Error())
			}
			for _, byt := range b {
				_, err := connect.Write([]byte{byt})
				if err != nil {
					log.Println(err.Error())
					os.Exit(1)
				}
			}
		}
	}
	buf := make([]byte, 1024)
	connect.Read(buf)
}

func handleConnection(conn *net.TCPConn) {

	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(10 * time.Second)

}
