package websocket

import (
	"log"
	"memes-hustle/internals/core/users"
	"net/http"
	"sync"

	ws "memes-hustle/internals/core/websocket"

	"github.com/gorilla/websocket"
)

type WebSocketUseCaseImpl struct {
	clients   map[*websocket.Conn]*users.User
	clientsMu sync.RWMutex
	broadcast chan *ws.WebSocketEvent
}

func NewWebSocketUseCase() ws.WebSocketUseCase {
	return &WebSocketUseCaseImpl{
		clients:   make(map[*websocket.Conn]*users.User),
		broadcast: make(chan *ws.WebSocketEvent),
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Adjust for production (e.g., check allowed origins)
	},
}

func (s *WebSocketUseCaseImpl) UpgradeToWebSocket(w http.ResponseWriter, r *http.Request, user *users.User) error {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	// Register client
	s.clientsMu.Lock()
	s.clients[conn] = user
	s.clientsMu.Unlock()

	// Handle WebSocket connection
	go s.handleWebSocketConnection(conn, user)

	return nil
}

func (s *WebSocketUseCaseImpl) handleWebSocketConnection(conn *websocket.Conn, user *users.User) {
	defer func() {
		s.clientsMu.Lock()
		delete(s.clients, conn)
		s.clientsMu.Unlock()
		conn.Close()
	}()

	// Send welcome message
	event := &ws.WebSocketEvent{
		Type:    "welcome",
		Payload: map[string]string{"message": "Welcome, " + user.Username},
	}
	if err := conn.WriteJSON(event); err != nil {
		log.Printf("Error sending welcome message: %v", err)
		return
	}

	// Listen for messages
	for {
		var event ws.WebSocketEvent
		if err := conn.ReadJSON(&event); err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		// Handle incoming messages (e.g., forward to broadcast channel)
		s.broadcast <- &event
	}
}

func (s *WebSocketUseCaseImpl) handleBroadcast() {
	for event := range s.broadcast {
		s.clientsMu.RLock()
		for conn := range s.clients {
			if err := conn.WriteJSON(event); err != nil {
				log.Printf("Error broadcasting message: %v", err)
			}
		}
		s.clientsMu.RUnlock()
	}
}
