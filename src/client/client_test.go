package client

import (
	"net/http"
	"testing"
)

func TestClient(t *testing.T) {
	url := "http://:8080/index"
	resp, err := doClientRequest(url)
	if err != nil {
		t.Errorf("%v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: %q", resp.StatusCode)
	}
}
