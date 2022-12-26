package main

import (
	"book-microservice/src/data"
	"book-microservice/src/database"
	"book-microservice/src/middleware"
	"book-microservice/src/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database.ConnectDb(data.Books)

	router := mux.NewRouter()

	router.Use(middleware.JsonContentTypeResponse)

	routes.SetupRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" //localhost
	}

	fmt.Println("Starting service at port: " + port)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedHeaders: []string{"Origin, Content-Type, Accept, Authorization"},
	})

	handler := c.Handler(router)

	err = http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatalf("Server crashed, reason:" + err.Error())
	}
}
