package usecase

import (
	"context"
	stderrors "errors"
	"fmt"
	"tourism/internal/common/errors"
	"tourism/internal/domain/event"
	"tourism/internal/infrastructure/repository"
)

type EventRepository interface {
	GetEvent(ctx context.Context, id int64) (*event.Event, error)
	GetEventsByCountry(ctx context.Context, countryID int64) ([]*event.EventsResponse, error)
}

type EventUseCase struct {
	eventRepo EventRepository
}

func NewEventUseCase(eventRepo EventRepository) *EventUseCase {
	return &EventUseCase{
		eventRepo: eventRepo,
	}
}

func (uc *EventUseCase) GetEvent(ctx context.Context, id int64) (*event.Event, error) {
	event, err := uc.eventRepo.GetEvent(ctx, id)
	if err != nil {
		if stderrors.Is(err, repository.ErrNotFound) {
			return nil, errors.NewNotFoundError("event was not found", "event")
		}
		return nil, fmt.Errorf("getting event %d: %w", id, err)
	}

	return event, nil
}

func (uc *EventUseCase) GetEventsByCountry(ctx context.Context, countryID int64) ([]*event.EventsResponse, error) {
	events, err := uc.eventRepo.GetEventsByCountry(ctx, countryID)
	if err != nil {
		return nil, fmt.Errorf("getting events for country %d: %w", countryID, err)
	}

	return events, nil
}
