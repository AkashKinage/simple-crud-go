package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/AkashKinage/simple-crud-go/internal/config"

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
	_ = db
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux() 
	mux.HandleFunc("/health", healthHandler)

	log.Fatal(http.ListenAndServe(":3000", mux))
}
