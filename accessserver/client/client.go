package client

import (
	"log"
	"net"
	"os"
	"time"
)

func ConnectToPort() {

	addr, err := net.ResolveTCPAddr("tcp", ":8080")

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

}
