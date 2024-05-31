package usecase

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
	"tourism/internal/domain/ws"
)

type MsgRepository interface {
	SaveMessage(ctx context.Context, msg *ws.Message) error
	GetMessagesByRoomID(ctx context.Context, roomID string) ([]*ws.MessageResponse, error)
	GetMessagesByClientID(ctx context.Context, clientID string) ([]*ws.MessageResponse, error)
}

type MsgUseCase struct {
	messageRepo MsgRepository
	hub         *ws.Hub
}

func NewMsgUseCase(messageRepo MsgRepository, hub *ws.Hub) *MsgUseCase {
	return &MsgUseCase{
		messageRepo: messageRepo,
		hub:         hub,
	}
}

func (uc *MsgUseCase) WriteMessage(client *ws.Client) {
	defer func() {
		client.Conn.Close()
	}()

	for {
		message, ok := <-client.Message
		if !ok {
			return
		}

		client.Conn.WriteJSON(message)
	}
}

func (uc *MsgUseCase) ReadMessage(ctx context.Context, client *ws.Client) {
	defer func() {
		uc.hub.Unregister <- client
		client.Conn.Close()
	}()

	for {
		_, m, err := client.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		msg := &ws.Message{
			Content:   string(m),
			RoomID:    client.RoomID,
			ClientID:  client.ID,
			CreatedAt: time.Now(),
		}

		err = uc.messageRepo.SaveMessage(ctx, msg)
		if err != nil {
			log.Printf("error saving message: %v", err)
		}

		uc.hub.Broadcast <- msg
	}
}

func (uc *MsgUseCase) GetMessagesByRoomID(ctx context.Context, roomID string) ([]*ws.MessageResponse, error) {
	messages, err := uc.messageRepo.GetMessagesByRoomID(ctx, roomID)
	if err != nil {
		return nil, fmt.Errorf("failed to get messages by roomId: %w", err)
	}

	return messages, nil
}

func (uc *MsgUseCase) GetMessagesByClientID(ctx context.Context, clientID string) ([]*ws.MessageResponse, error) {
	messages, err := uc.messageRepo.GetMessagesByRoomID(ctx, clientID)
	if err != nil {
		return nil, fmt.Errorf("failed to get messages by clientID: %w", err)
	}

	return messages, nil
}
