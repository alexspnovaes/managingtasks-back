package repository

import (
	"errors"
	"managingtasks/internal/entity"
)

type MockTaskRepository struct {
	Tasks []entity.Task
}

func (m *MockTaskRepository) GetAllTasks() ([]entity.Task, error) {
	return m.Tasks, nil
}

func (m *MockTaskRepository) CreateTask(task *entity.Task) error {
	if task.Title == "" {
		return errors.New("title cannot be empty")
	}
	task.ID = len(m.Tasks) + 1
	m.Tasks = append(m.Tasks, *task)
	return nil
}

func (m *MockTaskRepository) UpdateTask(task *entity.Task) error {
	for i, t := range m.Tasks {
		if t.ID == task.ID {
			m.Tasks[i] = *task
			return nil
		}
	}
	return errors.New("task not found")
}

func (m *MockTaskRepository) DeleteTask(id int) error {
	for i, t := range m.Tasks {
		if t.ID == id {
			m.Tasks = append(m.Tasks[:i], m.Tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
