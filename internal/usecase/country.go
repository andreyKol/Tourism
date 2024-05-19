package usecase

import (
	"context"
	stderrors "errors"
	"fmt"
	"tourism/internal/common/errors"
	"tourism/internal/domain/country"
	"tourism/internal/infrastructure/repository"
)

type CountryRepository interface {
	GetCountry(ctx context.Context, id int64) (*country.Country, error)
	GetCountries(ctx context.Context) ([]*country.CountriesResponse, error)
}

type CountryUseCase struct {
	countryRepo CountryRepository
}

func NewCountryUseCase(countryRepo CountryRepository) *CountryUseCase {
	return &CountryUseCase{
		countryRepo: countryRepo,
	}
}

func (uc *CountryUseCase) GetCountry(ctx context.Context, id int64) (*country.Country, error) {
	country, err := uc.countryRepo.GetCountry(ctx, id)
	if err != nil {
		if stderrors.Is(err, repository.ErrNotFound) {
			return nil, errors.NewNotFoundError("country was not found", "country")
		}
		return nil, fmt.Errorf("getting country %d: %w", id, err)
	}

	return country, nil
}

func (uc *CountryUseCase) GetCountries(ctx context.Context) ([]*country.CountriesResponse, error) {
	countries, err := uc.countryRepo.GetCountries(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting countries: %w", err)
	}

	return countries, nil
}
