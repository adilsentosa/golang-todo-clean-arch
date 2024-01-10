package repository

import (
	"database/sql"
	"todo-clean-arch/model"
)

type AuthorRepository interface {
	GetByEmail(email string) (model.Author, error)
	Get(id string) (model.Author, error)
	List(id string) ([]model.Author, error)
}

type authorRepository struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) AuthorRepository {
	return &authorRepository{
		db: db,
	}
}

func (a *authorRepository) Get(id string) (model.Author, error) {
	var author model.Author
	query := "SELECT id,name,email,created_at FROM authors WHERE id = $1"

	err := a.db.QueryRow(query, id).Scan(&author.ID, &author.Name, &author.Email, &author.CreatedAt)
	if err != nil {
		return model.Author{}, err
	}

	return author, nil
}

func (a *authorRepository) GetByEmail(email string) (model.Author, error) {
	var author model.Author
	query := "SELECT id,name,email,created_at FROM authors WHERE email = $1"

	err := a.db.QueryRow(query, email).Scan(&author.ID, &author.Name, &author.Email, &author.CreatedAt)
	if err != nil {
		return model.Author{}, err
	}

	return author, nil
}

func (t *authorRepository) List(id string) ([]model.Author, error) {
	var authors []model.Author
	query := "SELECT id,name,email,created_at FROM authors"

	rows, err := t.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var author model.Author
		err := rows.Scan(&author.ID, &author.Name, &author.Email, &author.CreatedAt)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}

	return authors, nil
}
