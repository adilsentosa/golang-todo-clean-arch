package repository

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"
	"time"
	"todo-clean-arch/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type AuthorRepositoryTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    AuthorRepository
}

var expectedAuthor = model.Author{
	ID:        "1",
	Name:      "test",
	Email:     "JLx4I@example.com",
	Password:  "test",
	Role:      "admin",
	CreatedAt: time.Now(),
	UpdatedAt: &time.Time{},
	Tasks:     nil,
}

func (s *AuthorRepositoryTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	s.mockDB = db
	s.mockSql = mock
	s.repo = NewAuthorRepository(s.mockDB)
}

func (s *AuthorRepositoryTestSuite) TestGetByEmail_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).
		AddRow(expectedAuthor.ID, expectedAuthor.Name, expectedAuthor.Email, expectedAuthor.Password, expectedAuthor.CreatedAt, expectedAuthor.UpdatedAt)
	s.mockSql.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, password,created_at, upadated_at FROM authors WHERE email=$1")).WithArgs(
		expectedAuthor.Email,
	).WillReturnRows(rows)
	actual, err := s.repo.GetByEmail(expectedAuthor.Email)
	s.Nil(err)
	s.NoError(err)
	s.Equal(expectedAuthor.ID, actual.ID)
	s.Equal(expectedAuthor.Email, actual.Email)
	s.Equal(expectedAuthor.Password, actual.Password)
	s.Equal(expectedAuthor.Name, actual.Name)
}

func (s *AuthorRepositoryTestSuite) TestGetByEmail_Fail() {
	s.mockSql.ExpectQuery(regexp.QuoteMeta("SELECT id,email,password,created_at,updated_at FROM authors WHERE email = $1")).WithArgs("xx@xx.com").WillReturnError(fmt.Errorf("error"))
	actual, err := s.repo.GetByEmail("xx@xx.com")
	s.Error(err)
	s.NotNil(err)
	s.Equal(model.Author{}, actual)
}

func TestAuthorRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AuthorRepositoryTestSuite))
}
