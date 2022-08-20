package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "eldoseldos"
	dbname   = "grindset"
)

type Users struct {
	Id       int
	Username string
	Email    string
	Password string
}

func setupRoutes() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		fmt.Fprintf(ctx.Writer, "Simple server")
	})
	r.GET("/ws", serveWs)
	r.Run(":8080")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func reader(conn *websocket.Conn) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)

	res, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	for res.Next() {
		var user Users
		err := res.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
		fmt.Println(user.Id)	
		if err != nil {
			panic(err)
		}
	}

	for {
		// reading messages
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		file, err := os.OpenFile("messages.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

		if err != nil {
			panic(err)
		}
		_, _ = file.WriteString(fmt.Sprint(time.Now().Format("2006-01-02 15:04:05"), " | message:", string(msg), "\n"))
		fmt.Println(time.Now().Format("15:04:05"), "message --->", string(msg))

		if err := conn.WriteMessage(msgType, msg); err != nil {
			log.Println(err)
			return
		}

	}
}

func serveWs(ctx *gin.Context) {
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	reader(ws)
}

func main() {
	setupRoutes()
}
