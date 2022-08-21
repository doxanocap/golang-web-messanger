package websocket

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

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
	Time     string `json: "time"`
	Username string `json: "username"`
	Message  string `json: "message"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}
	return ws, nil
}

func Reader(conn *websocket.Conn) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	ErrorHandler(err)

	for {
		msgType, msg, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}

		insert, err := db.Query(fmt.Sprintf("INSERT INTO messages (time, username, message) VALUES('%s','%s','%s')", string(time.Now().Format("2006-01-02 15:04:05")), "admin", string(msg)))
		ErrorHandler(err)

		if err := conn.WriteMessage(msgType, []byte("msgqweqeqwe")); err != nil {
			log.Println(err)
			defer func(db *sql.DB) {
				_ = db.Close()
			}(db)
			return
		}
		defer insert.Close()
	}
}

func Writer(conn *websocket.Conn) {
	for {
		fmt.Println("Sending")
		messageType, r, err := conn.NextReader()
		if err != nil {
			fmt.Println(err)
			return
		}
		w, err := conn.NextWriter(messageType)
		if err != nil {
			fmt.Println(err)
			return
		}
		if _, err := io.Copy(w, r); err != nil {
			fmt.Println(err)
			return
		}
		if err := w.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}
}

func ErrorHandler(err error) {
	if err != nil {
		panic(err)
	}
}
