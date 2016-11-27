package main

import (
	"im/accessserver/client"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	//client.ConnectToPort()
	client.ClientMulitiConnect()
}
