package handlers

import (
	"book-microservice/src/database"
	"book-microservice/src/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllPagedBooks(response http.ResponseWriter, request *http.Request) {

	query := request.URL.Query()

	limit, _ := strconv.ParseUint(query.Get("limit"), 10, 64)
	page, _ := strconv.ParseUint(query.Get("page"), 10, 64)

	sort := query.Get("sort")

	pagination := database.Pagination{
		Limit: int(limit),
		Page:  int(page),
		Sort:  sort,
	}

	books := database.GetAll(pagination)

	json.NewEncoder(response).Encode(books)
}

func GetOneBook(response http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	id, _ := strconv.ParseUint(params["id"], 10, 64)

	book := database.GetById(uint(id))

	if book == nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(response).Encode(book)
}

func CreateBook(response http.ResponseWriter, request *http.Request) {

	var book *model.Book

	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	result := database.Create(book)

	if result.Error != nil {
		response.WriteHeader(http.StatusBadRequest)
	} else {
		response.WriteHeader(http.StatusCreated)
	}
}

func UpdateBook(response http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	id, _ := strconv.ParseUint(params["id"], 10, 64)

	var book *model.Book

	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	result := database.Update(id, book)

	if result.Error != nil {
		response.WriteHeader(http.StatusBadRequest)
	} else {
		response.WriteHeader(http.StatusNoContent)
	}
}

func DeleteBook(response http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	id, _ := strconv.ParseUint(params["id"], 10, 64)

	book := database.GetById(uint(id))

	if book == nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	result := database.Delete(id)

	if result.Error != nil {
		response.WriteHeader(http.StatusBadRequest)
	} else {
		response.WriteHeader(http.StatusNoContent)
	}
}
