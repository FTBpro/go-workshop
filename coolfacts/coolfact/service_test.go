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

// generateRandomFactsDesc creates new random facts sorted by DESC
func generateRandomFactsDesc(n int) []coolfact.Fact {
	var facts []coolfact.Fact
	for i := 0; i < n; i++ {
		fact := randomFact()
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
