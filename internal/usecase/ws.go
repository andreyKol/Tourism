package usecase

import (
	"context"
	"fmt"
	"tourism/internal/domain"
	"tourism/internal/domain/ws"
)

//go:generate mockgen -source=ws.go -destination=./mocks/ws.go -package=mocks

type WsRepository interface {
	GetUserByID(ctx context.Context, id int64) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	CreateRoom(ctx context.Context, room *ws.Room) error
	AddClient(ctx context.Context, client *ws.Client) error
	GetClientsByRoomID(ctx context.Context, roomID string) ([]*ws.ClientResponse, error)
	GetRoomsByClientID(ctx context.Context, clientID string) ([]*ws.RoomResponse, error)
}

type WsUseCase struct {
	wsRepo WsRepository
}

func NewWsUseCase(wsRepo WsRepository) *WsUseCase {
	return &WsUseCase{
		wsRepo: wsRepo,
	}
}

func (uc *WsUseCase) GetUserByID(ctx context.Context, id int64) (*domain.User, error) {
	user, err := uc.wsRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *WsUseCase) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := uc.wsRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
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

func (uc *WsUseCase) AddClient(ctx context.Context, roomId string, client *ws.Client) error {
	client.RoomID = roomId

	err := uc.wsRepo.AddClient(ctx, client)
	if err != nil {
		return fmt.Errorf("failed to add client to the database: %w", err)
	}

	return nil
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
