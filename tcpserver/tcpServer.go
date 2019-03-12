package tcpserver

import (
	"net"
	"bufio"
	"fmt"

	"github.com/tinywarrior/logging_server/logger"
)

var logMessage *logger.Log

func init() {

	logMessage.Service = "HTTPServer"
	logMessage.RemoteServer = "Local Server"
}

func LaunchTCPServer() {

	logMessage.GenerateLogMessage("Starting TCP server...")
	li, err := net.Listen("tcp", ":1903")
	if err != nil {
		logMessage.GenerateErrorMessage(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println(err)
		}

		go handleTCPConnections(conn)
		logMessage.GenerateLogMessage("TCP Server started...")
	}

}

func handleTCPConnections(conn net.Conn) {

	logMessage.GenerateLogMessage("Incoming log message...")
	scanner := bufio.NewScanner(conn)

	var log *logger.Log
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println(message)
		log = logger.JsonUnmarshall([]byte(message))
		logger.WriteToFile(log)
		logMessage.GenerateLogMessage("Logs written successfully")
	}

	
	
	defer conn.Close()

}