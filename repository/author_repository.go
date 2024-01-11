package repository

import (
	"database/sql"
	"fmt"
	"time"
	"todo-clean-arch/config"
	"todo-clean-arch/model"
)

type AuthorRepository interface {
	GetByEmail(email string) (model.Author, error)
	Get(id string) (model.Author, error)
	List(id string) ([]model.Author, error)
	Update(author model.Author) (model.Author, error)
	Delete(id string) error
}

type authorRepository struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) AuthorRepository {
	return &authorRepository{
		db: db,
	}
}

func (a *authorRepository) Delete(id string) error {
	query := config.DeleteAuthorByID
	_, err := a.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (a *authorRepository) Update(author model.Author) (model.Author, error) {
	query := config.UpdateAuthorByID
	currentTime := time.Now()
	author.UpdatedAt = &currentTime

	_, err := a.db.Exec(query, author.ID, author.Name, author.Email, author.Password, author.Role, author.UpdatedAt)
	if err != nil {
		return model.Author{}, err
	}

	return author, nil
}

func (a *authorRepository) Get(id string) (model.Author, error) {
	var author model.Author
	query := config.SelectAuthorById

	err := a.db.QueryRow(query, id).Scan(&author.ID, &author.Name, &author.Email, &author.CreatedAt, &author.UpdatedAt)
	if err != nil {
		return model.Author{}, err
	}

	return author, nil
}

// func (a *authorRepository) Get(id string) (model.Author, error) {
// 	var author model.Author
// 	// query := "SELECT id,name,email,created_at,updated_at FROM authors WHERE id = $1"
// 	query := `SELECT
// 	  a.id,
// 	  a.name,
// 	  a.email,
//     a.updated_at,
// 	  a.created_at,
// 	  array_agg(
// 		  jsonb_build_object(
// 			  'id', t.id,
// 			  'title', t.title,
// 			  'content', t.content,
// 			  'author_ID', t.author_id,
// 			  'created_at', t.created_at
// 		  )
// 	  ) AS tasks
//   FROM authors a
//   JOIN tasks t  ON a.id = t.author_id
//   WHERE a.id = $1
//   GROUP BY a.id,a.name, a.email,a.created_at,a.updated_at`
// 	var tasks []model.Task
// 	// var tasks []map[string]interface{}
// 	var stringTasks string
//
// 	err := a.db.QueryRow(query, id).Scan(&author.ID, &author.Name, &author.Email, &author.CreatedAt, &author.UpdatedAt, &stringTasks)
// 	if err != nil {
// 		return model.Author{}, err
// 	}
// 	err = json.Unmarshal([]byte(stringTasks), &tasks)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Printf("%v", stringTasks)
// 	fmt.Println()
// 	fmt.Println(tasks)
//
// 	return author, nil
// }

func (a *authorRepository) GetByEmail(email string) (model.Author, error) {
	var author model.Author
	query := config.SelectAuthorByEmail
	err := a.db.QueryRow(query, email).Scan(&author.ID, &author.Name, &author.Email, &author.CreatedAt, &author.UpdatedAt)
	if err != nil {
		return model.Author{}, err
	}

	return author, nil
}

func (t *authorRepository) List(id string) ([]model.Author, error) {
	query := ""
	var role string

	queryAuthor := "SELECT role FROM authors WHERE id = $1"
	var err error
	err = t.db.QueryRow(queryAuthor, id).Scan(&role)
	if err != nil {
		return nil, err
	}
	fmt.Println("query author passed")
	var rows *sql.Rows

	if role == "admin" {
		query = config.SelectAuthorWithTasks
		rows, err = t.db.Query(query)
	} else {
		query = config.SelectAuthorWithTasksByID
		rows, err = t.db.Query(query, id)
	}
	if err != nil {
		return nil, err
	}

	fmt.Println("query all passed")
	defer rows.Close()

	authors := make(map[string]*model.Author)

	for rows.Next() {
		var task model.Task
		var author model.Author
		err := rows.Scan(&author.ID, &author.Name, &author.Email, &author.Role, &author.CreatedAt, &author.UpdatedAt, &task.ID, &task.Title,
			&task.Content, &task.AuthorID, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		if _, ok := authors[author.ID]; !ok {
			authors[author.ID] = &author
		}
		authors[author.ID].Tasks = append(authors[author.ID].Tasks, task)
	}

	authorSlice := make([]model.Author, 0, len(authors))
	for _, author := range authors {
		authorSlice = append(authorSlice, *author)
	}

	return authorSlice, nil
}
