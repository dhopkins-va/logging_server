package logger

import (
	"time"
	"os"
	"fmt"
	"encoding/json"
)

type Log struct {

	Service	string `json:"service"`
	RemoteServer string `json:"RemoteServer"`
	Time	time.Time `json:"time"`
	Message string `json:"message"`

}

func JsonUnmarshall(input []byte) (log *Log) {

	var unMarshalledLog *Log
	err := json.Unmarshal(input, &unMarshalledLog)
	if err != nil {
		fmt.Println("Error unmarshalling json")
		fmt.Println(err)
	}

	return unMarshalledLog
	
}

func (log *Log) GenerateErrorMessage(err error) {

	log.Time = time.Now()
	log.Message = err.Error()

	WriteToFile(log)
}

func (log *Log) GenerateLogMessage(message string) {

	log.Time = time.Now()
	log.Message = message

	WriteToFile(log)
}

func WriteToFile(log *Log) {

	fname := log.Service
	file, err := os.OpenFile("./logs/" + fname + ".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file for write.")
	}
	defer file.Close()

	message, err := json.Marshal(&log)
	if err != nil {
		fmt.Println(err)
	}
	if _, err := file.WriteString(string(message[:]) + "\n"); err != nil{
		fmt.Println("Error writing to file")
	}

}