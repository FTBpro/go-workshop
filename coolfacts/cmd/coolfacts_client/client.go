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
	// TODO: add fields
	// This struct represent the createFact API response body of the server.
	// We will decode the response into a variable of this struct type.
	// Since the server response is json, we will use json decode method.
	// For this be sure to add json tags on the struct. (https://gobyexample.com/json)
	// The response body is:
	// {
	//		"facts": [
	//			{
	//				"topic": "...",
	//				"description": "...",
	//          }
	//			...
	//		]
	// }
	//
}

func (r getFactsResponse) ToCoolFacts() []coolfact.Fact {
	// TODO: implement
	// loop over the response facts and convert them to the entity type []coolfact.Fact
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

func (c *client) GetAllFacts() ([]coolfact.Fact, error) {
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

	// TODO: handle response
	// this method returns *http.Response.
	//		- If response status code isn't 200 (http.StatusOK), you should read the error from the response.
	//			use method c.readError which is already implemented.
	//		- If the response is OK, use method readResponseGetFacts (which you will implement) to return the facts
}

type errorResponse struct {
	Error string `json:"error"`
}

func (c *client) readError(res *http.Response) (string, error) {
	var errRes errorResponse
	if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
		return "", fmt.Errorf("readBody failed to read response body: %v. \nbody string is: %s", err)
	}

	return errRes.Error, nil
}

func (c *client) readResponseGetFacts(res *http.Response) (getFactsResponse, error) {
	// TODO: implement - decode the json response into the target
	// Use variable of type getFactsResponse.
	// Use json.NewDecoder(...).Decode(...) (unlike the decoding in readError method)
}
