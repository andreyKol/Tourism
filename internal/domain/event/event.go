package event

import (
	"time"
)

type Event struct {
	ID          int64
	Name        string
	Description string
	CountryID   int64
	Date        time.Time
}

type EventsResponse struct {
	ID          int64
	Name        string
	Description string
	CountryID   int64
	Date        time.Time
}
