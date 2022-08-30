package websocket

import (
	"fmt"
	"time"

	"github.com/doxanocap/golang-react/backend/pkg/database"
)

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan ChatHistory),
	}
}

func Start(pool *Pool) {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client := range pool.Clients {
				fmt.Println(client)
				//client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
			}
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client := range pool.Clients {
				fmt.Println(client)
				//	client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			}
		case message := <-pool.Broadcast:

			fmt.Println("Sending message to all clients in Pool")
			fmt.Println(message.Message)
			_, err := database.DB.Query(fmt.Sprintf("INSERT INTO messages (time, username, message) VALUES('%s','%s','%s')", string(time.Now().Format("02.01.2006, 15:04:05")), "Doxa", string(message.Message)))
			if err != nil {
				fmt.Println(err)
			}
			for client := range pool.Clients {
				fmt.Println(client)
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
