package handlers

import "net/http"

func GetBooks(response http.ResponseWriter, request *http.Request) {

	response.WriteHeader(http.StatusOK)
}
