package repository

import (
	"database/sql"
	"log"
	"math"
	"time"
	"todo-clean-arch/model"
	sharedmodel "todo-clean-arch/shared/shared_model"
)

type TaskRepository interface {
	Create(payload model.Task) (model.Task, error)
	List(page int, size int) ([]model.Task, sharedmodel.Paging, error)
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
	currTime := time.Now()
	payload.UpdatedAt = &currTime
	err := t.db.QueryRow("INSERT INTO tasks (title, content,author_id,updated_at) VALUES ($1, $2, $3, $4) RETURNING id, created_at",
		payload.Title, payload.Content, payload.AuthorID, payload.UpdatedAt).Scan(&task.ID, &task.CreatedAt)
	if err != nil {
		log.Println("taskRepository.QueryRow", err.Error())
		return model.Task{}, err
	}
	task.Title = payload.Title
	task.Content = payload.Content
	task.AuthorID = payload.AuthorID
	task.UpdatedAt = payload.UpdatedAt
	return task, nil
}

func (t *taskRepository) GetByAuthorID(authorID string) ([]model.Task, error) {
	var tasks []model.Task
	query := "SELECT id,title,content,created_at,updated_at FROM tasks WHERE author_id = $1"

	rows, err := t.db.Query(query, authorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task model.Task
		task.AuthorID = authorID
		err := rows.Scan(&task.ID, &task.Title, &task.Content, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (t *taskRepository) List(page int, size int) ([]model.Task, sharedmodel.Paging, error) {
	var tasks []model.Task
	offset := (page - 1) * size
	query := `SELECT id,title,content,author_id,created_at 
  FROM tasks ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	rows, err := t.db.Query(query, offset)
	if err != nil {
		log.Println("taskRepository.Query:", err.Error())
		return nil, sharedmodel.Paging{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Content, &task.AuthorID, &task.CreatedAt)
		if err != nil {
			log.Println("taskRepository.Rows.Next():")
			return nil, sharedmodel.Paging{}, err
		}
		tasks = append(tasks, task)
	}

	totalRows := 0
	if err := t.db.QueryRow("SELECT COUNT(*) FROM tasks").Scan(&totalRows); err != nil {
		return nil, sharedmodel.Paging{}, err
	}

	paging := sharedmodel.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}

	return tasks, paging, nil
}
