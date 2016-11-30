package main

import (
	"im/imserver/server"
	"log"
	"runtime"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	if runtime.GOOS == "windows" {
		localUdpAddr := "localhost:6001"
		s := server.NEWServer(localUdpAddr)
		s.Run()
	} else {
		localUdpAddr := "172.17.0.5:6001"
		s := server.NEWServer(localUdpAddr)
		s.Run()
	}
}
