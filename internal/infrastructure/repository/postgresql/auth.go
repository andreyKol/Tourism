package postgresql

import (
	"context"
	"tourism/internal/domain"
)

func (r *Repository) GetUserByPhone(ctx context.Context, phone string) (*domain.User, error) {
	var user domain.User

	err := r.db.QueryRow(ctx, `
		select id,
		       name,
		       surname,
		       patronymic,
		       password_enc
		from users
		where phone = $1`, phone).Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Patronymic,
		&user.PasswordEncrypted,
	)
	if err != nil {
		return nil, parseError(err, "selecting user")
	}

	return &user, nil
}

func (r *Repository) GetAuthContextByUserID(ctx context.Context, id int64) (*domain.AuthContext, error) {
	ac := domain.AuthContext{ID: id}

	err := r.db.QueryRow(ctx, `
		select id, name
		from users
		where id=$1`, id,
	).Scan(&ac.ID, &ac.Name)
	if err != nil {
		return nil, parseError(err, "selecting email")
	}

	return &ac, nil
}
