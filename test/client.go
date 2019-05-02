package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

type Client struct {
	ServerURL string
}

func (c Client) Ping(t *testing.T) {
	req, err := http.NewRequest("GET", c.ServerURL+"/ping", nil)
	if err != nil {
		t.Fatalf("bug in test %s: error preparing request", t.Name())
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("error from http client: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("ping should return status code 200 OK, but got %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body: %s",body)
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			t.Errorf("error closing response body: %s", err)
		}
	}()

	if string(body) != "PONG" {
		t.Error(fmt.Sprintf("expected `PONG`, got `%s`", string(body)))
	}
}
