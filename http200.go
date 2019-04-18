// https://gist.github.com/reagent/043da4661d2984e9ecb1ccb5343bf438
package main

import (
	"log"
	"net/http"

	handler "github.com/abhishekkr/http200/handler"
	httplog "github.com/abhishekkr/http200/httplog"

	"github.com/abhishekkr/gol/golenv"
)

func main() {
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
