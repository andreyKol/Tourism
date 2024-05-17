package postgresql

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"Tourism/internal/infrastructure/repository"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func parseError(err error, prefix string) error {
	if errors.Is(err, pgx.ErrNoRows) {
		err = repository.ErrNotFound
	}

	if prefix != "" {
		return fmt.Errorf("%s: %w", prefix, err)
	}

	return err
}
