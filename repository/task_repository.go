package repository

import (
	"database/sql"
	"log"
	"todo-clean-arch/model"
)

type TaskRepository interface {
	Create(payload model.Task) (model.Task, error)
	List() ([]model.Task, error)
	GetByAuthorID(authorID string) ([]model.Task, error)
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (t *taskRepository) Create(payload model.Task) (model.Task, error) {
	var task model.Task
	err := t.db.QueryRow("INSERT INTO tasks (title, content,author_id) VALUES ($1, $2, $3) RETURNING id, created_at",
		payload.Title, payload.Content, payload.AuthorID).Scan(&task.ID, &task.CreatedAt)
	if err != nil {
		log.Println("taskRepository.QueryRow", err.Error())
		return model.Task{}, err
	}
	task.Title = payload.Title
	task.Content = payload.Content
	task.AuthorID = payload.AuthorID
	return task, nil
}

func (t *taskRepository) GetByAuthorID(authorID string) ([]model.Task, error) {
	var tasks []model.Task
	query := "SELECT id,title,content,created_at FROM tasks WHERE author_id = $1"

	rows, err := t.db.Query(query, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task model.Task
		task.AuthorID = authorID
		err := rows.Scan(&task.ID, &task.Title, &task.Content, &task.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (t *taskRepository) List() ([]model.Task, error) {
	var tasks []model.Task
	query := "SELECT id,title,content,author_id,created_at FROM tasks"

	rows, err := t.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Content, &task.AuthorID, &task.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
