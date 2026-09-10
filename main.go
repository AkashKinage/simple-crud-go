package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Healthy")
}

func main() {
	http.HandleFunc("/health", healthHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
