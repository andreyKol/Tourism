package event

import (
	"time"
)

type Event struct {
	ID          int64
	Name        string
	Description string
	CountryID   int64
	Image       string
	Date        time.Time
}

type EventsResponse struct {
	ID          int64
	Name        string
	Description string
	CountryID   int64
	Image       string
	Date        time.Time
}
