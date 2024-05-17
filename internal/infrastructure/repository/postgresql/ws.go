package postgresql

import (
	"Tourism/internal/domain"
	"Tourism/internal/domain/ws"
	"context"
)

func (r *Repository) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	var user domain.User

	err := r.db.QueryRow(ctx, `
		select id,
		       name,
		       surname,
		       patronymic
		from users
		where id = $1`, id).Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Patronymic,
	)
	if err != nil {
		return nil, parseError(err, "selecting user")
	}

	return &user, nil
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User

	err := r.db.QueryRow(ctx, `
		select id,
		       name,
		       surname,
		       patronymic
		from users
		where email = $1`, email).Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Patronymic,
	)
	if err != nil {
		return nil, parseError(err, "selecting user")
	}

	return &user, nil
}

func (r *Repository) CreateRoom(ctx context.Context, room *ws.Room) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO rooms (id)
		VALUES ($1)`,
		room.ID,
	)
	if err != nil {
		return parseError(err, "inserting room")
	}

	return nil
}

func (r *Repository) AddClient(ctx context.Context, client *ws.Client) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return parseError(err, "starting transaction")
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
			return
		}
		err = tx.Commit(ctx)
		if err != nil {
			err = parseError(err, "committing transaction")
		}
	}()

	_, err = tx.Exec(ctx, `
		INSERT INTO clients (client_id, room_id)
		VALUES ($1, $2)`,
		client.ID, client.RoomID,
	)
	if err != nil {
		return parseError(err, "inserting client")
	}

	_, err = tx.Exec(ctx, `
		INSERT INTO room_clients (room_id, client_id)
		VALUES ($1, $2)`,
		client.RoomID, client.ID,
	)
	if err != nil {
		return parseError(err, "inserting room_client")
	}

	return nil
}

func (r *Repository) GetClientsByRoomID(ctx context.Context, roomID string) ([]*ws.ClientResponse, error) {
	rows, err := r.db.Query(ctx, `
        SELECT c.id
        FROM clients c
        WHERE c.room_id = $1
    `, roomID)
	if err != nil {
		return nil, parseError(err, "selecting clients by room ID")
	}
	defer rows.Close()

	var clients []*ws.ClientResponse

	for rows.Next() {
		var client ws.ClientResponse
		if err = rows.Scan(&client.ID); err != nil {
			return nil, parseError(err, "scanning client row")
		}
		clients = append(clients, &client)
	}

	if err = rows.Err(); err != nil {
		return nil, parseError(err, "iterating over client rows")
	}

	return clients, nil
}

func (r *Repository) GetRoomsByClientID(ctx context.Context, clientID string) ([]*ws.RoomResponse, error) {
	rows, err := r.db.Query(ctx, `
        SELECT r.id
        FROM rooms r
        INNER JOIN room_clients rc ON r.id = rc.room_id
        WHERE rc.client_id = $1
    `, clientID)
	if err != nil {
		return nil, parseError(err, "selecting rooms by client ID")
	}
	defer rows.Close()

	var rooms []*ws.RoomResponse

	for rows.Next() {
		var room ws.RoomResponse
		if err = rows.Scan(&room.ID); err != nil {
			return nil, parseError(err, "scanning room row")
		}
		rooms = append(rooms, &room)
	}
	if err = rows.Err(); err != nil {
		return nil, parseError(err, "iterating over room rows")
	}

	return rooms, nil
}
