package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Define a simple struct
type Message struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

// In-memory data
var messages = []Message{
	{ID: 1, Text: "Hello, World!"},
	{ID: 2, Text: "Bonjour, Monde!"},
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Routes
	r.Get("/messages", getMessages)
	r.Post("/messages", createMessage)

	http.ListenAndServe(":8080", r)
}

// Handlers
func getMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func createMessage(w http.ResponseWriter, r *http.Request) {
	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	msg.ID = len(messages) + 1
	messages = append(messages, msg)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
}
