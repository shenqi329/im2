package main

import (
	"im/accessserver/server"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	server.Start()
}
