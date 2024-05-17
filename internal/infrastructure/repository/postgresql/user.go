package postgresql

import (
	"Tourism/internal/domain/users"
	"context"
	"fmt"
	"strings"

	"Tourism/internal/domain"
)

func (r *Repository) GetUser(ctx context.Context, id int64) (*domain.User, error) {
	var u domain.User

	err := r.db.QueryRow(ctx, `
		select id,
		       name,
		       surname,
		       patronymic,
		       role,
		       age,
		       gender,
		       image_id,
		       phone,
		       email,
		       created_at,
		       last_online
		from users
		where id = $1`, id).Scan(
		&u.ID,
		&u.Name,
		&u.Surname,
		&u.Patronymic,
		&u.Role,
		&u.Age,
		&u.Gender,
		&u.ImageID,
		&u.Phone,
		&u.Email,
		&u.CreatedAt,
		&u.LastOnline,
	)
	if err != nil {
		return nil, parseError(err, "selecting user")
	}

	return &u, nil
}

func (r *Repository) GetUsersByUserRole(ctx context.Context, role int16) ([]*domain.User, error) {
	var usersSlice []*domain.User

	rows, err := r.db.Query(ctx, `
        SELECT id,
               name,
               surname,
               patronymic,
               role,
               age,
               gender,
               image_id,
               phone,
               email,
               created_at,
               last_online
        FROM users
        WHERE role = $1
    `, role)
	if err != nil {
		return nil, parseError(err, "selecting users by role")
	}
	defer rows.Close()

	for rows.Next() {
		var u domain.User
		if err = rows.Scan(
			&u.ID,
			&u.Name,
			&u.Surname,
			&u.Patronymic,
			&u.Role,
			&u.Age,
			&u.Gender,
			&u.ImageID,
			&u.Phone,
			&u.Email,
			&u.CreatedAt,
			&u.LastOnline,
		); err != nil {
			return nil, parseError(err, "scanning user row")
		}
		usersSlice = append(usersSlice, &u)
	}
	if err = rows.Err(); err != nil {
		return nil, parseError(err, "iterating over user rows")
	}

	return usersSlice, nil
}

func (r *Repository) UpdateUser(ctx context.Context, user *domain.User) error {
	args := []any{user.ID}
	var fields []string
	argID := 2

	fields = append(fields, fmt.Sprintf("name=$%d", argID))
	argID++
	args = append(args, user.Name)

	fields = append(fields, fmt.Sprintf("phone=$%d", argID))
	argID++
	args = append(args, user.Phone)

	if user.Age != nil {
		fields = append(fields, fmt.Sprintf("age=$%d", argID))
		argID++
		args = append(args, *user.Age)
	}

	if user.Surname != nil {
		fields = append(fields, fmt.Sprintf("surname=$%d", argID))
		argID++
		args = append(args, *user.Surname)
	}

	if user.Patronymic != nil {
		fields = append(fields, fmt.Sprintf("patronymic=$%d", argID))
		argID++
		args = append(args, *user.Patronymic)
	}

	if user.Gender != nil {
		fields = append(fields, fmt.Sprintf("gender=$%d", argID))
		argID++
		args = append(args, *user.Gender)
	}

	if user.Email != nil {
		fields = append(fields, fmt.Sprintf("email=$%d", argID))
		argID++
		args = append(args, *user.Email)
	}

	query := strings.Join(fields, ",")

	_, err := r.db.Exec(ctx, fmt.Sprintf(`
		UPDATE users
		SET %s
		WHERE id=$1`, query),
		args...,
	)
	if err != nil {
		return parseError(err, "updating user")
	}

	return nil
}

func (r *Repository) CheckPhoneUnique(ctx context.Context, phone string) error {
	var v int

	err := r.db.QueryRow(ctx, `
		select 1
		from users
		where phone = $1`, phone,
	).Scan(&v)
	if err != nil {
		return parseError(err, "selecting phone")
	}

	return nil
}

func (r *Repository) CheckEmailUnique(ctx context.Context, email string) error {
	var v int

	err := r.db.QueryRow(ctx, `
		select 1
		from users
		where email = $1`, email,
	).Scan(&v)
	if err != nil {
		return parseError(err, "selecting email")
	}

	return nil
}

func (r *Repository) CreateUser(ctx context.Context, user *domain.User) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO users(name, surname, patronymic, age, gender, image_id, phone, email, role, last_online, created_at, password_enc)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
		user.Name,
		user.Surname,
		user.Patronymic,
		user.Age,
		user.Gender,
		user.ImageID,
		user.Phone,
		user.Email,
		user.Role,
		user.LastOnline,
		user.CreatedAt,
		user.PasswordEncrypted,
	)
	if err != nil {
		return parseError(err, "inserting user")
	}

	return nil
}

func (r *Repository) SetUserImage(ctx context.Context, id int64, imageID string) error {
	_, err := r.db.Exec(ctx, `
		UPDATE users 
		SET image_id = $2
		WHERE id = $1`,
		id,
		imageID,
	)
	if err != nil {
		return fmt.Errorf("updating user: %w", err)
	}

	return nil
}

func (r *Repository) CreatePatient(ctx context.Context, patient *users.Patient) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO patients(name, surname, patronymic, age, gender, image_id, phone, email, last_online, created_at, password_enc)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		patient.Name,
		patient.Surname,
		patient.Patronymic,
		patient.Age,
		patient.Gender,
		patient.ImageID,
		patient.Phone,
		patient.Email,
		patient.LastOnline,
		patient.CreatedAt,
		patient.PasswordEncrypted,
	)
	if err != nil {
		return parseError(err, "inserting patient")
	}

	return nil
}
