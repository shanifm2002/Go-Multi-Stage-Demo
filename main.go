package main

import (
	"encoding/json"
	"net/http"
	"os"
)

type Health struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, This is Go App 🚀"))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	response := Health{
		Status:      "healthy",
		Environment: env,
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)

	http.ListenAndServe(":8080", nil)
}
