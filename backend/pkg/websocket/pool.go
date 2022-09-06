package websocket

import (
	"fmt"
	"github.com/doxanocap/golang-react/backend/pkg/database"
	"github.com/doxanocap/golang-react/backend/pkg/models"
	"time"
)

func Start(pool *models.Pool) {
	for {
		select {
		case client := <-pool.Register:
			currUser := ShowCurrentUser(client)
			res, _ := database.DB.Query(fmt.Sprintf("SELECT * FROM onlineUsers WHERE id = '%d'", currUser.Id))
			i := 0
			for res.Next() {
				i++
				break
			}
			if i == 0 && currUser.Id != 0 && currUser.Token != "" {
				database.DB.Query(fmt.Sprintf("INSERT INTO onlineUsers (id, username, email) VALUES('%d','%s','%s')", currUser.Id, currUser.Username, currUser.Email))
			}
			pool.Clients[client] = true
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			currUser := ShowCurrentUser(client)
			database.DB.Query(fmt.Sprintf("DELETE FROM onlineUsers WHERE id = '%d'", currUser.Id))
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			fmt.Println(message.Message)
			_, err := database.DB.Query(fmt.Sprintf("INSERT INTO messages (time, username, message) VALUES('%s','%s','%s')", string(time.Now().Format("02.01.2006, 15:04:05")), string(message.Username), string(message.Message)))
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
