package websocket

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "eldoseldos"
	dbname   = "webchat"
)

type chatHistory struct {
	Time     string
	Username string
	Message  string
}

var currentChatHistory = []chatHistory{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

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
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	res, err := db.Query("SELECT * FROM messages")
	if err != nil {
		panic(err)
	}
	currentChatHistory = []chatHistory{}
	for res.Next() {
		var current chatHistory
		err = res.Scan(&current.Time, &current.Username, &current.Message)
		if err != nil {
			panic(err)
		}
		currentChatHistory = append(currentChatHistory, current)
	}
	data, _ := json.MarshalIndent(currentChatHistory, "", "\t")
	ctx.JSON(http.StatusOK, string(data))
}
