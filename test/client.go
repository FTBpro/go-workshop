package test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin/json"
)

type Client struct {
	ServerURL string
}

func (c Client) Do(t *testing.T, method, path string, body interface{}) *http.Response {
	t.Helper()
	var bodyReader io.Reader
	if body != nil {
		bodyJ, err := json.Marshal(body)
		if err != nil {
			t.Fatalf("bug in test %s: error marshaling body", t.Name())
		}
		bodyReader = bytes.NewReader(bodyJ)
	}
	req, err := http.NewRequest(method, c.ServerURL+path, bodyReader)
	if err != nil {
		t.Fatalf("bug in test %s: error preparing request", t.Name())
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("error from http client: %s", err)
	}
	return resp
}

func (c Client) Ping(t *testing.T) {
	resp := c.Do(t, "GET", "/ping", nil)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("ping should return status code 200 OK, but got %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body: %s", body)
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			t.Fatalf("error closing response body: %s", err)
		}
	}()

	if string(body) != "PONG" {
		t.Error(fmt.Sprintf("expected `PONG`, got `%s`", string(body)))
	}
}
