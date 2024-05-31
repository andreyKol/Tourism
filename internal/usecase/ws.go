package usecase

import (
	"context"
	"fmt"
	"tourism/internal/domain/ws"
)

type WsRepository interface {
	CreateRoom(ctx context.Context, room *ws.Room) error
	GetRoomByID(ctx context.Context, roomID string) (*ws.RoomResponse, error)
	GetRooms(ctx context.Context) ([]*ws.RoomResponse, error)
	AddClient(ctx context.Context, client *ws.Client) error
	IsClientInRoom(ctx context.Context, roomID, clientID string) (bool, error)
	GetClientsByRoomID(ctx context.Context, roomID string) ([]*ws.ClientResponse, error)
	GetRoomsByClientID(ctx context.Context, clientID string) ([]*ws.RoomResponse, error)
}

type WsUseCase struct {
	wsRepo WsRepository
	hub    *ws.Hub
}

func NewWsUseCase(wsRepo WsRepository, hub *ws.Hub) *WsUseCase {
	return &WsUseCase{
		wsRepo: wsRepo,
		hub:    hub,
	}
}

func (uc *WsUseCase) SyncRooms(ctx context.Context) error {
	rooms, err := uc.wsRepo.GetRooms(ctx)
	if err != nil {
		return err
	}

	for _, room := range rooms {
		uc.hub.Rooms[room.ID] = &ws.Room{
			ID:      room.ID,
			Clients: make(map[string]*ws.Client),
		}
	}
	return nil
}

func (uc *WsUseCase) CreateRoom(ctx context.Context, req *ws.CreateRoomRequest) error {
	room := &ws.Room{
		ID:      req.ID,
		Clients: make(map[string]*ws.Client),
	}

	err := uc.wsRepo.CreateRoom(ctx, room)
	if err != nil {
		return fmt.Errorf("failed to save room to the database: %w", err)
	}

	return nil
}

func (uc *WsUseCase) GetRoomByID(ctx context.Context, roomID string) (*ws.RoomResponse, error) {
	room, err := uc.wsRepo.GetRoomByID(ctx, roomID)
	if err != nil {
		return nil, fmt.Errorf("failed to get rooms: %w", err)
	}

	return room, nil
}

func (uc *WsUseCase) GetRooms(ctx context.Context) ([]*ws.RoomResponse, error) {
	rooms, err := uc.wsRepo.GetRooms(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get rooms: %w", err)
	}

	return rooms, nil
}

func (uc *WsUseCase) AddClient(ctx context.Context, roomId string, client *ws.Client) error {
	client.RoomID = roomId
	err := uc.wsRepo.AddClient(ctx, client)
	if err != nil {
		return fmt.Errorf("failed to add client to the database: %w", err)
	}

	return nil
}

func (uc *WsUseCase) IsClientInRoom(ctx context.Context, roomID, clientID string) (bool, error) {
	isInRoom, err := uc.wsRepo.IsClientInRoom(ctx, roomID, clientID)
	if err != nil {
		return false, fmt.Errorf("failed to check if client is in room: %w", err)
	}

	return isInRoom, nil
}

func (uc *WsUseCase) GetClientsByRoomID(ctx context.Context, roomId string) ([]*ws.ClientResponse, error) {
	clients, err := uc.wsRepo.GetClientsByRoomID(ctx, roomId)
	if err != nil {
		return nil, fmt.Errorf("failed to get clients by roomId: %w", err)
	}

	return clients, nil
}

func (uc *WsUseCase) GetRoomsByClientID(ctx context.Context, clientID string) ([]*ws.RoomResponse, error) {
	rooms, err := uc.wsRepo.GetRoomsByClientID(ctx, clientID)
	if err != nil {
		return nil, fmt.Errorf("failed to get rooms by client ID: %w", err)
	}

	return rooms, nil
}
