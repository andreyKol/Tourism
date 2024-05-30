package country

import "tourism/internal/domain/event"

type Country struct {
	ID          int64
	Name        string
	Description string
	Image       string
	Events      map[string]*event.Event
}

type CountriesResponse struct {
	ID          string
	Name        string
	Description string
	Image       string
}
