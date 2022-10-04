package websocket

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/doxanocap/golang-react/backend/pkg/database"
	"github.com/doxanocap/golang-react/backend/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var currentChatHistory []models.ChatHistory

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Upgrade(ctx *gin.Context) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}

func Sender(ctx *gin.Context) {
	res, err := database.DB.Query("SELECT * FROM messages")
	if err != nil {
		panic(err)
	}
	defer res.Close()
	currentChatHistory = []models.ChatHistory{}
	for res.Next() {
		var current models.ChatHistory
		err = res.Scan(&current.Time, &current.Username, &current.Message)
		if err != nil {
			panic(err)
		}
		currentChatHistory = append(currentChatHistory, current)
	}
	data, _ := json.MarshalIndent(currentChatHistory, "", "\t")
	ctx.JSON(http.StatusOK, string(data))
}

func ListofUsers(ctx *gin.Context) {
	res, err := database.DB.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	defer res.Close()
	var data []models.User
	for res.Next() {
		var Current models.User
		err = res.Scan(&Current.Id, &Current.Token, &Current.Username, &Current.Email, &Current.Password)
		if err != nil {
			panic(err)
		}
		data = append(data, Current)
	}
	dataJSON, _ := json.MarshalIndent(data, "", "\t")
	ctx.JSON(http.StatusOK, string(dataJSON))
}

func ListofOnlineUsers(ctx *gin.Context) {
	res, err := database.DB.Query("SELECT * FROM onlineUsers")
	if err != nil {
		panic(err)
	}
	defer res.Close()
	var data []models.User
	for res.Next() {
		var Current models.User
		err = res.Scan(&Current.Id, &Current.Username, &Current.Email)
		if err != nil {
			panic(err)
		}
		if Current.Username != "" && Current.Email != "" && Current.Id != 0 {
			data = append(data, Current)
		}
	}
	dataJSON, _ := json.MarshalIndent(data, "", "\t")
	ctx.JSON(http.StatusOK, string(dataJSON))
}
