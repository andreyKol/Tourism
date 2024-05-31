package postgresql

import (
	"context"
	"tourism/internal/domain/ws"
)

func (r *Repository) SaveMessage(ctx context.Context, msg *ws.Message) error {
	query := `INSERT INTO messages (content, room_id, client_id, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.db.QueryRow(ctx, query, msg.Content, msg.RoomID, msg.ClientID, msg.CreatedAt).Scan(&msg.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetMessagesByRoomID(ctx context.Context, roomID string) ([]*ws.MessageResponse, error) {
	rows, err := r.db.Query(ctx, `
        SELECT id, content, room_id, client_id, created_at
        FROM messages
        WHERE room_id = $1
    `, roomID)
	if err != nil {
		return nil, parseError(err, "selecting messages by room ID")
	}
	defer rows.Close()

	var messages []*ws.MessageResponse

	for rows.Next() {
		var msg ws.MessageResponse
		if err = rows.Scan(&msg.ID, &msg.Content, &msg.RoomID, &msg.ClientID, &msg.CreatedAt); err != nil {
			return nil, parseError(err, "scanning message row")
		}
		messages = append(messages, &msg)
	}

	if err = rows.Err(); err != nil {
		return nil, parseError(err, "iterating over message rows")
	}

	return messages, nil
}

func (r *Repository) GetMessagesByClientID(ctx context.Context, clientID string) ([]*ws.MessageResponse, error) {
	rows, err := r.db.Query(ctx, `
        SELECT id, content, room_id, client_id, created_at
        FROM messages
        WHERE client_id = $1
    `, clientID)
	if err != nil {
		return nil, parseError(err, "selecting messages by client ID")
	}
	defer rows.Close()

	var messages []*ws.MessageResponse

	for rows.Next() {
		var msg ws.MessageResponse
		if err = rows.Scan(&msg.ID, &msg.Content, &msg.RoomID, &msg.ClientID, &msg.CreatedAt); err != nil {
			return nil, parseError(err, "scanning message row")
		}
		messages = append(messages, &msg)
	}

	if err = rows.Err(); err != nil {
		return nil, parseError(err, "iterating over message rows")
	}

	return messages, nil
}
