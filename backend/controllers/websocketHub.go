package controllers

import ()

type Hub struct {
	clients map[*client]bool
	broadcast chan []byte
	register chan *client
	unregister chan *client

	content []byte
}

var h = Hub{
	broadcast:   make(chan []byte),
	register:    make(chan *client),
	unregister:  make(chan *client),
	clients: 	 make(map[*client]bool),
}

func InitHub() *Hub{
	return &h
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.register:
			h.clients[c] = true
			c.send <- h.content
			break

		case c := <-h.unregister:
			_, ok := h.clients[c]
			if ok {
				delete(h.clients, c)
				close(c.send)
			}
			break

		case m := <-h.broadcast:
			h.content = m
			h.broadcastMessage()
			break
		}
	}
}

func (h *Hub) broadcastMessage() {
	for c := range h.clients {
		select {
		case c.send <- h.content:
			break

		// We can't reach the client
		default:
			close(c.send)
			delete(h.clients, c)
		}
	}
}
