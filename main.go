package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	server "github.com/miro-puetz/http-server/src/server"
)

var addr string = ":8080"

func main() {

	logFilename := flag.String("logFilename", "", "log output to given filename or leave empty for standard output")

	flag.Parse()

	if *logFilename != "" {
		file, err := os.OpenFile(*logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(file)
	}

	mux := server.NewMux()

	cert, err := tls.LoadX509KeyPair("cert/localhost.crt", "cert/localhost.key")
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		IdleTimeout:       5 * time.Minute,
		ReadHeaderTimeout: time.Minute,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	log.Println("Server started")
	log.Fatal(srv.ListenAndServeTLS("", ""))
}
