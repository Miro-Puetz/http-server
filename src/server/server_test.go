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
		method string
		body   io.Reader
		code   int
	}{
		{http.MethodGet, nil, http.StatusOK},
		{http.MethodGet, bytes.NewBufferString("Can I have a body?"), http.StatusOK},
		{http.MethodPost, bytes.NewBufferString("I hava a small body"), http.StatusOK},
		{http.MethodPost, nil, http.StatusOK},
		{http.MethodHead, nil, http.StatusMethodNotAllowed},
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

		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("%d: %v", i, err)
		}
	}
}
