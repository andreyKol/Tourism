package http

import (
	"context"
	"github.com/go-chi/render"
	"net/http"
	"tourism/internal/domain/country"
	"tourism/internal/handlers/httphelp"
)

type CountryUseCase interface {
	GetCountry(ctx context.Context, id int64) (*country.Country, error)
	GetCountries(ctx context.Context) ([]*country.CountriesResponse, error)
}

// @Summary      GetCountry
// @Description  Returns information about country
// @Tags         country
// @Accept       json
// @Produce      json
// @Param		 country_id path int true "Country ID"
// @Success      200 {object} country.Country
// @Failure      404  {object} Error
// @Failure      500  {object} Error
// @Router       /countries/{countryId} [get]
func (h HttpHandler) GetCountry(w http.ResponseWriter, r *http.Request) {
	countryID := httphelp.ParseParamInt64("countryId", r)
	country, err := h.countryUseCase.GetCountry(r.Context(), countryID)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	render.JSON(w, r, country)
}

// @Summary      GetCountries
// @Description  Returns list of all countries
// @Tags         country
// @Accept       json
// @Produce      json
// @Success      200 {array} country.CountriesResponse
// @Failure      500  {object}  Error
// @Router       /countries [get]
func (h HttpHandler) GetCountries(w http.ResponseWriter, r *http.Request) {
	countries, err := h.countryUseCase.GetCountries(r.Context())
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	var response []*country.CountriesResponse
	for _, c := range countries {
		response = append(response, c)
	}

	render.JSON(w, r, response)
}
