package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	server "github.com/miro-puetz/http-server/src"
)

var addr string = ":8080"

func main() {

	logFilename := flag.String("logFilename", "", "log output to given file name or leave empty for standard output")

	flag.Parse()

	if *logFilename != "" {
		file, err := os.OpenFile(*logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(file)
	}

	srv := &http.Server{
		Addr:              addr,
		Handler:           http.HandlerFunc(server.DefaultHandler),
		IdleTimeout:       5 * time.Minute,
		ReadHeaderTimeout: time.Minute,
	}

	log.Println("Server started")
	log.Fatal(srv.ListenAndServe())
}
