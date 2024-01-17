package usecase

import (
	"fmt"
	"todo-clean-arch/model"
	"todo-clean-arch/repository"
	sharedmodel "todo-clean-arch/shared/shared_model"
)

type TaskUsecase interface {
	RegisterNewTask(payload model.Task) (model.Task, error)
	FindTaskByAuthor(id string) ([]model.Task, error)
	FindAllTask(page int, size int) ([]model.Task, sharedmodel.Paging, error)
	RemoveTask(id string) error
}

type taskUsecase struct {
	taskRepository repository.TaskRepository
	authorUC       AuthorUsecase
}

func (t *taskUsecase) RemoveTask(id string) error {
	err := t.taskRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete %v", err)
	}
	return nil
}

func NewTaskUseCase(taskRepository repository.TaskRepository, authorUC AuthorUsecase) TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
		authorUC:       authorUC,
	}
}

func (t *taskUsecase) RegisterNewTask(payload model.Task) (model.Task, error) {
	// Validate payload
	_, err := t.authorUC.FindAuthorByID(payload.AuthorID)
	if err != nil {
		return model.Task{}, fmt.Errorf("user doesnt exist: %v", err)
	}

	task, err := t.taskRepository.Create(payload)
	if err != nil {
		return model.Task{}, err
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

func (t *taskUsecase) FindAllTask(page int, size int) ([]model.Task, sharedmodel.Paging, error) {
	tasks, paging, err := t.taskRepository.List(page, size)
	if err != nil {
		return nil, paging, err
	}
	return tasks, paging, nil
}
