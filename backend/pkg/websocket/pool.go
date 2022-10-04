package websocket

import (
	"fmt"
	"time"

	"github.com/doxanocap/golang-react/backend/pkg/database"
	"github.com/doxanocap/golang-react/backend/pkg/models"
)

func Start(pool *models.Pool) {
	for {
		select {
		case client := <-pool.Register:
			currUser := ShowCurrentUser(client)
			res, err := database.DB.Query(fmt.Sprintf("SELECT * FROM onlineUsers WHERE id = '%d'", currUser.Id))
			if err != nil {
				panic(err)
			}
			i := 0
			for res.Next() {
				i++
				break
			}
			res.Close()
			if i == 0 && currUser.Id != 0 && currUser.Token != "" {
				res1, err1 := database.DB.Query(fmt.Sprintf("INSERT INTO onlineUsers (id, username, email) VALUES('%d','%s','%s')", currUser.Id, currUser.Username, currUser.Email))
				if err1 != nil {
					panic(err)
				}
				res1.Close()
			}
			pool.Clients[client] = true
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			currUser := ShowCurrentUser(client)
			res, err := database.DB.Query(fmt.Sprintf("DELETE FROM onlineUsers WHERE id = '%d'", currUser.Id))
			if err != nil {
				panic(err)
			}
			res.Close()
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			fmt.Println(message.Message)
			res, err := database.DB.Query(fmt.Sprintf("INSERT INTO messages (time, username, message) VALUES('%s','%s','%s')", string(time.Now().Format("02.01.2006, 15:04:05")), string(message.Username), string(message.Message)))
			if err != nil {
				fmt.Println(err)
			}
			res.Close()
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
