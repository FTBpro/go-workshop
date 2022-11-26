package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

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
		CreatedAt:   allFacts[0].CreatedAt,
	}, nil

}

func (c *client) GetAllFacts() ([]coolfact.Fact, error) {
	ul := c.endpoint + pathCreateFact
	res, err := c.httpClient.Get(ul)
	if err != nil {
		return nil, fmt.Errorf("client.GetLastCreatedFact to do request: %v", err)
	}

	defer func() {
		if res != nil && res.Body != nil {
			_, _ = io.Copy(ioutil.Discard, res.Body)
			_ = res.Body.Close()
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

	return getFactsRes.ToCoolFacts(), nil
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

	req, err := http.NewRequest(http.MethodPost, ul, responseBody)
	if err != nil {
		return fmt.Errorf("client.CreateFact failed to create request: %v", err)
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("client.CreateFact failed to do request: %v", err)
	}

	defer func() {
		if res != nil && res.Body != nil {
			_, _ = io.Copy(ioutil.Discard, res.Body)
			_ = res.Body.Close()
		}
	}()

	if res.StatusCode != http.StatusOK {
		errMessage, err := c.readError(res)
		if err != nil {
			return fmt.Errorf("client.CreateFact: %s", err)
		}

		return fmt.Errorf("client.CreateFact got an error from server. status: %d. error: %s", res.StatusCode, errMessage)
	}

	return nil
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

type getFactsResponse struct {
	Facts []struct {
		Image       string    `json:"image"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"createdAt"`
	} `json:"facts"`
}

func (r getFactsResponse) ToCoolFacts() []coolfact.Fact {
	coolfacts := make([]coolfact.Fact, len(r.Facts))
	for i, fact := range r.Facts {
		coolfacts[i] = coolfact.Fact(fact)
	}

	return coolfacts
}

func (r getFactsResponse) getLastFact() coolfact.Fact {
	if len(r.Facts) == 0 {
		return coolfact.Fact{}
	}

	return coolfact.Fact{
		Image:       r.Facts[0].Image,
		Description: r.Facts[0].Description,
		CreatedAt:   r.Facts[0].CreatedAt,
	}
}

func (r getFactsResponse) ToCoolFacts() []coolfact.Fact {
	// TODO: implement
	// loop over the response facts and convert them to the entity type []coolfact.Fact
}

func (c *client) readResponseGetFacts(res *http.Response) (getFactsResponse, error) {
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return getFactsResponse{}, fmt.Errorf("readResponseGetFacts failed to read response body: %v", err)
	}

	var factsRes getFactsResponse
	if err = json.Unmarshal(resBody, &factsRes); err != nil {
		return getFactsResponse{}, fmt.Errorf("readResponseGetFacts failed to read response body: %v. \nbody string is: %s", err, string(resBody))
	}

	return factsRes, nil
}
