package config

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"         // isi dengan user klean
	password = "trancendsdb@3821" // isi dengan password user klean
	dbName   = "todo_app_db"
)

func ConnectDB() *sql.DB {
	fmt.Println("Welcome to the Todo APP")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Database")
	return db
}
