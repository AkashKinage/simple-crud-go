package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/AkashKinage/simple-crud-go/internal/config"
	"github.com/AkashKinage/simple-crud-go/internal/handler"
	"github.com/AkashKinage/simple-crud-go/internal/repository"
	"github.com/AkashKinage/simple-crud-go/internal/service"

	_ "github.com/lib/pq"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	response := map[string] string {"status": "OK"}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	cfg := config.Load()
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux() 
	mux.HandleFunc("/health", healthHandler)

	taskRepo := repository.NewPostgresTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	mux.HandleFunc("GET /tasks", taskHandler.List)
	mux.HandleFunc("POST /tasks", taskHandler.Create)
	mux.HandleFunc("GET /tasks/{id}", taskHandler.GetByID)

	log.Fatal(http.ListenAndServe(":3000", mux))
}
