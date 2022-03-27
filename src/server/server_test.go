package server

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestRequests(t *testing.T) {
	testCases := []struct {
		method   string
		body     io.Reader
		code     int
		response string
	}{
		{http.MethodGet, nil, http.StatusOK, "You have send me a get request"},
		{http.MethodGet, bytes.NewBufferString("Can I have a body?"), http.StatusOK, "You have send me a get request"},
		{http.MethodPost, bytes.NewBufferString("I am a small body"), http.StatusOK, "You have send me a post request: I am a small body"},
		{http.MethodPost, nil, http.StatusOK, "You have send me a post request: "},
		{http.MethodHead, nil, http.StatusMethodNotAllowed, ""},
	}

	client := new(http.Client)
	path := "http://:8080"

	for i, c := range testCases {
		request, err := http.NewRequest(c.method, path, c.body)
		if err != nil {
			t.Errorf("%d: %v", i, err)
			continue
		}

		resp, err := client.Do(request)
		if err != nil {
			t.Errorf("%d: %v", i, err)
			continue
		}

		if resp.StatusCode != c.code {
			t.Errorf("%d: unexpected status code: %q", i, resp.StatusCode)
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("%d: %v", i, err)
		}
		if string(b) != c.response {
			t.Errorf("%d: expected %q; actual %q", i, c.response, b)
		}
	}
}
