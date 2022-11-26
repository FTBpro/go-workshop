package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	
	"github.com/FTBpro/go-workshop/coolfacts/coolfact"
)

const (
	pathGetFacts   = "/facts"
	pathCreateFact = "/facts"
)

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

type getFactsResponse struct {
	// TODO: add fields
	// This struct represent the createFact API response body of the server.
	// We will use this struct to
	// The response is json:
	// {
	//		"facts": [
	//			{
	//				"image": "...",
	//				"description": "...",
	//				"createdAt": "...",
	//          }
	//			...
	//		]
	// }
	//
	
}

func (c *client) GetLastCreatedFact() (coolfact.Fact, error) {
	allFacts, err := c.GetAllFacts()
	if err != nil {
		return coolfact.Fact{}, fmt.Errorf("GetLastCreatedFact: %v", err)
	}
	
	if len(allFacts) == 0 {
		return coolfact.Fact{}, fmt.Errorf("GetLastCreatedFact didn't find facts")
	}
	
	return coolfact.Fact{
		Image:       allFacts[0].Image,
		Description: allFacts[0].Description,
		CreateAt:    allFacts[0].CreatedAt,
	}, nil
	
}

func (c *client) GetAllFacts() ([]coolfact.Fact, error) {
	// TODO: implement
	// 1. For calling a simple GET request, you can use the c.httpClient.Get method.
	//	  You just need to build the url. Use the client endpoint and the const for the path
	// 2. this method returns *http.Response.
	//		- The client must close the body of the request after use. Do this in defer
	//		- If response status code isn't 200 (http.StatusOK), you should read the error from the response.
	//			use method c.readError which is already implemented.
	//		- If the response is OK, use method readResponseGetFacts (which you will implement) to return the facts
}

func (c *client) CreateFact(fct coolfact.Fact) error {
	ul := c.endpoint + pathCreateFact
	
	// First we are preparing the payload
	payload := map[string]interface{}{
		"image":       fct.Image,
		"description": fct.Description,
	}
	
	// we need io.Reader to create a new http request.
	// we will create bytes.Buffer which implement this interface
	postBody, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("client.CreateFact failed to marshal payload: %v", err)
	}
	responseBody := bytes.NewBuffer(postBody)
	
	// TODO:
	// 1. create a new request. Use http.NewRequestWithContext
	// 2. Do the request using c.httpClient
	// 3. Ad in GetAllFacts, in case of a failure (response status code is not 200), return error using ReadErro
}

type errorResponse struct {
	Error string `json:"error"`
}

func (c *client) readError(res *http.Response) (string, error) {
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("readBody failed to read response body: %v", err)
	}
	
	var errRes errorResponse
	if err = json.NewDecoder(res.Body).Decode(&errRes); err != nil {
		return "", fmt.Errorf("readBody failed to read response body: %v. \nbody string is: %s", err, string(resBody))
	}
	
	if errRes.Error == "" {
		return string(resBody), nil
	}
	
	return errRes.Error, nil
}

func (c *client) readResponseGetFacts(res *http.Response) (getFactsResponse, error) {
	// TODO: implement - decode the json response into the target
	// Use variable of type getFactsResponse.
	// Use json.NewDecoder(...).Decode(...) (unlike the decoding in readError method)
}
