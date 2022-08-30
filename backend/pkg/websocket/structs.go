package websocket

import "github.com/gorilla/websocket"

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
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
