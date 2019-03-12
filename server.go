package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
	"time"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type Log struct {

	Service	string `json:"service"`
	RemoteServer string `json:"RemoteServer"`
	Time	time.Time `json:"time"`
	Message string `json:"message"`

}

func main() {

	mux := httprouter.New()
	mux.POST("/write", writeLogs)
	http.ListenAndServe(":8080", mux)
}



func writeLogs(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		fmt.Println(err)
	}

	var log Log
	fmt.Printf("%T", req.RemoteAddr)
	log.RemoteServer = strings.Split(req.RemoteAddr, ":")[0]
	err = json.Unmarshal(body, &log)
	if err != nil {
		fmt.Println(err)
	}

	writeToFile(log)

	res.Header().Set("Content-Type", "application/json")
	res.Write(body)

}

func writeToFile(log Log) {

	fname := log.Service
	file, err := os.OpenFile("./logs/" + fname, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file for write.")
	}
	defer file.Close()

	message, err := json.Marshal(log)
	if err != nil {
		fmt.Println(err)
	}
	if _, err := file.WriteString(string(message[:]) + "\n"); err != nil{
		fmt.Println("Error writing to file")
	}

}