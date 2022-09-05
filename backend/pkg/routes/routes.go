package routes

import (
	"fmt"
	"github.com/doxanocap/golang-react/backend/pkg/controllers"
	"github.com/doxanocap/golang-react/backend/pkg/models"
	"github.com/doxanocap/golang-react/backend/pkg/websocket"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET", "PATCH", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Accept", "Accept-Encoding", "Authorization", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Authorization"},
		AllowCredentials: true,
	}))
	pool := models.NewPool()

	go websocket.Start(pool)

	r.GET("/api/all-users", websocket.ListofUsers)
	r.GET("api/online-users", websocket.ListofOnlineUsers)
	r.GET("/api/fetch", websocket.Sender)
	r.GET("/api/user", controllers.User)
	r.GET("/api/websocket", func(ctx *gin.Context) {
		serveWs(pool, ctx)
	})
	r.POST("/api/register", controllers.Register)
	r.POST("/api/login", controllers.Login)
	r.POST("/api/logout", controllers.Logout)

	r.Run(":8080")

}

func serveWs(pool *models.Pool, ctx *gin.Context) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(ctx, ctx.Request)
	if err != nil {
		fmt.Fprintf(ctx.Writer, "%+V\n", err)
	}
	client := &models.Client{
		Conn: conn,
		Pool: pool,
		Ctx:  ctx,
	}

	pool.Register <- client
	websocket.Read(client)
}
