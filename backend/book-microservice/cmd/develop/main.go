package main

import (
	"book-microservice/data"
	"book-microservice/internal/database"
	"book-microservice/internal/handlers"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	err := database.CreateDatabase(data.Books)

	if err != nil {
		fmt.Println(err)
		return
	}

	router := mux.NewRouter()

	router.HandleFunc("/book", handlers.GetAllPaged).Methods("GET")
	router.HandleFunc("/book/{id}", handlers.GetOne).Methods("GET")

	// router.HandleFunc("/book", handlers.GetBooks).Methods("POST")
	// router.HandleFunc("/book/{id}", handlers.GetBooks).Methods("PUT")
	// router.HandleFunc("/book/{id}", handlers.GetBooks).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "9001" //localhost
	}

	fmt.Println("Starting service at port: " + port)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedHeaders: []string{"Origin, Content-Type, Accept, Authorization"},
	})

	handler := c.Handler(router)

	err = http.ListenAndServe(":"+port, handler)
	if err != nil {
		fmt.Print(err)
	}
}
