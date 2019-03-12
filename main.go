package main

import (
	"github.com/tinywarrior/logging_server/httpserver"
	"github.com/tinywarrior/logging_server/tcpserver"
)

func main() {

	go tcpserver.LaunchTCPServer()	
	httpserver.LaunchHTTPServer()
}




