package usecase

import (
	"todo-clean-arch/model"
	"todo-clean-arch/repository"
)

type AuthorUsecase interface {
	FindAuthorByID(id string) (model.Author, error)
	FindAuthorByEmail(email string) (model.Author, error)
	FindAllAuthor() ([]model.Author, error)
}

type authorUsecase struct {
	authorRepository repository.AuthorRepository
}

func (a *authorUsecase) FindAuthorByID(id string) (model.Author, error) {
	author, err := a.authorRepository.Get(id)
	if err != nil {
		return model.Author{}, err
	}
	return author, nil
}

func (a *authorUsecase) FindAuthorByEmail(email string) (model.Author, error) {
	author, err := a.authorRepository.GetByEmail(email)
	if err != nil {
		return model.Author{}, err
	}
	return author, nil
}

func (a *authorUsecase) FindAllAuthor() ([]model.Author, error) {
	authors, err := a.authorRepository.List()
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func NewAuthorUseCase(authorRepository repository.AuthorRepository) AuthorUsecase {
	return &authorUsecase{
		authorRepository: authorRepository,
	}
}
