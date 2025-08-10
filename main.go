package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
)

// Define a simple struct
type Message struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

var db *sql.DB

// In-memory data
// var messages = []Message{
// 	{ID: 1, Text: "Hello, World!"},
// 	{ID: 2, Text: "Bonjour, Monde!"},
// }

func main() {
	// r := chi.NewRouter()
	// r.Use(middleware.Logger)

	// // Routes
	// r.Get("/messages", getMessages)
	// r.Post("/messages", createMessage)

	// http.ListenAndServe(":8080", r)

	//#####22222

	// var err error

	// // DB connection from env vars (docker-compose will set these)
	// connStr := os.Getenv("DATABASE_URL")
	// if connStr == "" {
	// 	connStr = "postgres://user:pass@db:5432/mydb?sslmode=disable"
	// }

	// db, err = sql.Open("postgres", connStr)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Ensure table exists
	// _, err = db.Exec(`CREATE TABLE IF NOT EXISTS messages (
	// 	id SERIAL PRIMARY KEY,
	// 	text TEXT NOT NULL
	// )`)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// r := chi.NewRouter()
	// r.Use(middleware.Logger)

	// r.Get("/messages", getMessages)
	// r.Post("/messages", createMessage)

	// log.Println("API running on :8080")
	// http.ListenAndServe(":8080", r)

	// Read DATABASE_URL from environment or fallback to local
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/flowers4you?sslmode=disable"
	}

	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Database is not reachable:", err)
	}

	log.Println("Connected to database")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/messages", getMessages)
	r.Post("/messages", createMessage)

	log.Println("API running on :8080")
	http.ListenAndServe(":8080", r)
}

// Handlers
func getMessages(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(messages)

	rows, err := db.Query("SELECT id, text FROM messages ORDER BY id ASC")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var msgs []Message
	for rows.Next() {
		var m Message
		if err := rows.Scan(&m.ID, &m.Text); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		msgs = append(msgs, m)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msgs)
}

func createMessage(w http.ResponseWriter, r *http.Request) {
	// var msg Message
	// err := json.NewDecoder(r.Body).Decode(&msg)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// msg.ID = len(messages) + 1
	// messages = append(messages, msg)

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(msg)

	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err := db.QueryRow(
		"INSERT INTO messages (text) VALUES ($1) RETURNING id", msg.Text,
	).Scan(&msg.ID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
}
