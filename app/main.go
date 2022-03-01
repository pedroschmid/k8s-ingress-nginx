package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", Default)

	http.ListenAndServe(":8080", nil)
}

func Default(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	data := Response{Status: http.StatusOK, Message: "Welcome to Golang API =D"}

	json.NewEncoder(w).Encode(data)
}
