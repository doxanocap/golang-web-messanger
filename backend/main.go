package main

import (
	"fmt"

	"github.com/doxanocap/golang-react/backend/pkg/websocket"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func setupRoutes() {
	r := gin.Default()
	pool := websocket.NewPool()

	go websocket.Start(pool)

	r.GET("/put", websocket.Sender)
	r.GET("/ws", func(ctx *gin.Context) {
		serveWs(pool, ctx)
	})

	r.Run(":8080")
}

func serveWs(pool *websocket.Pool, ctx *gin.Context) {
	fmt.Println("WebSocket Endpoint Hit")
	websocket.EnableCors(ctx)
	conn, err := websocket.Upgrade(ctx.Writer, ctx.Request)
	if err != nil {
		fmt.Fprintf(ctx.Writer, "%+V\n", err)
	}
	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func main() {
	setupRoutes()
}
