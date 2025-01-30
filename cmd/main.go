package main

import (
	"log"
	"managingtasks/config"
	"managingtasks/infrastructure/database"
	"managingtasks/interface/handler"
	"managingtasks/internal/repository"
	"managingtasks/internal/usecase"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()

	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	err = database.InitDB(db)
	if err != nil {
		log.Fatal("Failed to initialize database: ", err)
	}

	taskRepo := repository.NewTaskRepository(db)

	taskUseCase := usecase.NewTaskUseCase(taskRepo)

	taskHandler := handler.NewTaskHandler(taskUseCase)

	r := mux.NewRouter()
	r.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", taskHandler.AddTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
