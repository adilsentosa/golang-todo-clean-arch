package delivery

import (
	"fmt"
	"todo-clean-arch/config"
	"todo-clean-arch/delivery/controller"
	"todo-clean-arch/repository"
	"todo-clean-arch/usecase"

	"github.com/gin-gonic/gin"
)

type Server struct {
	authorUC usecase.AuthorUsecase
	taskUC   usecase.TaskUsecase
	engine   *gin.Engine
	host     string
}

func (s *Server) initRoute() {
	rg := s.engine.Group("/api/v1")
	controller.NewAuthorHandler(s.authorUC, rg).Route()
}

func (s *Server) Run() {
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("failed to start server %v", err))
	}
}

func NewServer() *Server {
	db := config.ConnectDB()

	taskRepository := repository.NewTaskRepository(db)
	authorRepository := repository.NewAuthorRepository(db)
	authorUseCase := usecase.NewAuthorUseCase(authorRepository)
	taskUseCase := usecase.NewTaskUseCase(taskRepository, authorUseCase)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", "8080")

	return &Server{
		authorUC: authorUseCase,
		taskUC:   taskUseCase,
		engine:   engine,
		host:     host,
	}
}
