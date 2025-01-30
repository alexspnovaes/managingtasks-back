package repository

import (
	"database/sql"
	"managingtasks/internal/entity"
)

type TaskRepository interface {
	GetAllTasks() ([]entity.Task, error)
	CreateTask(task *entity.Task) error
	UpdateTask(task *entity.Task) error
	DeleteTask(id int) error
}

type TaskRepositoryImpl struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepositoryImpl {
	return &TaskRepositoryImpl{db: db}
}

func (r *TaskRepositoryImpl) GetAllTasks() ([]entity.Task, error) {
	rows, err := r.db.Query("SELECT id, title, completed FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []entity.Task
	for rows.Next() {
		var task entity.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Completed); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *TaskRepositoryImpl) CreateTask(task *entity.Task) error {
	_, err := r.db.Exec("INSERT INTO tasks (title, completed) VALUES (?, ?)", task.Title, task.Completed)
	return err
}

func (r *TaskRepositoryImpl) UpdateTask(task *entity.Task) error {
	_, err := r.db.Exec("UPDATE tasks SET title = ?, completed = ? WHERE id = ?", task.Title, task.Completed, task.ID)
	return err
}

func (r *TaskRepositoryImpl) DeleteTask(id int) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}
