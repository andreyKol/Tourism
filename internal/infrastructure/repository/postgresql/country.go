package postgresql

import (
	"context"
	"tourism/internal/domain/country"
)

func (r *Repository) GetCountry(ctx context.Context, id int64) (*country.Country, error) {
	var c country.Country

	err := r.db.QueryRow(ctx, `
		select id,
		       name,
		       description
		from countries
		where id = $1`, id).Scan(
		&c.ID,
		&c.Name,
		&c.Description,
	)
	if err != nil {
		return nil, parseError(err, "selecting country")
	}

	return &c, nil
}

func (r *Repository) GetCountries(ctx context.Context) ([]*country.CountriesResponse, error) {
	rows, err := r.db.Query(ctx, `
        SELECT c.id, c.name, c.description
        FROM countries c
    `)
	if err != nil {
		return nil, parseError(err, "selecting countries")
	}
	defer rows.Close()

	var countries []*country.CountriesResponse

	for rows.Next() {
		var country country.CountriesResponse
		if err = rows.Scan(&country.ID, &country.Name, &country.Description); err != nil {
			return nil, parseError(err, "scanning country row")
		}
		countries = append(countries, &country)
	}

	if err = rows.Err(); err != nil {
		return nil, parseError(err, "iterating over country rows")
	}

	return countries, nil
}
