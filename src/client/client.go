package client

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func doClientRequest(url string) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
	}

	resp, err := http.DefaultClient.Do(req)
	fmt.Println(resp)
	if err != nil {
		log.Println(err)
	}

	_ = resp.Body.Close()
	return resp, err
}
