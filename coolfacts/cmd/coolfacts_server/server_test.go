package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	server "github.com/FTBpro/go-workshop/coolfacts/cmd/coolfacts_server"
	"github.com/FTBpro/go-workshop/coolfacts/coolfact"
)

func Test_Server_GetFacts(t *testing.T) {
	facts := generateRandomFactsDesc(10)
	tests := []struct {
		name               string
		queryParamsToSend  string
		expectedFilters    coolfact.Filters
		want               []coolfact.Fact
		wantErr            bool
		expectedHTTPStatus int
	}{
		{
			name:              "10 facts with filters",
			queryParamsToSend: "?limit=10&topic=TV",
			expectedFilters: coolfact.Filters{
				Topic: "TV",
				Limit: 10,
			},
			want:               facts,
			expectedHTTPStatus: http.StatusOK,
		},
		{
			name:              "no topic",
			queryParamsToSend: "?limit=10",
			expectedFilters: coolfact.Filters{
				Topic: "",
				Limit: 10,
			},
			want:               facts,
			expectedHTTPStatus: http.StatusOK,
		},
		{
			name:               "no limit - expect bad request",
			queryParamsToSend:  "",
			want:               nil,
			wantErr:            true,
			expectedHTTPStatus: http.StatusBadRequest,
		},
		{
			name:               "limit is not an int - expect bad request",
			queryParamsToSend:  "?limit=one",
			want:               nil,
			wantErr:            true,
			expectedHTTPStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := mockFactsService{
				factsToReturn: tt.want,
			}

			srv := server.NewServer(&mockService)
			ts := httptest.NewServer(srv)

			res, err := http.Get(ts.URL + "/facts" + tt.queryParamsToSend)
			require.NoError(t, err)
			require.Equal(t, tt.expectedHTTPStatus, res.StatusCode)

			if tt.wantErr {
				return
			}

			gotFacts, err := factsFromResponse(t, res)
			require.NoError(t, err)

			require.Equal(t, tt.expectedFilters, mockService.filtersGot)
			expectEqualFacts(t, tt.want, gotFacts)
		})
	}
}

func Test_Server_CreateFacts(t *testing.T) {
	facts := generateRandomFactsDesc(10)
	tests := []struct {
		name               string
		queryParamsToSend  string
		factToCreate       coolfact.Fact
		wantErr            bool
		expectedHTTPStatus int
	}{
		{
			name:               "10 facts with filters",
			factToCreate:       facts[0],
			expectedHTTPStatus: http.StatusOK,
		},
		{
			name:               "no topic",
			factToCreate:       facts[0],
			expectedHTTPStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := mockFactsService{}
			srv := server.NewServer(&mockService)
			ts := httptest.NewServer(srv)

			payload := map[string]interface{}{
				"topic":       tt.factToCreate.Topic,
				"description": tt.factToCreate.Description,
			}

			postBody, err := json.Marshal(payload)
			require.NoError(t, err)

			responseBody := bytes.NewBuffer(postBody)

			res, err := http.Post(ts.URL+"/facts", "application/json", responseBody)
			require.NoError(t, err)
			require.Equal(t, tt.expectedHTTPStatus, res.StatusCode)

			if tt.wantErr {
				return
			}
			require.Equal(t, tt.factToCreate.Topic, mockService.createFactGotFact.Topic)
			require.Equal(t, tt.factToCreate.Description, mockService.createFactGotFact.Description)
		})
	}
}

func generateRandomFactsDesc(n int) []coolfact.Fact {
	var facts []coolfact.Fact
	for i := 0; i < n; i++ {
		fact := randomFact()
		fact.CreatedAt = time.Now().Add(-(time.Duration(i) * time.Hour)).UTC()
		facts = append(facts, fact)
	}

	return facts
}

func randomFact() coolfact.Fact {
	rand.Seed(time.Now().UnixNano())
	return coolfact.Fact{
		Topic:       fmt.Sprintf("Topic %d", rand.Intn(10000)),
		Description: fmt.Sprintf("Some Description %d", rand.Intn(10000)),
	}
}

type getFactsResponse struct {
	Facts []struct {
		Topic       string    `json:"topic"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"createdAt"`
	} `json:"facts"`
}

func factsFromResponse(t *testing.T, res *http.Response) ([]coolfact.Fact, error) {
	var factsResponse getFactsResponse
	err := json.NewDecoder(res.Body).Decode(&factsResponse)
	require.NoErrorf(t, err, "factsFromResponse failed decode get facts response")

	facts := make([]coolfact.Fact, len(factsResponse.Facts))
	for i, fact := range factsResponse.Facts {
		facts[i] = coolfact.Fact(fact)
	}

	return facts, nil
}

type mockFactsService struct {
	filtersGot        coolfact.Filters
	factsToReturn     []coolfact.Fact
	shouldReturnError bool

	createFactGotFact coolfact.Fact
}

func (m *mockFactsService) GetFacts(filters coolfact.Filters) ([]coolfact.Fact, error) {
	m.filtersGot = filters

	if m.shouldReturnError {
		return nil, fmt.Errorf("mockFactsService asked to return an error")
	}

	return m.factsToReturn, nil
}

func (m *mockFactsService) CreateFact(fact coolfact.Fact) error {
	m.createFactGotFact = fact

	if m.shouldReturnError {
		return fmt.Errorf("mockFactsService asked to return an error")
	}

	return nil
}

func expectEqualFacts(t *testing.T, expected, got []coolfact.Fact) {
	require.Equalf(t, len(expected), len(got), "expectEqualFacts: different length")

	for i, gotFact := range got {
		expectedFact := expected[i]
		require.Equal(t, expectedFact.Topic, gotFact.Topic)
		require.Equal(t, expectedFact.Description, gotFact.Description)
		fmt.Println("------------------ ", gotFact.CreatedAt)
		require.Equal(t, expectedFact.CreatedAt, gotFact.CreatedAt)
	}
}
