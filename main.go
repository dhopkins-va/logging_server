package main

import (
	"httpserver"
	"tcpserver"
)

func main() {

	go tcpserver.LaunchTCPServer()	
	httpserver.LaunchHTTPServer()
}




