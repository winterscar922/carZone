package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDb() {
	connStr := fmt.Sprintf("host=%s user=%s port=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	fmt.Println(connStr)

	pgdb, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error while opening database connection!")
	}

	if pgdb.Ping() != nil {
		log.Fatal("error while connecting to database!")
	}
	fmt.Println("connected to database successfully!")

	db = pgdb
}

func GetDb() *sql.DB {
	return db
}

func CloseDb() {
	if db.Close() != nil {
		log.Fatal("error while closing connection!")
	}
}
