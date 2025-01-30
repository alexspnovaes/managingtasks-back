package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"managingtasks/interface/handler"
	"managingtasks/internal/entity"
	"managingtasks/internal/repository"
	"managingtasks/internal/usecase"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetTasks(t *testing.T) {
	mockRepo := &repository.MockTaskRepository{}
	mockUseCase := usecase.NewTaskUseCase(mockRepo)
	taskHandler := handler.NewTaskHandler(mockUseCase)
	req, _ := http.NewRequest("GET", "/tasks", nil)
	res := httptest.NewRecorder()
	taskHandler.GetTasks(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestAddTask(t *testing.T) {
	mockRepo := &repository.MockTaskRepository{}
	mockUseCase := usecase.NewTaskUseCase(mockRepo)
	taskHandler := handler.NewTaskHandler(mockUseCase)

	task := entity.Task{Title: "Test Task"}
	body, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	taskHandler.AddTask(res, req)

	assert.Equal(t, http.StatusCreated, res.Code)
}

func TestAddTaskWithEmptyTitle(t *testing.T) {
	mockRepo := &repository.MockTaskRepository{}
	mockUseCase := usecase.NewTaskUseCase(mockRepo)
	taskHandler := handler.NewTaskHandler(mockUseCase)

	task := entity.Task{Title: ""}
	body, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	taskHandler.AddTask(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code)
}

func TestUpdateTask(t *testing.T) {
	mockRepo := &repository.MockTaskRepository{}
	mockUseCase := usecase.NewTaskUseCase(mockRepo)
	taskHandler := handler.NewTaskHandler(mockUseCase)
	mockRepo.CreateTask(&entity.Task{Title: "Task to be updated"})

	task := entity.Task{ID: 1, Title: "Updated Task", Completed: true}
	body, _ := json.Marshal(task)

	req, _ := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	vars := map[string]string{"id": "1"}
	req = mux.SetURLVars(req, vars)
	taskHandler.UpdateTask(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestDeleteTask(t *testing.T) {
	mockRepo := &repository.MockTaskRepository{}
	mockUseCase := usecase.NewTaskUseCase(mockRepo)
	taskHandler := handler.NewTaskHandler(mockUseCase)
	mockRepo.CreateTask(&entity.Task{Title: "Task to be deleted"})

	req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
	vars := map[string]string{"id": "1"}
	req = mux.SetURLVars(req, vars)
	res := httptest.NewRecorder()

	taskHandler.DeleteTask(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}
