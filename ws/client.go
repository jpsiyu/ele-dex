package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	hub  *Hub
	conn *websocket.Conn
}

func NewClient(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{
		hub:  hub,
		conn: conn,
	}

	hub.register <- client
	client.serve()
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) Send(msg []byte) {
	c.conn.WriteMessage(websocket.TextMessage, msg)
}

func (c *Client) serve() {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			c.hub.unregister <- c
			break
		}
		log.Println("rece message:", string(message))
	}
}
