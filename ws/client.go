package ws

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	user *User
}

func NewClient(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	user := &User{}
	client := &Client{
		hub:  hub,
		conn: conn,
	}
	client.user = user
	client.user.SetClient(client)

	client.Register()
}

func (c *Client) Register() {
	c.hub.register <- c
	c.user.Register()
	c.serve()
}

func (c *Client) Unregister() {
	c.user.Unregister()
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
		c.user.ReceMsg(message)
	}
}
