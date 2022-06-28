package main

import "net/http"

type event struct {
	Method  string            `json:"method"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

func NewEvent(method, body string, headers http.Header) *event {
	event := &event{
		Method: method,
		Body:   body,
	}
	event.Headers = make(map[string]string)
	for header, values := range headers {
		for _, value := range values {
			event.Headers[header] = value
		}
	}
	return event
}

type Response struct {
	Headers    map[string]string `json:"headers"`
	StatusCode int               `json:"statusCode"`
	Body       string            `json:"body"`
}
