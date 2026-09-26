package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	response := map[string] string {"status": "OK"}
	json.NewEncoder(w).Encode(response)

	w.WriteHeader(http.StatusOK)
}

func main() {
	mux := http.NewServeMux() 
	mux.HandleFunc("/health", healthHandler)

	log.Fatal(http.ListenAndServe(":3000", mux))
}
