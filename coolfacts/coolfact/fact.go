package coolfact

import "time"

type Fact struct {
	Topic       string
	Description string
	CreatedAt   time.Time
}

// TODO: add struct Filters with:
// - Topic
// - Limit
