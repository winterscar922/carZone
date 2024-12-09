package driver

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var db *sql.DB

func InitDb() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error while opening database connection!")
	}

	if db.Ping() != nil {
		log.Fatal("error while connecting to database!")
	}
	fmt.Println("connected to database successfully!")
}

func GetDb() *sql.DB {
	return db
}

func CloseDb() {
	if db.Close() != nil {
		log.Fatal("error while closing connection!")
	}
}
