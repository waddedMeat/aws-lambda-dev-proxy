package main

type Event struct {
	Body string `json:"body"`
}

type Response struct {
	Headers    map[string]string `json:"headers"`
	StatusCode int               `json:"statusCode"`
	Body       string            `json:"body"`
}
