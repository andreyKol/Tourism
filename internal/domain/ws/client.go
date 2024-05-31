package ws

import (
	"github.com/gorilla/websocket"
	"time"
)

type Client struct {
	ID      string `json:"id"`
	Conn    *websocket.Conn
	Message chan *Message
	RoomID  string `json:"roomId"`
}

type Message struct {
	ID        int64
	Content   string
	RoomID    string
	ClientID  string
	CreatedAt time.Time
}

type MessageResponse struct {
	ID        int64
	Content   string
	RoomID    string
	ClientID  string
	CreatedAt time.Time
}

type ClientResponse struct {
	ID string `json:"id"`
}
