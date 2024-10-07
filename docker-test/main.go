package main

import (
	"docker-app/handler"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	handler.Get()
	// Define some routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Chi Router 222222222222222222!"))
	})

	r.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	fmt.Println("server started at :8080")
	http.ListenAndServe(":8080", r)
}
