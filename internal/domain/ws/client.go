package ws

import (
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	ID      string `json:"id"`
	Conn    *websocket.Conn
	Message chan *Message
	RoomID  string `json:"roomId"`
}

type Message struct {
	Content  string `json:"content"`
	RoomID   string `json:"roomId"`
	ClientID string `json:"clientId"`
}

type ClientResponse struct {
	ID string `json:"id"`
}

func (c *Client) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) ReadMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		msg := &Message{
			Content:  string(m),
			RoomID:   c.RoomID,
			ClientID: c.ID,
		}

		hub.Broadcast <- msg
	}
}
