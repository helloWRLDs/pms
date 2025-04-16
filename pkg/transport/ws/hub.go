package ws

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type Hub struct {
	clients map[*websocket.Conn]bool
	mu      sync.Mutex

	Cache map[string]interface{}
}

func (wh *Hub) GetCache() map[string]interface{} {
	wh.mu.Lock()
	defer wh.mu.Unlock()

	return wh.Cache
}

func (wh *Hub) SetCache(id string, val ...interface{}) {
	wh.mu.Lock()
	defer wh.mu.Unlock()

	if wh.Cache == nil {
		wh.Cache = make(map[string]interface{})
	}
	wh.Cache[id] = val
}

func (wh *Hub) Clean() {
	wh.mu.Lock()
	defer wh.mu.Unlock()

	wh.Cache = nil
}

func (wh *Hub) GetClients() map[*websocket.Conn]bool {
	wh.mu.Lock()
	defer wh.mu.Unlock()

	clientsCopy := make(map[*websocket.Conn]bool)
	for conn, val := range wh.clients {
		clientsCopy[conn] = val
	}
	return clientsCopy
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (wh *Hub) AddClient(conn *websocket.Conn) {
	wh.mu.Lock()
	defer wh.mu.Unlock()

	wh.clients[conn] = true
}

func (wh *Hub) RemoveClient(c *websocket.Conn) {
	wh.mu.Lock()
	defer wh.mu.Unlock()

	for conn := range wh.clients {
		if conn == c {
			delete(wh.clients, conn)
			break
		}
	}
}

func (wh *Hub) Broadcast(message []byte) {
	wh.mu.Lock()
	defer wh.mu.Unlock()

	for conn := range wh.clients {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			conn.Close()
			delete(wh.clients, conn)
		}
	}
}
