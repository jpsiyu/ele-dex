package ws

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				client.Unregister()
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				client.Send(message)
			}
		}
	}
}

func (h *Hub) Broadcast(msg []byte) {
	for client := range h.clients {
		client.Send(msg)
	}
}
