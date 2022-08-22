package main

import (
	"fmt"

	"github.com/doxanocap/golang-react/backend/pkg/websocket"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func setupRoutes() {
	r := gin.Default()
	r.GET("/ws", serveWs)
	r.GET("/12", websocket.Sender)
	r.Run(":8080")
}

func serveWs(ctx *gin.Context) {
	ws, err := websocket.Upgrade(ctx.Writer, ctx.Request)
	if err != nil {
		fmt.Fprintf(ctx.Writer, "%+V\n", err)
	}
	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func main() {
	fmt.Println("Start:")
	setupRoutes()
}

func errorChecker(err error) {
	if err != nil {
		panic(err)
	}
}
