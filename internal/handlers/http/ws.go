package http

import (
	"Tourism/internal/domain"
	"Tourism/internal/domain/ws"
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/gorilla/websocket"
	"net/http"
)

type WsUseCase interface {
	GetUserByID(ctx context.Context, id int64) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	CreateRoom(ctx context.Context, req *ws.CreateRoomRequest) error
	AddClient(ctx context.Context, roomId string, client *ws.Client) error
	GetClientsByRoomID(ctx context.Context, roomId string) ([]*ws.ClientResponse, error)
	GetRoomsByClientID(ctx context.Context, clientID string) ([]*ws.RoomResponse, error)
}

// @Summary      CreateRoom
// @Description  Creates a new chat room.
// @Tags         ws
// @Accept       json
// @Produce      json
// @Param        request body ws.CreateRoomRequest true "Request body"
// @Success      200
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /ws/createRoom [post]
func (h HttpHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var req ws.CreateRoomRequest
	err := render.DecodeJSON(r.Body, &req)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}
	h.hub.Rooms[req.ID] = &ws.Room{
		ID:      req.ID,
		Clients: make(map[string]*ws.Client),
	}

	err = h.wsUseCase.CreateRoom(r.Context(), &req)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// @Summary      JoinRoom
// @Description  Joins a client to a room.
// @Tags         ws
// @Accept       json
// @Produce      json
// @Param        roomId path string true "Room Id"
// @Success      200
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /ws/joinRoom/{roomId} [get]
func (h HttpHandler) JoinRoom(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	roomId := chi.URLParam(r, "roomId")
	clientId := r.URL.Query().Get("clientId")

	cl := &ws.Client{
		Conn:    conn,
		Message: make(chan *ws.Message, 10),
		ID:      clientId,
		RoomID:  roomId,
	}
	m := &ws.Message{
		Content:  "A new client has joined the room",
		RoomID:   roomId,
		ClientID: clientId,
	}

	err = h.wsUseCase.AddClient(r.Context(), roomId, cl)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	h.hub.Register <- cl
	h.hub.Broadcast <- m

	go cl.WriteMessage()
	cl.ReadMessage(h.hub)
}

// @Summary      GetRooms
// @Tags         ws
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /ws/getRooms [get]
func (h HttpHandler) GetRooms(w http.ResponseWriter, r *http.Request) {
	rooms := make([]ws.RoomResponse, 0)

	for _, r := range h.hub.Rooms {
		rooms = append(rooms, ws.RoomResponse{
			ID: r.ID,
		})
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, rooms)
}

// @Summary      GetClientsByRoomID
// @Tags         Ws
// @Accept       json
// @Produce      json
// @Param        roomId path string true "Room ID"
// @Success      200
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /ws/getClients/ [get]
func (h HttpHandler) GetClientsByRoomID(w http.ResponseWriter, r *http.Request) {
	roomId := chi.URLParam(r, "roomId")
	clients, err := h.wsUseCase.GetClientsByRoomID(r.Context(), roomId)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, clients)
}

// @Summary      GetRoomsByClientID
// @Description  Retrieves the rooms in which the client is currently joined.
// @Tags         ws
// @Accept       json
// @Produce      json
// @Param        clientId query string true "Client ID"
// @Success      200  {object}  []RoomResponse
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /ws/getRooms/ [get]
func (h HttpHandler) GetRoomsByClientID(w http.ResponseWriter, r *http.Request) {
	clientId := r.URL.Query().Get("clientId")
	rooms, err := h.wsUseCase.GetRoomsByClientID(r.Context(), clientId)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, rooms)
}
