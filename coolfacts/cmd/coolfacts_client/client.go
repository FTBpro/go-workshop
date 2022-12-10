package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/FTBpro/go-workshop/coolfacts/coolfact"
)

const (
	pathGetFacts = "/facts"
)

type getFactsResponse struct {
	Facts []struct {
		Topic       string `json:"topic"`
		Description string `json:"description"`
	} `json:"facts"`
}

func (r getFactsResponse) toCoolFacts() []coolfact.Fact {
	coolfacts := make([]coolfact.Fact, len(r.Facts))
	for i, fact := range r.Facts {
		coolfacts[i] = coolfact.Fact(fact)
	}

	return coolfacts
}

type client struct {
	endpoint   string
	httpClient *http.Client
}

func NewClient(endpoint string) *client {
	return &client{
		endpoint:   endpoint,
		httpClient: &http.Client{},
	}
}

func (c *client) GetFacts() ([]coolfact.Fact, error) {
	ul := c.endpoint + pathGetFacts
	res, err := c.httpClient.Get(ul)
	if err != nil {
		return nil, fmt.Errorf("client.GetLastCreatedFact to do request: %v", err)
	}

	// The client must close the body after the response is handled
	// We must read all the body before closing it, so for reading the body and copying to ioutil.Discard, which does nothing
	defer func() {
		if res != nil && res.Body != nil {
			io.Copy(ioutil.Discard, res.Body)
			res.Body.Close()
		}
	}()

	if res.StatusCode != http.StatusOK {
		errMessage, err := c.readError(res)
		if err != nil {
			return nil, fmt.Errorf("client.CreateFact: %s", err)
		}

		return nil, fmt.Errorf("client.GetLastCreatedFact got an error from server. status: %d. error: %s", res.StatusCode, errMessage)
	}

	getFactsRes, err := c.readResponseGetFacts(res)
	if err != nil {
		return nil, fmt.Errorf("client.GetLastCreatedFact: %s", err)
	}

	return getFactsRes.toCoolFacts(), nil
}

type errorResponse struct {
	Error string `json:"error"`
}

func (c *client) readError(res *http.Response) (string, error) {
	var errRes errorResponse
	if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
		return "", fmt.Errorf("readBody failed to read response body: %v. \n", err)
	}

	return errRes.Error, nil
}

func (c *client) readResponseGetFacts(res *http.Response) (getFactsResponse, error) {
	var factsRes getFactsResponse
	if err := json.NewDecoder(res.Body).Decode(&factsRes); err != nil {
		return getFactsResponse{}, fmt.Errorf("readResponseGetFacts failed to read response body: %v. \nbody string is: %s", err)
	}

	return factsRes, nil
}
