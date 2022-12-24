package routes

import (
	"book-microservice/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/book", handlers.GetAllPaged).Methods("GET")
	router.HandleFunc("/book/{id}", handlers.GetOne).Methods("GET")

	// router.HandleFunc("/book", handlers.GetBooks).Methods("POST")
	// router.HandleFunc("/book/{id}", handlers.GetBooks).Methods("PUT")
	// router.HandleFunc("/book/{id}", handlers.GetBooks).Methods("DELETE")
}
