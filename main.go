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

func main() {
	logFilename := flag.String("logFilename", "", "log output to given filename or leave empty for standard output")
	certFilename := flag.String("cert", "", "certificate")
	pkeyFilename := flag.String("key", "", "private key")

	flag.Parse()

	if *logFilename != "" {
		file, err := os.OpenFile(*logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(file)
	}

	mux := server.NewMux()

	httpSrv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		IdleTimeout:       5 * time.Minute,
		ReadHeaderTimeout: time.Minute,
	}

	// HTTPS Server
	if *certFilename != "" && *pkeyFilename != "" {
		cert, err := tls.LoadX509KeyPair(*certFilename, *pkeyFilename)
		if err != nil {
			log.Fatal(err)
		}

		httpsSrv := &http.Server{
			Addr:              ":10443",
			Handler:           mux,
			IdleTimeout:       5 * time.Minute,
			ReadHeaderTimeout: time.Minute,
			TLSConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
			},
		}

		go func() {
			log.Println("HTTPS Server started")
			log.Fatal(httpsSrv.ListenAndServeTLS("", ""))
		}()
	}

	// HTTP Server
	log.Println("HTTP Server started")
	log.Fatal(httpSrv.ListenAndServe())
}
