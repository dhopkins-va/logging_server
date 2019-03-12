package main

import (
	"os"
	"log"
	"sync"
	
	"github.com/tinywarrior/logging_server/httpserver"
	"github.com/tinywarrior/logging_server/tcpserver"
)

type logger struct {
    filename string
    *log.Logger
}

var customLogger *logger
var once sync.Once

// start logging
func GetInstance() *logger {

	once.Do(func() {
		customLogger = createLogger("server.log")
	})

	return customLogger
}

func createLogger(fname string) *logger {

	file, _ := os.OpenFile("./logs/" + fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	return &logger{
		filename: fname,
		Logger: log.New(file, "Server: ", log.Lshortfile|log.LstdFlags),
	}
}



func main() {

	go tcpserver.LaunchTCPServer()	
	httpserver.LaunchHTTPServer()
}




