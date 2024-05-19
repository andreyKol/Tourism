package http

import (
	"context"
	"github.com/go-chi/render"
	"net/http"
	"tourism/internal/domain/event"
	"tourism/internal/handlers/httphelp"
)

type EventUseCase interface {
	GetEvent(ctx context.Context, id int64) (*event.Event, error)
	GetEventsByCountry(ctx context.Context, countryID int64) ([]*event.EventsResponse, error)
}

// @Summary      GetEvent
// @Description  Returns information about an event
// @Tags         event
// @Accept       json
// @Produce      json
// @Param        event_id path int true "Event ID"
// @Success      200 {object} event.Event
// @Failure      404  {object} Error
// @Failure      500  {object} Error
// @Router       /events/{eventId} [get]
func (h HttpHandler) GetEvent(w http.ResponseWriter, r *http.Request) {
	eventID := httphelp.ParseParamInt64("eventId", r)
	event, err := h.eventUseCase.GetEvent(r.Context(), eventID)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	render.JSON(w, r, event)
}

// @Summary      GetEventsByCountry
// @Description  Returns a list of events for a specific country
// @Tags         event
// @Accept       json
// @Produce      json
// @Param        country_id path int true "Country ID"
// @Success      200 {array} event.EventsResponse
// @Failure      404  {object} Error
// @Failure      500  {object} Error
// @Router       /events/{country_id} [get]
func (h HttpHandler) GetEventsByCountry(w http.ResponseWriter, r *http.Request) {
	countryID := httphelp.ParseParamInt64("countryId", r)
	events, err := h.eventUseCase.GetEventsByCountry(r.Context(), countryID)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	var response []*event.EventsResponse
	for _, e := range events {
		response = append(response, e)
	}

	render.JSON(w, r, response)
}
