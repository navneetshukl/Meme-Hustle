package websocket

import (
	"memes-hustle/internals/core/users"
	"net/http"
)

type WebSocketEvent struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type WebSocketUseCase interface {
	UpgradeToWebSocket(w http.ResponseWriter, r *http.Request, user *users.User) error
}
