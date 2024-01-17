package usecase_mock

import (
	"todo-clean-arch/model"

	"github.com/stretchr/testify/mock"
)

type AuthorUseCaseMock struct {
	mock.Mock
}

func (m *AuthorUseCaseMock) FindAllAuthor(id string) ([]model.Author, error) {
	args := m.Called(id)
	return args.Get(0).([]model.Author), args.Error(1)
}

func (m *AuthorUseCaseMock) FindAuthorByID(id string) (model.Author, error) {
	args := m.Called(id)
	return args.Get(0).(model.Author), args.Error(1)
}

func (m *AuthorUseCaseMock) FindAuthorByEmail(email string) (model.Author, error) {
	args := m.Called(email)
	return args.Get(0).(model.Author), args.Error(1)
}

func (m *AuthorUseCaseMock) UpdateAuthor(payload model.Author) (model.Author, error) {
	panic("not implemented")
}

func (m *AuthorUseCaseMock) RemoveAuthor(id string) error {
	panic("not impl")
}
