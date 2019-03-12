package tcpserver

import (
	"net"
	"bufio"
	"fmt"

	"github.com/tinywarrior/logging_server/logger"
)

func LaunchTCPServer() {

	li, err := net.Listen("tcp", ":1903")
	if err != nil {
		fmt.Println(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println(err)
		}

		go handleTCPConnections(conn)
	}

}

func handleTCPConnections(conn net.Conn) {

	scanner := bufio.NewScanner(conn)

	var log *logger.Log
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println(message)
		log = logger.JsonUnmarshall([]byte(message))
		logger.WriteToFile(log)
	}

	
	
	defer conn.Close()

}