package tcpserver

import (
	"net"
	"bufio"
	"time"

	"logger"
)

var logMessage *logger.Log

func init() {

	logMessage = &logger.Log{
		Service: "TCP Server",
		Time: time.Now(),
		RemoteServer: "",
		Message: "",
	}
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
			logMessage.GenerateErrorMessage(err)
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
		log = logger.JsonUnmarshall([]byte(message))
		logger.WriteToFile(log)
		logMessage.GenerateLogMessage("Logs written successfully")
	}

	
	
	defer conn.Close()

}