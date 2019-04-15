// https://gist.github.com/reagent/043da4661d2984e9ecb1ccb5343bf438
package main

import (
	"log"
	"net/http"

	handler "./handler"

	"github.com/abhishekkr/gol/golenv"
)

func main() {
	listenAt := golenv.OverrideIfEnv("LISTEN_AT", ":9000")
	log.Printf("listening at: %s", listenAt)
	err := http.ListenAndServe(listenAt, handler.AppHandler())

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
