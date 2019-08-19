package apilogger

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

func Println(msg string) {
	logger.Println(msg)
}

func Printf(format string, v ...interface{}) {
	logger.Printf(format, v)
}

func Fatal(msg interface{}) {
	logger.Fatalln(msg)
}

func LogApiRequest(method string, jsonPayload interface{}, headers map[string]string) {
	logger.Printf(fmt.Sprintf("Sending request. Method: '%v'", method))
	if headers != nil {
		logger.Println("with headers:")
		for k, v := range headers {
			logger.Printf(fmt.Sprintf("  %v: %v", k, v))
		}
	}
	bytes, e := json.MarshalIndent(jsonPayload, "", "    ")
	if e != nil {
		log.Fatal(e)
	}
	logger.Printf("json payload:\n%v", string(bytes))
}

func LogApiResponse(httpResponse *http.Response, jsonPayload interface{}) {
	logger.Println("Received response:")
	logger.Println(fmt.Sprintf("http status: %s", httpResponse.Status))
	bytes, e := json.MarshalIndent(jsonPayload, "", "    ")
	if e != nil {
		log.Fatal(e)
	}
	logger.Println(fmt.Sprintf("response body:\n%v", string(bytes)))
}
