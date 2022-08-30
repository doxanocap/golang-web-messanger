package websocket

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/doxanocap/golang-react/backend/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var currentChatHistory []ChatHistory

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}

func EnableCors(ctx *gin.Context) {
	(ctx.Writer).Header().Set("Access-Control-Allow-Origin", "*")
}

func Sender(ctx *gin.Context) {
	EnableCors(ctx)
	res, err := database.DB.Query("SELECT * FROM messages")
	if err != nil {
		panic(err)
	}

	currentChatHistory = []ChatHistory{}
	for res.Next() {
		var current ChatHistory
		err = res.Scan(&current.Time, &current.Username, &current.Message)
		if err != nil {
			panic(err)
		}
		currentChatHistory = append(currentChatHistory, current)
	}
	data, _ := json.MarshalIndent(currentChatHistory, "", "\t")
	ctx.JSON(http.StatusOK, string(data))
}
