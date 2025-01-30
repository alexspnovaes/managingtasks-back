package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"managingtasks/internal/usecase"

	"github.com/gorilla/mux"
)

type TaskHandler struct {
	taskUseCase *usecase.TaskUseCase
}

func NewTaskHandler(taskUseCase *usecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{taskUseCase: taskUseCase}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.taskUseCase.GetTasks()
	if err != nil {
		http.Error(w, "Failed to retrieve tasks", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) AddTask(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string `json:"title"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := h.taskUseCase.CreateTask(input.Title)
	if err != nil {
		if err.Error() == "title cannot be empty" {
			http.Error(w, "Title is required", http.StatusBadRequest)
			return
		}
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err = h.taskUseCase.UpdateTask(id, input.Title, input.Completed)
	if err != nil {
		if err.Error() == "task not found" {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	err = h.taskUseCase.DeleteTask(id)
	if err != nil {
		if err.Error() == "task not found" {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
