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

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) Send(msg []byte) {
	c.conn.WriteMessage(websocket.TextMessage, msg)
}

func (c *Client) serve() {
	for {
		mt, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			c.hub.unregister <- c
			break
		}
		log.Printf("recv: %s", message)
		err = c.conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func ServeWS(hub *Hub, w http.ResponseWriter, r *http.Request) {
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
