package httpserver

import (
	"net/http"
	"html/template"
	"io/ioutil"
	"time"
	// "fmt"
	"bufio"
	"os"

	"logger"
	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

var logMessage *logger.Log

func init() {


	tpl = template.Must(template.ParseGlob("vendor/httpserver/templates/*.gohtml"))

	logMessage = &logger.Log{
		Service: "HTTP Server",
		Time: time.Now(),
		RemoteServer: "",
		Message: "",
	}
}

func LaunchHTTPServer() {

	mux := httprouter.New()
	logMessage.GenerateLogMessage("Starting HTTP server...")
	mux.POST("/write", writeLogs)
	mux.GET("/logs", getLogs)
	logMessage.GenerateLogMessage("HTTP Server started")
	http.ListenAndServe(":8080", mux)
}

func getLogs(res http.ResponseWriter, req *http.Request, _ httprouter.Params){

	files, err := ioutil.ReadDir("./logs")
	if err != nil {
		logMessage.GenerateErrorMessage(err)
	}

	type Data struct {
		Filename string
		AllLogs []string
	}
	var mapSlice []Data
	logMessage.GenerateLogMessage("Searching for files")
	for _, file := range files {
	
		logs, err := os.Open("./logs/" + file.Name(), )
		if err != nil {
			logMessage.GenerateErrorMessage(err)
		}
		defer logs.Close()

		var allLogs []string
		scanner := bufio.NewScanner(logs)
		for scanner.Scan() {
			log := scanner.Text()
			allLogs = append(allLogs, log)
		}
		m := Data{Filename: file.Name(), AllLogs: allLogs}
		mapSlice = append(mapSlice, m)

	}

	err = tpl.ExecuteTemplate(res, "logs.gohtml", mapSlice)
	if err != nil {
		logMessage.GenerateErrorMessage(err)
	}
}

func writeLogs(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	
	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		logMessage.GenerateErrorMessage(err)
	}
	logMessage.GenerateLogMessage("Parsing message body")
	var log *logger.Log

	
	// log.RemoteServer = strings.Split(req.RemoteAddr, ":")[0]
	log = logger.JsonUnmarshall(body)

	logMessage.GenerateLogMessage("Writing to file")
	logger.WriteToFile(log)

	res.Header().Set("Content-Type", "application/json")
	res.Write(body)

}