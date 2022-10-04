package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/doxanocap/golang-react/backend/pkg/controllers"
	"github.com/doxanocap/golang-react/backend/pkg/models"
	"github.com/doxanocap/golang-react/backend/pkg/websocket"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET", "PATCH", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Accept", "Accept-Encoding", "Authorization", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Authorization"},
		AllowCredentials: true,
	}))
	pool := models.NewPool()
	r.Use(static.Serve("/", static.LocalFile("./web", true)))
	r.Use(static.Serve("/webchat", static.LocalFile("./web", true)))
	r.Use(static.Serve("/login", static.LocalFile("./web", true)))
	r.Use(static.Serve("/register", static.LocalFile("./web", true)))

	go websocket.Start(pool)
	api := r.Group("/api")
	api.GET("/all-users", websocket.ListofUsers)
	api.GET("/online-users", websocket.ListofOnlineUsers)
	api.GET("/fetch", websocket.Sender)
	api.GET("/user", controllers.User)
	api.GET("/logout", controllers.Logout)
	api.GET("/websocket", func(ctx *gin.Context) {
		serveWs(pool, ctx)
	})
	api.GET("/check", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "Work pls!!"}) })
	api.POST("/register", controllers.Register)
	api.POST("/login", controllers.Login)
	r.Run(":" + port)

}

func serveWs(pool *models.Pool, ctx *gin.Context) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(ctx)
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
