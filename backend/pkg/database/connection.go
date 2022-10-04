package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var psqlInfo = "postgres://zkcontozffaiso:421f8c9ea47124bdd2e5ba023ff74a98f76a87a1e4c9fa7320ee2820dca94872@ec2-34-241-90-235.eu-west-1.compute.amazonaws.com:5432/dd99hq0m6gtap6"

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "eldoseldos"
// 	dbname   = "webchat"
// )

// var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
// 	"password=%s dbname=%s sslmode=disable",
// 	host, port, user, password, dbname)

var DB *sql.DB

func Connect() {
	connection, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	DB = connection
}
