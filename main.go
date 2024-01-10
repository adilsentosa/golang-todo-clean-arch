package main

import (
	"fmt"
	"todo-clean-arch/config"
	"todo-clean-arch/handler"
	"todo-clean-arch/repository"
	"todo-clean-arch/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db := config.ConnectDB()

	taskRepository := repository.NewTaskRepository(db)
	authorRepository := repository.NewAuthorRepository(db)
	authorUseCase := usecase.NewAuthorUseCase(authorRepository)
	taskUseCase := usecase.NewTaskUseCase(taskRepository, authorUseCase)

	taskHandler := handler.NewTaskHandler(taskUseCase)
	authorHanlder := handler.NewAuthorHandler(authorUseCase)

	route := gin.Default()

	tasks := route.Group("/api/v1/tasks")
	{
		tasks.GET("/:id", taskHandler.FindTaskByAuthor)
		tasks.POST("/create", taskHandler.CreateHandler)
	}

	authors := route.Group("/api/v1/authors")
	{
		authors.GET("/list", authorHanlder.ListAuthor)
		authors.GET("/:id", authorHanlder.GetAuthor)
	}

	if err := route.Run(":8080"); err != nil {
		panic(fmt.Errorf("failed to start server: %v", err))
	}
}
