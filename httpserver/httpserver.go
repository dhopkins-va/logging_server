package httpserver

import (
	"net/http"
	"io/ioutil"
	"fmt"

	"github.com/tinywarrior/logging_server/logger"
	"github.com/julienschmidt/httprouter"
)

func LaunchHTTPServer() {

	mux := httprouter.New()
	mux.POST("/write", writeLogs)
	http.ListenAndServe(":8080", mux)
}

func writeLogs(res http.ResponseWriter, req *http.Request, params httprouter.Params) {

	// logger := GetInstance()

	
	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		// logger.Println(err)
	}
	// logger.Println("Parsing message body")
	var log *logger.Log

	
	// log.RemoteServer = strings.Split(req.RemoteAddr, ":")[0]
	log = logger.JsonUnmarshall(body)

	fmt.Println("Writing to file")
	logger.WriteToFile(log)

	res.Header().Set("Content-Type", "application/json")
	res.Write(body)

}