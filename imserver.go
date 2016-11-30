package main

import (
	"im/imserver"
	controller "im/imserver/controller"
	protocolBean "im/protocol/bean"
	"log"
	"runtime"
)

func DeviceLogin(c imserver.Context) error {
	return nil
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	var localUdpAddr string
	if runtime.GOOS == "windows" {
		localUdpAddr = "localhost:6001"
	} else {
		localUdpAddr = "172.17.0.5:6001"
	}

	s := imserver.NEWServer(localUdpAddr)

	s.Handle(protocolBean.MessageTypeDeviceRegisteRequest, controller.HandleRegiste)

	s.Run()
}