package models

type User struct {
	Id       uint   `json:"id"`
	Token    string `json:"token"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
}
