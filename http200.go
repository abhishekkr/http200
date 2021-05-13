// https://gist.github.com/reagent/043da4661d2984e9ecb1ccb5343bf438
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	handler "github.com/abhishekkr/http200/handler"
	httplog "github.com/abhishekkr/http200/httplog"

	"github.com/abhishekkr/gol/golenv"
)

var (
	needHelp = flag.Bool("help", false, "need help with usage")
)

func main() {
	flag.Parse()
	if *needHelp {
		displayHelp()
		return
	}
	listenAt := golenv.OverrideIfEnv("HTTP200_LISTEN_AT", ":9000")
	log.Printf("listening at: %s", listenAt)
	err := http.ListenAndServe(listenAt, httplogHandler(handler.AppHandler()))

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}

func httplogHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httplog.LogRequest(r)
		next.ServeHTTP(w, r)
	})
}

func displayHelp() {
	fmt.Println(`
http200

latest release: https://github.com/abhishekkr/http200/releases/latest

It's your friendly http server to use as placeholder for integration points of your service under development.

### it provides:

* listens default at port ':9000', allows to change it using environment variable like 'HTTP200_LISTEN_AT=:8080'

* enable printing request body using environment variable 'HTTP200_BODY=true'

* shows this wiki at '/wiki'

* a simple placeholder http server providing '/200','/400','/404','/500' for respective HTTP response codes

* returns '404' response code for any non-default or non-customized route

* un-handled route's response status code could be customized via env 'HTTP200_DEFAULT_ROUTE' with values 'Route200', 'Route400', 'Route404', 'Route500'

* add custom route and request method with response code and body if required

---
	`)
}
