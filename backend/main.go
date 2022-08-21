package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/doxanocap/golang-react/backend/pkg/websocket"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "eldoseldos"
	dbname   = "webchat"
)

type chatHistory struct {
	Time     string `json: "time"`
	Username string `json: "username"`
	Message  string `json: "message"`
}

func setupRoutes() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		fmt.Fprintf(ctx.Writer, "Simple server")
	})
	r.GET("/ws", serveWs)
	r.GET("/sendChatToFront", sendChatToFront)
	r.Run(":8080")
}

var currentChatHistory = []chatHistory{}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}
	return ws, nil
}

func sendChatToFront(ctx *gin.Context) {
	data, _ := json.MarshalIndent(currentChatHistory, "", "\t")
	ctx.JSON(http.StatusOK, string(data))
	fmt.Println(string(data))
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func main() {
	setupRoutes()
}

func errorChecker(err error) {
	if err != nil {
		panic(err)
	}
}
