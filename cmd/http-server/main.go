// http-response.go
// Esempio di un semplice server HTTP in Go che risponde a richieste su due
package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Benvenuto alla homepage!")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Chi siamo")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	fmt.Println("Server attivo su :8080...")
	http.ListenAndServe(":8080", nil)
}
