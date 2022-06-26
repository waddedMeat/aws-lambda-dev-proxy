package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var lambdaEndpoint string
var serverPort int

func init() {
	protocol := flag.String(`protocol`, `http`, `lambda protocol`)
	host := flag.String(`host`, `localhost`, `lambda host`)
	port := flag.Int(`port`, 8080, `lambda port`)

	flag.IntVar(&serverPort, `server-port`, 9000, `server port`)

	flag.Parse()

	lambdaEndpoint = fmt.Sprintf(
		`%s://%s:%d/2015-03-31/functions/function/invocations`,
		*protocol,
		*host,
		*port,
	)
}

func main() {
	// register handler
	http.HandleFunc(`/`, proxyHandler)

	log.Printf(`starting proxy for lambda: %s`, lambdaEndpoint)
	err := http.ListenAndServe(fmt.Sprintf(`:%d`, serverPort), nil)
	if err != nil {
		log.Fatalln(`Failed to start server:`, err)
	}
}

type Event struct {
	Body string `json:"body"`
}

type Response struct {
	Headers    map[string]string `json:"headers"`
	StatusCode int               `json:"statusCode"`
	Body       string            `json:"body"`
}

func proxyHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set(`Content-Type`, `application/json`)

	body := make([]byte, request.ContentLength)
	// @todo add some error handling to this
	request.Body.Read(body)

	payload, _ := json.Marshal(Event{Body: string(body)})

	log.Printf(`Calling Lambda: %s`, lambdaEndpoint)

	resp, _ := http.Post(
		lambdaEndpoint,
		request.Header.Get(`Content-Type`),
		bytes.NewBuffer(payload),
	)

	lambdaResponse := &Response{}
	err := json.NewDecoder(resp.Body).Decode(lambdaResponse)
	if err != nil {
		return
	}

	if lambdaResponse.Headers != nil {
		for header, value := range lambdaResponse.Headers {
			w.Header().Set(header, value)
		}
	}

	if lambdaResponse.StatusCode > 0 {
		w.WriteHeader(lambdaResponse.StatusCode)
	}

	if lambdaResponse.Body != `` {
		// @todo add some error handling to this
		w.Write([]byte(lambdaResponse.Body))
	}

}
