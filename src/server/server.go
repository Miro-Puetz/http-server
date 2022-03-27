package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", noContent)
	mux.HandleFunc("/index", DefaultHandler)
	mux.Handle("/files/", http.StripPrefix("/files", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/status-codes/", handleCodes)
	return mux
}

func noContent(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	var response string
	switch r.Method {
	case http.MethodGet:
		response = "You have send me a get request"
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		response = fmt.Sprintf("You have send me a post request: %s", body)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	fmt.Fprint(w, string(response))
}

func handleCodes(w http.ResponseWriter, r *http.Request) {
	code := strings.Split(r.URL.Path, "/")[2]
	switch code {
	case "200":
		w.WriteHeader(http.StatusOK)
	case "203":
		w.WriteHeader(http.StatusNonAuthoritativeInfo)
	case "400":
		w.WriteHeader(http.StatusBadRequest)
	case "404":
		w.WriteHeader(http.StatusNotFound)
	case "500":
		w.WriteHeader(http.StatusInternalServerError)
	default:
		w.WriteHeader(http.StatusNoContent)
	}
}
