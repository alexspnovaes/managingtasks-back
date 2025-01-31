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
	r.Use(corsMiddleware)
	r.HandleFunc("/tasks", taskHandler.GetTasks).Methods("GET", "OPTIONS")
	r.HandleFunc("/tasks", taskHandler.AddTask).Methods("POST", "OPTIONS")
	r.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT", "OPTIONS")
	r.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE", "OPTIONS")
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
