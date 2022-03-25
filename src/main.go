package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(defaultHandler))
	http.ListenAndServe(":8080", nil)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Word!")
}
