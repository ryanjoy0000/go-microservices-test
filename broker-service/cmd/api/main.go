// BROKER Service

package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":80"

type BrokerAppConfig struct{}

func main() {

	// create app config
	brokerAppConf := BrokerAppConfig{}

	// define http server
	srv := &http.Server{
		Addr:    port,
		Handler: brokerAppConf.routes(),
	}

	// start server
	fmt.Println("Starting Broker service on address ", port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
