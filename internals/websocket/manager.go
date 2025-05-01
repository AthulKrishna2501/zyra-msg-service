package websocket

import "github.com/gorilla/websocket"

var Clients = make(map[string]*websocket.Conn) 

func AddClient(userID string, conn *websocket.Conn) {
	Clients[userID] = conn
}

func RemoveClient(userID string) {
	delete(Clients, userID)
}

func GetClient(userID string) *websocket.Conn {
	return Clients[userID]
}
