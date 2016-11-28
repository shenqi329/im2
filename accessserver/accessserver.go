package main

import (
	"im/accessserver/server"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	localTcpAddr := "localhost:6000"
	proxyUdpAddr := "localhost:6001"

	s := server.NEWServer(localTcpAddr, proxyUdpAddr)
	s.Run()
}
