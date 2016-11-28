package main

import (
	"im/proxyserver/server"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	localUdpAddr := "localhost:6001"
	s := server.NEWServer(localUdpAddr)
	s.Run()
}
