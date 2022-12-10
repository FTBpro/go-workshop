package coolfact_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/FTBpro/go-workshop/coolfacts/coolfact"
	"github.com/FTBpro/go-workshop/coolfacts/inmem"
)

func Test_service_GetFacts(t *testing.T) {
	facts := generateRandomFactsDesc(10)

	testCases := []struct {
		name      string
		repoField coolfact.Repository
		want      []coolfact.Fact
		wantErr   bool
	}{
		{
			name:      "with facts",
			repoField: inmem.NewFactsRepository(facts...),
			want:      facts,
			wantErr:   false,
		},
		{
			name:      "add unsorted facts",
			repoField: inmem.NewFactsRepository(facts[5], facts[4], facts[2]),
			want:      []coolfact.Fact{facts[2], facts[4], facts[5]},
			wantErr:   false,
		},
		{
			name:      "no facts - should get nil",
			repoField: inmem.NewFactsRepository(),
			want:      nil,
			wantErr:   false,
		},
		{
			name:      "repo returns error",
			repoField: mockRepoError{},
			want:      nil,
			wantErr:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := coolfact.NewService(tc.repoField)
			got, err := s.GetFacts()
			if err != nil {
				require.True(t, tc.wantErr, "got an unexpected error from service")
				return
			}

			require.False(t, tc.wantErr, "expected an error but didn't receive one.")
			expectEqualFacts(t, tc.want, got)

		})
	}
}

func Test_service_CreateFact(t *testing.T) {
	facts := generateRandomFactsDesc(10)

	tests := []struct {
		name          string
		repoField     coolfact.Repository
		factsToCreate []coolfact.Fact
		want          []coolfact.Fact
		wantErr       bool
	}{
		{
			name:          "base case - adding sorted",
			repoField:     inmem.NewFactsRepository(facts[5]),
			factsToCreate: []coolfact.Fact{facts[2]},
			want:          []coolfact.Fact{facts[2], facts[5]},
		},
		{
			name:          "add fact from the past",
			repoField:     inmem.NewFactsRepository(facts[3]),
			factsToCreate: []coolfact.Fact{facts[5]},
			want:          []coolfact.Fact{facts[3], facts[5]},
		},
		{
			name:          "add many mixed facts",
			repoField:     inmem.NewFactsRepository(facts[3], facts[5], facts[1], facts[9]),
			factsToCreate: []coolfact.Fact{facts[4], facts[2], facts[0]},
			want:          []coolfact.Fact{facts[0], facts[1], facts[2], facts[3], facts[4], facts[5], facts[9]},
		},
		{
			name:          "repo returns error",
			repoField:     mockRepoError{},
			factsToCreate: []coolfact.Fact{facts[4], facts[2], facts[0]},
			want:          nil,
			wantErr:       true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			s := coolfact.NewService(tc.repoField)
			for _, fact := range tc.factsToCreate {
				err := s.CreateFact(fact)
				if err != nil {
					require.True(t, tc.wantErr, "got an unexpected error from service")
					return
				}

				require.False(t, tc.wantErr, "expected an error but didn't receive one.")
				return
			}

			gotFacts, err := s.GetFacts()
			require.NoError(t, err)
			require.Equal(t, gotFacts, tc.want)
		})
	}
}

// generateRandomFactsDesc creates new random facts sorted by DESC
func generateRandomFactsDesc(n int) []coolfact.Fact {
	var facts []coolfact.Fact
	for i := 0; i < n; i++ {
		fact := randomFact()
		fact.CreatedAt = time.Now().Add(-(time.Duration(i) * time.Hour))
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

func expectEqualFacts(t *testing.T, expected, got []coolfact.Fact) {
	require.Equalf(t, len(expected), len(got), "expectEqualFacts: different length")

	for _, fact := range got {
		require.Contains(t, expected, fact, "expectEqualFacts: got unexpected fact")
	}

	for _, fact := range expected {
		require.Contains(t, got, fact, "expectEqualFacts: didn't got expected fact")
	}
}

type mockRepoError struct {
}

func (m mockRepoError) GetFacts() ([]coolfact.Fact, error) {
	return nil, fmt.Errorf("mock repo returns error")
}

func (m mockRepoError) CreateFact(fact coolfact.Fact) error {
	return fmt.Errorf("mock repo returns error")
}
