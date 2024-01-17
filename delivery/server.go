package delivery

import (
	"fmt"
	"log"
	"todo-clean-arch/config"
	"todo-clean-arch/delivery/controller"
	"todo-clean-arch/repository"
	"todo-clean-arch/shared/service"
	"todo-clean-arch/usecase"

	"github.com/gin-gonic/gin"
)

type Server struct {
	authorUC usecase.AuthorUsecase
	taskUC   usecase.TaskUsecase
	authUC   usecase.AuthUseCase
	engine   *gin.Engine
	host     string
}

func (s *Server) initRoute() {
	rg := s.engine.Group("/api/v1")
	controller.NewAuthorHandler(s.authorUC, rg).Route()
	controller.NewTaskHandler(s.taskUC, rg).Route()
	controller.NewAuthController(s.authUC, rg).Route()
}

func (s *Server) Run() {
	s.initRoute()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("failed to start server %v", err))
	}
}

func NewServer() *Server {
	db := config.ConnectDB()
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("config error: %v", err))
	}

	taskRepository := repository.NewTaskRepository(db)
	authorRepository := repository.NewAuthorRepository(db)
	authorUseCase := usecase.NewAuthorUseCase(authorRepository)
	jwtService := service.NewJwtService(cfg.TokenConfig)
	authUC := usecase.NewAuthUseCase(authorUseCase, jwtService)
	taskUseCase := usecase.NewTaskUseCase(taskRepository, authorUseCase)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		authorUC: authorUseCase,
		taskUC:   taskUseCase,
		authUC:   authUC,
		engine:   engine,
		host:     host,
	}
}
