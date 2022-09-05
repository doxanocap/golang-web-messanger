package websocket

import (
	"github.com/doxanocap/golang-react/backend/pkg/models"
	"log"
	"time"
)

func Read(c *models.Client) {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		user := ShowCurrentUser(c)
		message := models.ChatHistory{Type: messageType, Time: string(time.Now().Format("02.01.2006, 15:04:05")), Username: user.Username, Message: string(p)}
		c.Pool.Broadcast <- message
	}
}
