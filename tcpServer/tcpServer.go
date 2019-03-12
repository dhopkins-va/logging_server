package tcpServer

import (
	"net"
	"bufio"
)

func launchTCPServer() {

	logger := GetInstance()

	logger.Println("Starting tcp server...")
	li, err := net.Listen("tcp", ":1903")
	if err != nil {
		logger.Println(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			logger.Println(err)
		}

		logger.Println("Listening for connections...")
		go handleTCPConnections(conn)
	}

}

func handleTCPConnections(conn net.Conn) {

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println(message)
	}

	defer conn.Close()

}