package repo_mock

import (
	"todo-clean-arch/model"
	sharedmodel "todo-clean-arch/shared/shared_model"

	"github.com/stretchr/testify/mock"
)

type TaskRepositoryMock struct {
	mock.Mock
}

func (m *TaskRepositoryMock) List(page, size int) ([]model.Task, sharedmodel.Paging, error) {
	args := m.Called(page, size)
	return args.Get(0).([]model.Task), args.Get(1).(sharedmodel.Paging), args.Error(2)
}

func (m *TaskRepositoryMock) Create(payload model.Task) (model.Task, error) {
	args := m.Called(payload)
	return args.Get(0).(model.Task), args.Error(1)
}

func (m *TaskRepositoryMock) GetByAuthorID(authorID string) ([]model.Task, error) {
	args := m.Called(authorID)
	return args.Get(0).([]model.Task), args.Error(1)
}

func (m *TaskRepositoryMock) Delete(taskID string) error {
	panic("not impl")
}
