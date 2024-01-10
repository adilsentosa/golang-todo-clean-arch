package repository

import (
	"database/sql"
	"log"
	"todo-clean-arh/model"
)

type TaskRepository interface {
	Create(payload model.Task) (model.Task, error)
	List() []model.Task
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
	err := t.db.QueryRow("INSERT INTO tasks (title, content,author_id) VALUES ($1 $2 $3) RETURNING id, created_at",
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
	return []model.Task{}, nil
}

func (t *taskRepository) List() []model.Task {
	return []model.Task{}
}
