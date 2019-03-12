package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"os"
	"time"
	"strings"
	"./customLogger"

	"github.com/julienschmidt/httprouter"
)

type Log struct {

	Service	string `json:"service"`
	RemoteServer string `json:"RemoteServer"`
	Time	time.Time `json:"time"`
	Message string `json:"message"`

}

func main() {

	logger := customLogger.GetInstance()
	logger.Println("Starting server...")
	mux := httprouter.New()
	mux.POST("/write", writeLogs)
	logger.Println("Server started")
	http.ListenAndServe(":8080", mux)
	
}



func writeLogs(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

	logger := customLogger.GetInstance()

	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		logger.Println(err)
	}
	logger.Println("Parsing message body")
	var log Log

	log.RemoteServer = strings.Split(req.RemoteAddr, ":")[0]
	err = json.Unmarshal(body, &log)
	if err != nil {
		logger.Println(err)
	}

	logger.Println("Writing to file")
	writeToFile(log)

	res.Header().Set("Content-Type", "application/json")
	res.Write(body)

}

func writeToFile(log Log) {

	logger := customLogger.GetInstance()

	fname := log.Service
	file, err := os.OpenFile("./logs/" + fname + ".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Println("Error opening file for write.")
	}
	defer file.Close()

	message, err := json.Marshal(log)
	if err != nil {
		logger.Println(err)
	}
	if _, err := file.WriteString(string(message[:]) + "\n"); err != nil{
		logger.Println("Error writing to file")
	}

}
