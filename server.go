package main

import (
	"fmt"
	"log"
	"bufio"
	"net"
	"strings"
	"os"
)

func main() {

	// Create server
	li, err := net.Listen("tcp", ":1902")
	if err != nil {
		panic(err)
	}

	defer li.Close()


	// Create the logs folder if it doesn't exist

	if _, err := os.Stat("./logs"); os.IsNotExist(err) {
		os.Mkdir("./logs", os.ModePerm)
		fmt.Println("Logs folder created")
	}
	// Handle incoming connections
	for {

		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		message := scanner.Text()
		project := strings.Split(message, " ")[0]
		writeLogs(project, message)
	}

	defer conn.Close()
}

func writeLogs(project string, message string) {

	_, err := os.Stat("./logs/" + project + "/" + project + ".log")
	if err != nil {
		if os.IsNotExist(err) {
			createLogFile(project)
		} else {
			fmt.Println("Error checking if logs already exist")
		}
	}

	file, err := os.OpenFile("./logs/" + project + "/" + project + ".log", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("Error opening file for write.")
	}
	defer file.Close()

	if _, err := file.WriteString(message + "\n"); err != nil{
		fmt.Println("Error writing to file")
	}

}

func createLogFile(project string) {

	os.Mkdir("./logs/" + project, os.ModePerm)
	file, err := os.Create("./logs/" + project + "/" + project + ".log")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("New project created")
	file.Close()
}