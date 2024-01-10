package repository

import "todo-clean-arch/model"

type AuthorRepository interface {
	GetByEmail(email string) (model.Author, error)
	Get(id string) (model.Author, error)
	List() ([]model.Author, error)
}
