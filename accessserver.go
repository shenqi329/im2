package main

import (
	"im/accessserver"
	"log"
	"runtime"
)

func main() {

	log.SetFlags(log.Lshortfile | log.LstdFlags)

	if runtime.GOOS == "windows" {
		localTcpAddr := "localhost:6000"
		proxyUdpAddr := "localhost:6001"
		s := accessserver.NEWServer(localTcpAddr, proxyUdpAddr)
		s.Run()
	} else {
		localTcpAddr := "localhost:6000"
		proxyUdpAddr := "localhost:6001"
		s := accessserver.NEWServer(localTcpAddr, proxyUdpAddr)
		s.Run()
	}
}
