package usecase

import (
	"fmt"
	"todo-clean-arch/model"
	"todo-clean-arch/repository"
)

type TaskUsecase interface {
	RegisterNewTask(payload model.Task) (model.Task, error)
	FindTaskByAuthor(id string) ([]model.Task, error)
	FindAllTask() ([]model.Task, error)
}

type taskUsecase struct {
	taskRepository repository.TaskRepository
	authorUC       AuthorUsecase
}

func (t *taskUsecase) RegisterNewTask(payload model.Task) (model.Task, error) {
	// Validate payload
	_, err := t.authorUC.FindAuthorByID(payload.AuthorID)
	if err != nil {
		return model.Task{}, fmt.Errorf("user doesnt exist: %v", err)
	}
	if payload.Content == "" || payload.Title == "" {
		return model.Task{}, fmt.Errorf("content or cannot be empty")
	}

	task, err := t.taskRepository.Create(payload)
	if err != nil {
		return model.Task{}, fmt.Errorf("failed to create task %v", err)
	}
	return task, nil
}

func (t *taskUsecase) FindTaskByAuthor(id string) ([]model.Task, error) {
	task, err := t.taskRepository.GetByAuthorID(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (t *taskUsecase) FindAllTask() ([]model.Task, error) {
	tasks, err := t.taskRepository.List()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func NewTaskUseCase(taskRepository repository.TaskRepository, authorUC AuthorUsecase) TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
		authorUC:       authorUC,
	}
}
