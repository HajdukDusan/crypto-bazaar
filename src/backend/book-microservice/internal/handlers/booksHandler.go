package handlers

import (
	"book-microservice/internal/database"
	"encoding/json"
	"net/http"
)

func GetAll(response http.ResponseWriter, request *http.Request) {

	books := database.GetAll()

	json.NewEncoder(response).Encode(books)
}
