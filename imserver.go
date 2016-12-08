package main

import (
	"im/imserver"
	controller "im/imserver/controller"
	protocolClient "im/protocol/client"
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
		localUdpAddr = "localhost:6001"
	}

	s := imserver.NEWServer(localUdpAddr)

	s.Handle(protocolClient.MessageTypeDeviceRegisteRequest, controller.HandleRegiste)
	s.Handle(protocolClient.MessageTypeDeviceLoginRequest, controller.HandleLogin)
	s.Handle(protocolClient.MessageTypeRPCRequest, controller.HandleRpc)

	s.Run()
}
