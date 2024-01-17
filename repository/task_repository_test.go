package repository

import (
	"database/sql"
	"testing"
	"time"
	"todo-clean-arch/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type TaskRepositoryTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    TaskRepository
}

var expectedTask = model.Task{
	ID:        "1",
	Title:     "title",
	Content:   "content",
	AuthorID:  "1",
	CreatedAt: time.Now(),
	UpdatedAt: &time.Time{},
}

func (s *TaskRepositoryTestSuite) Setuptask() {
	db, mock, _ := sqlmock.New()
	s.mockDB = db
	s.mockSql = mock
	s.repo = NewTaskRepository(db)
}

func (s *TaskRepositoryTestSuite) TestCreate_Success() {
	s.mockSql.ExpectQuery("INSERT INTO tasks").WithArgs(
		expectedTask.Title,
		expectedTask.Content,
		expectedTask.AuthorID,
	).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at"}).AddRow(expectedTask.ID, expectedTask.CreatedAt))
	actual, err := s.repo.Create(expectedTask)
	s.Nil(err)
	s.Equal(expectedTask.Title, actual.Title)
}

func TestTaskRepositorySuite(t *testing.T) {
	suite.Run(t, new(TaskRepositoryTestSuite))
}
