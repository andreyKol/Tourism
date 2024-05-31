package http

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
	"tourism/internal/domain/ws"
)

type WsUseCase interface {
	CreateRoom(ctx context.Context, req *ws.CreateRoomRequest) error
	SyncRooms(ctx context.Context) error
	GetRoomByID(ctx context.Context, roomID string) (*ws.RoomResponse, error)
	GetRooms(ctx context.Context) ([]*ws.RoomResponse, error)
	AddClient(ctx context.Context, roomId string, client *ws.Client) error
	IsClientInRoom(ctx context.Context, roomID, clientID string) (bool, error)
	GetClientsByRoomID(ctx context.Context, roomId string) ([]*ws.ClientResponse, error)
	GetRoomsByClientID(ctx context.Context, clientID string) ([]*ws.RoomResponse, error)
}

type MsgUseCase interface {
	WriteMessage(client *ws.Client)
	ReadMessage(ctx context.Context, client *ws.Client)
	GetMessagesByRoomID(ctx context.Context, roomID string) ([]*ws.MessageResponse, error)
	GetMessagesByClientID(ctx context.Context, clientID string) ([]*ws.MessageResponse, error)
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
		conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, err.Error()))
		conn.Close()
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
		ID:       int64(uuid.New().ID()),
		Content:  "A new client has joined the room",
		RoomID:   roomId,
		ClientID: clientId,
	}

	room, err := h.wsUseCase.GetRoomByID(r.Context(), roomId)
	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte("Error getting room: "+err.Error()))
		conn.Close()
		return
	}

	if room == nil {
		conn.WriteMessage(websocket.TextMessage, []byte("Room not found"))
		conn.Close()
		return
	}

	isInRoom, err := h.wsUseCase.IsClientInRoom(r.Context(), roomId, clientId)
	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte("Error checking if client is in room: "+err.Error()))
		conn.Close()
		return
	}
	if !isInRoom {
		err = h.wsUseCase.AddClient(r.Context(), roomId, cl)
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte("Error joining room: "+err.Error()))
			conn.Close()
			return
		}
	}
	h.hub.Register <- cl
	h.hub.Broadcast <- m

	go h.msgUseCase.WriteMessage(cl)
	h.msgUseCase.ReadMessage(r.Context(), cl)
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
	rooms, err := h.wsUseCase.GetRooms(r.Context())
	if err != nil {
		ErrorResponse(w, r, err)
		return
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

// @Summary      GetMessagesByRoomID
// @Description  Retrieves messages by room ID.
// @Tags         ws
// @Accept       json
// @Produce      json
// @Param        roomId path string true "Room ID"
// @Success      200  {object}  []ws.MessageResponse
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /ws/getMessagesByRoomID/{roomId} [get]
func (h HttpHandler) GetMessagesByRoomID(w http.ResponseWriter, r *http.Request) {
	roomId := chi.URLParam(r, "roomId")
	messages, err := h.msgUseCase.GetMessagesByRoomID(r.Context(), roomId)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, messages)
}

// @Summary      GetMessagesByClientID
// @Description  Retrieves messages by client ID.
// @Tags         ws
// @Accept       json
// @Produce      json
// @Param        clientId query string true "Client ID"
// @Success      200  {object}  []ws.MessageResponse
// @Failure      404  {object}
func (h HttpHandler) GetMessagesByClientID(w http.ResponseWriter, r *http.Request) {
	clientId := r.URL.Query().Get("clientId")
	messages, err := h.msgUseCase.GetMessagesByClientID(r.Context(), clientId)
	if err != nil {
		ErrorResponse(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, messages)
}
