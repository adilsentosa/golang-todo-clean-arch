package main

import (
	"database/sql"
	"fmt"
	"log"
	"todo-clean-arch/model"
	"todo-clean-arch/repository"

	_ "github.com/lib/pq"
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

func main() {
	db := ConnectDB()

	task := model.Task{
		Title:    "Task 1",
		Content:  "Task 1 content",
		AuthorID: "e1bac6e0-8f96-42f9-b3c3-e5040c4b70bf",
	}
	taskRepository := repository.NewTaskRepository(db)

	task, err := taskRepository.Create(task)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(task)
	}
}
