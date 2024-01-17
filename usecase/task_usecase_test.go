package usecase

import (
	"fmt"
	"testing"
	"time"
	"todo-clean-arch/mock/repo_mock"
	"todo-clean-arch/mock/usecase_mock"
	"todo-clean-arch/model"
	sharedmodel "todo-clean-arch/shared/shared_model"

	"github.com/stretchr/testify/suite"
)

var expectedTask = model.Task{
	ID:        "1",
	Title:     "This is test task",
	Content:   "This is test content",
	AuthorID:  expectedAuthor.ID,
	CreatedAt: time.Now(),
	UpdatedAt: &time.Time{},
}

var expectedAuthor = model.Author{
	ID:        "1",
	Name:      "test author",
	Email:     "test@example.com",
	Password:  "test password",
	Role:      "user",
	CreatedAt: time.Now(),
	UpdatedAt: &time.Time{},
	Tasks: []model.Task{
		{
			ID:        "1",
			Title:     "this is test task",
			Content:   "this is test content",
			AuthorID:  "1",
			CreatedAt: time.Now(),
			UpdatedAt: &time.Time{},
		},
	},
}

type TaskUseCaseSuite struct {
	suite.Suite
	trm *repo_mock.TaskRepositoryMock
	aum *usecase_mock.AuthorUseCaseMock
	tuc TaskUsecase
}

func (s *TaskUseCaseSuite) SetupTest() {
	s.trm = new(repo_mock.TaskRepositoryMock)
	s.aum = new(usecase_mock.AuthorUseCaseMock)
	s.tuc = NewTaskUseCase(s.trm, s.aum)
}

func (s *TaskUseCaseSuite) TestRegisterNewTask_Success() {
	s.aum.On("FindAuthorByID", expectedTask.AuthorID).Return(expectedAuthor, nil)
	s.trm.On("Create", expectedTask).Return(expectedTask, nil)
	actuel, err := s.tuc.RegisterNewTask(expectedTask)
	s.NoError(err)
	s.Nil(err)
	s.Equal(expectedTask.Title, actuel.Title)
}

func (s *TaskUseCaseSuite) TestRegisterNewTaskFindAuthorByID_Failed() {
	s.aum.On("FindAuthorByID", expectedTask.AuthorID).Return(model.Author{}, fmt.Errorf("error"))
	_, err := s.tuc.RegisterNewTask(expectedTask)
	s.Error(err)
	s.NotNil(err)
}

func (s *TaskUseCaseSuite) TestRegisterNewTask_EmptyField() {
	s.aum.On("FindAuthorByID", expectedTask.AuthorID).Return(expectedAuthor, nil)
	payloadMock := model.Task{
		Title:    "",
		Content:  "",
		AuthorID: expectedTask.AuthorID,
	}
	_, err := s.tuc.RegisterNewTask(payloadMock)
	s.Error(err)
	s.NotNil(err)
}

func (s *TaskUseCaseSuite) TestRegisterNewTask_Fail() {
	s.aum.On("FindAuthorByID", expectedTask.AuthorID).Return(expectedAuthor, nil)
	s.trm.On("Create", expectedTask).Return(model.Task{}, fmt.Errorf("error"))
	_, err := s.tuc.RegisterNewTask(expectedTask)
	s.Error(err)
	s.NotNil(err)
}

func (s *TaskUseCaseSuite) TestFindAllTask_Success() {
	mockData := []model.Task{expectedTask}
	mockPaging := sharedmodel.Paging{
		Page:        1,
		RowsPerPage: 1,
		TotalRows:   5,
		TotalPages:  1,
	}
	s.trm.On("List", 1, 5).Return(mockData, mockPaging, nil)
	actual, paging, err := s.tuc.FindAllTask(1, 5)
	s.NoError(err)
	s.Nil(err)
	s.Equal(expectedTask.Title, actual[0].Title)
	s.Len(actual, 0)
	s.Equal(mockPaging, paging)
}

func TestTaskUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(TaskUseCaseSuite))
}
