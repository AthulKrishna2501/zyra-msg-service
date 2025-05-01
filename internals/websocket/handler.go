package websocket

import (
	"context"
	"net/http"
	"time"

	"github.com/AthulKrishna2501/zyra-msg-service/internals/config"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func WebSocketHandler(c *gin.Context) {
	userID := c.Query("userId")

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to upgrade connection"})
		return
	}
	AddClient(userID, conn)
	defer RemoveClient(userID)

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			break
		}

		msg.Timestamp = time.Now()

		_, _ = config.ChatCollection.InsertOne(context.Background(), msg)

		receiverConn := GetClient(msg.ReceiverID)
		if receiverConn != nil {
			_ = receiverConn.WriteJSON(msg)
		}
	}
}
