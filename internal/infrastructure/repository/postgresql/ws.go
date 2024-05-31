package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"tourism/internal/domain/ws"
)

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

func (r *Repository) GetRoomByID(ctx context.Context, roomID string) (*ws.RoomResponse, error) {
	row := r.db.QueryRow(ctx, `
        SELECT id
        FROM rooms
        WHERE id = $1
    `, roomID)

	var room ws.RoomResponse
	if err := row.Scan(&room.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("room not found")
		}
		return nil, parseError(err, "scanning room row")
	}

	return &room, nil
}

func (r *Repository) GetRooms(ctx context.Context) ([]*ws.RoomResponse, error) {
	rows, err := r.db.Query(ctx, `
        SELECT id
        FROM rooms
    `)
	if err != nil {
		return nil, parseError(err, "selecting rooms")
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

func (r *Repository) AddClient(ctx context.Context, client *ws.Client) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO clients (client_id, room_id)
		VALUES ($1, $2)`,
		client.ID, client.RoomID,
	)
	if err != nil {
		return parseError(err, "inserting client")
	}

	return nil
}

func (r *Repository) IsClientInRoom(ctx context.Context, roomID, clientID string) (bool, error) {
	row := r.db.QueryRow(ctx, `
        SELECT EXISTS (
            SELECT 1
            FROM clients
            WHERE room_id = $1 AND client_id = $2
        )`, roomID, clientID)

	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, parseError(err, "checking if client is in room")
	}

	return exists, nil
}

func (r *Repository) GetClientsByRoomID(ctx context.Context, roomID string) ([]*ws.ClientResponse, error) {
	rows, err := r.db.Query(ctx, `
        SELECT client_id
        FROM clients
        WHERE room_id = $1
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
        SELECT room_id
        FROM clients
        WHERE client_id = $1
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
