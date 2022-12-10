package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/FTBpro/go-workshop/coolfacts/coolfact"
)

const (
	pathGetFacts   = "/facts"
	pathCreateFact = "/facts"
)

type getFactsResponse struct {
	Facts []struct {
		Topic       string `json:"topic"`
		Description string `json:"description"`
		// TODO: add a field for createdAt
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

func (c *client) GetLastCreatedFact() (coolfact.Fact, error) {
	// TODO: implement this method.
	// Use the method GetFacts for getting all the facts
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

func (c *client) CreateFact(fact coolfact.Fact) error {
	ul := c.endpoint + pathCreateFact

	// First we are preparing the payload
	payload := map[string]interface{}{
		"topic":       fact.Topic,
		"description": fact.Description,
	}

	// we need io.Reader to create a new http request.
	// we will create bytes.Buffer which implement this interface
	postBody, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("client.CreateFact failed to marshal payload: %v", err)
	}
	responseBody := bytes.NewBuffer(postBody)

	// TODO:
	// 1. create a new request. Use http.NewRequestWithContext. For argument use the ul and the responseBody.
	// 2. Do the request using c.httpClient
	// 3. As in GetFacts, in case of a failure (response status code is not 200), return error using readError
	// * don't forget to close the body like we did in GetFacts method
}

type errorResponse struct {
	Error string `json:"error"`
}

func (c *client) readError(res *http.Response) (string, error) {
	var errRes errorResponse
	if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
		return "", fmt.Errorf("readBody failed to read response body: %v", err)
	}

	return errRes.Error, nil
}

func (c *client) readResponseGetFacts(res *http.Response) (getFactsResponse, error) {
	var factsRes getFactsResponse
	if err := json.NewDecoder(res.Body).Decode(&factsRes); err != nil {
		return getFactsResponse{}, fmt.Errorf("readResponseGetFacts failed to read response body: %v", err)
	}

	return factsRes, nil
}
