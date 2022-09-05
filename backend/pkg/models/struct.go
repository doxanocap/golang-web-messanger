package models

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
	Ctx  *gin.Context
}

type ChatHistory struct {
	Time     string `json:"time"`
	Username string `json:"username"`
	Message  string `json:"message"`
	Type     int    `json:"type"`
}

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan ChatHistory
}

type User struct {
	Id       uint   `json:"id"`
	Token    string `json:"token"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan ChatHistory),
	}
}
