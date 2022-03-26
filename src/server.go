package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

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
