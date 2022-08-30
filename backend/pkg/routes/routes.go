package routes

import (
	"fmt"

	"github.com/doxanocap/golang-react/backend/pkg/controllers"
	"github.com/doxanocap/golang-react/backend/pkg/websocket"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	r := gin.Default()
	pool := websocket.NewPool()

	go websocket.Start(pool)

	r.GET("/api/fetch", websocket.Sender)
	r.GET("/api/websocket", func(ctx *gin.Context) {
		serveWs(pool, ctx)
	})
	r.GET("/api/register", controllers.Register)
	r.Run(":8080")

}

func serveWs(pool *websocket.Pool, ctx *gin.Context) {
	fmt.Println("WebSocket Endpoint Hit")
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
