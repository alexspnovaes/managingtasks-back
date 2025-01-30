package usecase

import (
	"errors"
	"managingtasks/internal/entity"
	"managingtasks/internal/repository"
)

type TaskUseCase struct {
	taskRepo repository.TaskRepository
}

func NewTaskUseCase(taskRepo repository.TaskRepository) *TaskUseCase {
	return &TaskUseCase{taskRepo: taskRepo}
}

func (uc *TaskUseCase) GetTasks() ([]entity.Task, error) {
	return uc.taskRepo.GetAllTasks()
}

func (uc *TaskUseCase) CreateTask(title string) error {
	if title == "" {
		return errors.New("title cannot be empty")
	}
	task := &entity.Task{Title: title, Completed: false}
	return uc.taskRepo.CreateTask(task)
}

func (uc *TaskUseCase) UpdateTask(id int, title string, completed bool) error {
	if title == "" {
		return errors.New("title cannot be empty")
	}
	task := &entity.Task{ID: id, Title: title, Completed: completed}
	return uc.taskRepo.UpdateTask(task)
}

func (uc *TaskUseCase) DeleteTask(id int) error {
	return uc.taskRepo.DeleteTask(id)
}
