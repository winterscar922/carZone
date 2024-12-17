package driver

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDb() {
	connStr := "user=postgres port=5432 dbname=postgres password=winterscar sslmode=disable"

	fmt.Println(connStr)

	temp, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error while opening database connection!")
	}

	if temp.Ping() != nil {
		log.Fatal("error while connecting to database!")
	}
	fmt.Println("connected to database successfully!")

	db = temp
}

func GetDb() *sql.DB {
	return db
}

func CloseDb() {
	if db.Close() != nil {
		log.Fatal("error while closing connection!")
	}
}
