package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", DefaultHandler)
	mux.Handle("/files/", http.StripPrefix("/files", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/status-codes/", handleCodes)
	return mux
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
	err := tpl.ExecuteTemplate(w, "index.gohtml", response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
}

func handleCodes(w http.ResponseWriter, r *http.Request) {
	var code string
	splitUrl := strings.Split(r.URL.Path, "/")
	fmt.Println(splitUrl, len(splitUrl))
	if len(splitUrl) > 2 {
		code = splitUrl[2]
	}
	fmt.Println(code)
	if code != "" {
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
	} else {
		err := tpl.ExecuteTemplate(w, "status-codes.gohtml", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
		}
	}
}
