package main

import (
	"net/http"

	server "github.com/miro-puetz/http-server/src"
)

func main() {
	http.Handle("/", http.HandlerFunc(server.DefaultHandler))
	http.ListenAndServe(":8080", nil)
}
