package coolfact

import "time"

type Fact struct {
	Topic       string
	Description string
	CreatedAt   time.Time
}

type Filters struct {
	Topic string
	Limit int
}
