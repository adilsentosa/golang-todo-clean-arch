package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"todo-clean-arch/mock/middleware_mock"
	"todo-clean-arch/mock/usecase_mock"
	"todo-clean-arch/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type AuthorControllerTestSuite struct {
	suite.Suite
	rg  *gin.RouterGroup
	aum *usecase_mock.AuthorUseCaseMock
	amm *middleware_mock.AuthorMiddlewareMock
}

func (s *AuthorControllerTestSuite) SetupTest() {
	s.aum = new(usecase_mock.AuthorUseCaseMock)
	s.amm = new(middleware_mock.AuthorMiddlewareMock)
	r := gin.Default()
	gin.SetMode(gin.TestMode)
	rg := r.Group("/api/v1")
	rg.Use(s.amm.RequireToken("admin"))
	s.rg = rg
}

func (s *AuthorControllerTestSuite) TestListHandler_Success() {
	mockAuthor := []model.Author{
		{
			ID:        "1",
			Name:      "Name",
			Email:     "name@mail.com",
			Role:      "admin",
			CreatedAt: time.Now(),
			UpdatedAt: &time.Time{},
			Tasks: []model.Task{
				{
					ID:        "1",
					Title:     "title",
					Content:   "content",
					CreatedAt: time.Now(),
					UpdatedAt: &time.Time{},
				},
			},
		},
	}
	s.aum.On("FindAllAuthor", mockAuthor[0]).Return(mockAuthor, nil)
	authorController := NewAuthorHandler(s.aum, s.rg, s.amm)
	authorController.Route()
	req, err := http.NewRequest("GET", "/api/v1/authors/list", nil)
	s.NoError(err)
	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	ctx.Set("author", mockAuthor[0].ID)
	s.Equal(http.StatusOK, record.Code)
}

func (s *AuthorControllerTestSuite) TestListHandler_Fail() {
	s.aum.On("FindAllAuthor", "2").Return([]model.Author{}, fmt.Errorf("error"))
	authorController := NewAuthorHandler(s.aum, s.rg, s.amm)
	authorController.Route()
	req, err := http.NewRequest("GET", "/api/v1/authors/list", nil)
	s.NoError(err)
	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = req
	ctx.Set("author", "2")
	authorController.ListAuthor(ctx)
	s.Equal(http.StatusInternalServerError, record.Code)
}

func TestAuthorControllerTestSuite(t *testing.T) {
	suite.Run(t, new(AuthorControllerTestSuite))
}
