package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
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

	taskRepository := repository.NewTaskRepository(db)
	authorRepository := repository.NewAuthorRepository(db)
	// task := model.Task{
	// 	Title:    "Task 1",
	// 	Content:  "Task 1 content",
	// 	AuthorID: "e1bac6e0-8f96-42f9-b3c3-e5040c4b70bf",
	// }

	// task, err := taskRepository.Create(task)
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println(task)
	// }
	author, err := authorRepository.Get("40af5ee7-5a6f-4193-9d6f-66e9d2cf12fb")
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.Repeat("=", 10), "Get Author By AuthorID", strings.Repeat("=", 10))
	printAuthor(author)
	fmt.Println()

	author, err = authorRepository.GetByEmail("naruto@konoha.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.Repeat("=", 10), "Get Author By Email", strings.Repeat("=", 10))
	printAuthor(author)
	fmt.Println()

	tasks, err := taskRepository.GetByAuthorID("40af5ee7-5a6f-4193-9d6f-66e9d2cf12fb")
	if err != nil {
		panic(err)
	}
	fmt.Println(strings.Repeat("=", 10), "GetTask By AuthorID", strings.Repeat("=", 10))
	printTask(tasks)

	tasks, err = taskRepository.List()
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println(strings.Repeat("=", 10), "List all Tasks", strings.Repeat("=", 10))
	printTask(tasks)
}

func printAuthor(author model.Author) {
	fmt.Println()
	fmt.Println("ID:", author.ID)
	fmt.Println("Name:", author.Name)
	fmt.Println("Email:", author.Email)
	fmt.Println("Created_At:", author.CreatedAt)
}

func printTask(tasks []model.Task) {
	for _, task := range tasks {
		fmt.Println()
		fmt.Println("ID:", task.ID)
		fmt.Println("Title:", task.Title)
		fmt.Println("Content:", task.Content)
		fmt.Println("AuthorID:", task.AuthorID)
		fmt.Println("Created_At:", task.CreatedAt)
	}
}
