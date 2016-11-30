package main

import (
	"im/accessserver/server"
	"log"
	"runtime"
)

func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	if runtime.GOOS == "windows" {
		localTcpAddr := "localhost:6000"
		proxyUdpAddr := "localhost:6001"
		s := server.NEWServer(localTcpAddr, proxyUdpAddr)
		s.Run()
	} else {
		localTcpAddr := "172.17.0.4:6000"
		proxyUdpAddr := "172.17.0.5:6001"
		s := server.NEWServer(localTcpAddr, proxyUdpAddr)
		s.Run()
	}
}
