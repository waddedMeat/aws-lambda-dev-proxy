package main

import (
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
