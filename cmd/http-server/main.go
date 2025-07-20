package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Modello dati
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// Memoria simulata
var todos = []Todo{
	{ID: 1, Title: "Comprare il latte", Done: false},
	{ID: 2, Title: "Studiare Go", Done: true},
}

// 🔍 GET /todos - restituisce tutti i todo
func getTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

// 🔍 GET /todos/{id}
func getTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, todo := range todos {
		if todo.ID == id {
			json.NewEncoder(w).Encode(todo)
			return
		}
	}

	http.NotFound(w, r)
}

// ➕ POST /todos - crea un nuovo todo
func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = len(todos) + 1
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

// 📝 PUT /todos/{id} - aggiorna un todo esistente
func updateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, t := range todos {
		if t.ID == id {
			var updated Todo
			_ = json.NewDecoder(r.Body).Decode(&updated)
			updated.ID = id
			todos[i] = updated
			json.NewEncoder(w).Encode(updated)
			return
		}
	}

	http.NotFound(w, r)
}

// ❌ DELETE /todos/{id}
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			fmt.Fprintf(w, "Eliminato todo con ID %d", id)
			return
		}
	}

	http.NotFound(w, r)
}

// 🧠 main
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/todos", getTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", getTodo).Methods("GET")
	r.HandleFunc("/todos", createTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", updateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")

	fmt.Println("REST server su http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
