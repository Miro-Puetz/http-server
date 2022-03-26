package main

import (
	"log"
	"net/http"
	"time"

	server "github.com/miro-puetz/http-server/src"
)

var addr string = ":8080"

func main() {

	srv := &http.Server{
		Addr:              addr,
		Handler:           http.HandlerFunc(server.DefaultHandler),
		IdleTimeout:       5 * time.Minute,
		ReadHeaderTimeout: time.Minute,
	}

	log.Fatal(srv.ListenAndServe())
}
