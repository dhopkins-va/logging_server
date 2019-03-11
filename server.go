package main

import (
	"fmt"
	"bufio"
	"net"
	"strings"
	"os"
	"./customLogger"
)

func init() {

	// Create log folder
	if _, err := os.Stat("./logs"); os.IsNotExist(err) {
		os.Mkdir("./logs", os.ModePerm)
		fmt.Println("Logs folder created")
	}

	logger := customLogger.GetInstance()
	logger.Println("Init complete")
}


func main() {

	logger := customLogger.GetInstance()

	// Create server
	li, err := net.Listen("tcp", ":1902")
	if err != nil {
		panic(err)
	}

	logger.Println("Server listening...")

	defer li.Close()
	// Handle incoming connections
	for {

		conn, err := li.Accept()
		if err != nil {
			logger.Fatalln(err)
		}

		logger.Println("Connection received")
		go handle(conn)
	}
}

func handle(conn net.Conn) {

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		message := scanner.Text()
		fname := strings.Split(message, " ")[0]
		fname+= ".txt"
		writeLogs(fname, message)
	}

	defer conn.Close()
}

func writeLogs(fname string, message string) {


	logger := customLogger.GetInstance()

	file, err := os.OpenFile("./logs/" + fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Println("Error opening file for write.")
	}
	defer file.Close()

	if _, err := file.WriteString(message + "\n"); err != nil{
		fmt.Println("Error writing to file")
	}

}