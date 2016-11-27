package main

import (
	"im/accessserver/server"
	"log"
	"net"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	clientMulitiConnect()
}

func clientMulitiConnect() {

	var count int = 100000
	for i := 0; i < count; i++ {
		//raddr, err := net.ResolveTCPAddr("tcp", "172.17.0.2:6000")
		raddr, err := net.ResolveTCPAddr("tcp", server.ServerAddr())
		if err != nil {
			log.Println("net.ResolveTCPAddr fail.", err)
			//os.Exit(1)
			break
		}
		conn, err := net.DialTCP("tcp", nil, raddr)
		if err != nil {
			log.Println(err.Error())
			break
		}

		if i < count-1 {
			go mulitihandleConnection(conn)
		} else {
			mulitihandleConnection(conn)
		}
	}
	for true {
		time.Sleep(60 * time.Minute)
	}
}

func mulitihandleConnection(conn *net.TCPConn) {

	conn.SetKeepAlive(true)
	conn.SetKeepAlivePeriod(10 * time.Second)

	b := make([]byte, 1)

	conn.Read(b)

	defer conn.Close()
}
