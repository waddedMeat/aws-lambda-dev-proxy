package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func proxyHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if request.Method == `OPTIONS` {
		return
	}

	// read the request body into a []byte
	body := make([]byte, request.ContentLength)
	if read, err := request.Body.Read(body); err != nil && read < len(body) {
		// @todo respond w/ error
		//       log error
		return
	}

	// this is the event that is going to be passed to the lambda
	lambdaEvent := NewEvent(request.Method, string(body), request.Header)

	// marshal the json into a []byte for the request payload
	payload, _ := json.Marshal(lambdaEvent)

	log.Printf(`Calling Lambda: %s`, lambdaEndpoint)

	// call the lambda function endpoint
	resp, _ := http.Post(
		lambdaEndpoint,
		`application/json`,
		bytes.NewBuffer(payload), // the Post requires a *Buffer
	)

	// struct reference that represents the lambda response
	lambdaResponse := &Response{}
	// decode the json into the lambdaResponse
	if err := json.NewDecoder(resp.Body).Decode(lambdaResponse); err != nil {
		return
	}

	// set headers if they exist
	if lambdaResponse.Headers != nil {
		for header, value := range lambdaResponse.Headers {
			w.Header().Set(header, value)
		}
	}

	// set status code if it is set; int has a default value of 0
	if lambdaResponse.StatusCode > 0 {
		w.WriteHeader(lambdaResponse.StatusCode)
	}

	// set the response body if it exists; string has a default value of ''
	if lambdaResponse.Body != `` {
		if _, err := w.Write([]byte(lambdaResponse.Body)); err != nil {
			return
		}
		w.Header().Set(`Content-Type`, `application/json`)
	}

}
