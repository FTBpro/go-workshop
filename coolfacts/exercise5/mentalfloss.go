package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type mentalfloss struct {}

func (mf mentalfloss) Facts()([]fact, error) {
	resp, err := http.Get("http://mentalfloss.com/api/facts")
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}

	// A `defer` statement defers the execution of a function until the surrounding function returns.
	// This is how we make sure that we close the response body before we exit the function.
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error readAll: %v", err)
	}

	facts, err := parseFromRawItems(b)
	if err != nil {
		return nil, fmt.Errorf("error parsing data: %v", err)
	}

	return facts, nil
}

func parseFromRawItems(b []byte) ([]fact, error) {
	var items []struct {
		FactText     string `json:"fact"`
		PrimaryImage string `json:"primaryImage"`
	}
	if err := json.Unmarshal(b, &items); err != nil {
		return nil, err
	}

	var facts []fact
	for _, it := range items {
		newFact := fact{
			Image:       it.PrimaryImage,
			Description: it.FactText,
		}
		facts = append(facts, newFact)
	}

	return facts, nil
}