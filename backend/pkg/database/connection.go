package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "eldoseldos"
	dbname   = "webchat"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

var DB *sql.DB

func Connect() {
	connection, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	DB = connection
}
