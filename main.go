package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	Connect()
	res, err := DB.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	arr := []User{}
	for res.Next() {
		var user User
		err := res.Scan(&user.Id, &user.Token, &user.Username, &user.Email, &user.Password)
		if err != nil {
			panic(err)
		}
		fmt.Println(user)
		arr = append(arr, user)
	}
}

type User struct {
	Id       uint   `json:"id"`
	Token    string `json:"token"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password []byte `json:"-"`
}

var psqlInfo = "postgres://zkcontozffaiso:421f8c9ea47124bdd2e5ba023ff74a98f76a87a1e4c9fa7320ee2820dca94872@ec2-34-241-90-235.eu-west-1.compute.amazonaws.com:5432/dd99hq0m6gtap6"

var DB *sql.DB

func Connect() {
	connection, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	DB = connection
}
