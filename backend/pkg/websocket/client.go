package websocket

import (
	"fmt"
	"log"
	"time"
)

func (c *Client) Read() {
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
		message := ChatHistory{Type: messageType, Time: string(time.Now().Format("02.01.2006, 15:04:05")), Username: "Doxa", Message: string(p)}
		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}

// if err := c.Conn.WriteMessage(websocket.TextMessage, []byte("qwe")); err != nil {
// 	fmt.Println("Can't send")
// } else {
// 	fmt.Println("Sent message")
// }
