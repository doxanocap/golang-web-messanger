package models

import "github.com/gorilla/websocket"

type User struct {
	Id       uint
	Name     string
	Email    string
	Password string
}

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan ChatHistory
}

type ChatHistory struct {
	Time     string `json:"time"`
	Username string `json:"username"`
	Message  string `json:"message"`
	Type     int    `json:"type"`
}

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}
