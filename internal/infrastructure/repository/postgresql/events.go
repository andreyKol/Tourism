package postgresql

import (
	"context"
	"tourism/internal/domain/event"
)

func (r *Repository) GetEvent(ctx context.Context, id int64) (*event.Event, error) {
	var e event.Event

	err := r.db.QueryRow(ctx, `
        SELECT id,
               name,
               description,
               country_id,
               image,
               date
        FROM events
        WHERE id = $1`, id).Scan(
		&e.ID,
		&e.Name,
		&e.Description,
		&e.CountryID,
		&e.Image,
		&e.Date,
	)
	if err != nil {
		return nil, parseError(err, "selecting event")
	}

	return &e, nil
}

func (r *Repository) GetEventsByCountry(ctx context.Context, countryID int64) ([]*event.EventsResponse, error) {
	rows, err := r.db.Query(ctx, `
        SELECT e.id, e.name, e.description, e.country_id, e.image, e.date 
        FROM events e
        WHERE e.country_id = $1
    `, countryID)
	if err != nil {
		return nil, parseError(err, "selecting events by country")
	}
	defer rows.Close()

	var events []*event.EventsResponse

	for rows.Next() {
		var event event.EventsResponse
		if err = rows.Scan(&event.ID, &event.Name, &event.Description, &event.CountryID, &event.Image, &event.Date); err != nil {
			return nil, parseError(err, "scanning event row")
		}
		events = append(events, &event)
	}

	if err = rows.Err(); err != nil {
		return nil, parseError(err, "iterating over event rows")
	}

	return events, nil
}
