package customLogger

import (
	"log"
	"os"
	"sync"
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