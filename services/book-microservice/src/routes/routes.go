package routes

import (
	"book-microservice/src/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/books", handlers.GetAllPagedBooks).Methods("GET")
	router.HandleFunc("/books/{id}", handlers.GetOneBook).Methods("GET")

	router.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")
}
