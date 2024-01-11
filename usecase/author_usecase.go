package usecase

import (
	"fmt"
	"todo-clean-arch/model"
	"todo-clean-arch/repository"
)

type AuthorUsecase interface {
	FindAuthorByID(id string) (model.Author, error)
	FindAuthorByEmail(email string) (model.Author, error)
	FindAllAuthor(id string) ([]model.Author, error)
	UpdateAuthor(payload model.Author) (model.Author, error)
	RemoveAuthor(id string) error
}

type authorUsecase struct {
	authorRepository repository.AuthorRepository
}

func (a *authorUsecase) RemoveAuthor(id string) error {
	_, err := a.FindAuthorByID(id)
	if err != nil {
		return fmt.Errorf("user doesnt exist: %v", err)
	}
	err = a.authorRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete %v", err)
	}
	return nil
}

func (a *authorUsecase) UpdateAuthor(payload model.Author) (model.Author, error) {
	_, err := a.FindAuthorByID(payload.ID)
	if err != nil {
		return model.Author{}, fmt.Errorf("user doesnt exist: %v", err)
	}
	author, err := a.authorRepository.Update(payload)
	if err != nil {
		return model.Author{}, err
	}
	return author, nil
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

func (a *authorUsecase) FindAllAuthor(id string) ([]model.Author, error) {
	authors, err := a.authorRepository.List(id)
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
