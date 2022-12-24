package handlers

import (
	"book-microservice/database"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllPaged(response http.ResponseWriter, request *http.Request) {

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

func GetOne(response http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)

	id, _ := strconv.ParseUint(params["id"], 10, 64)

	books := database.GetById(uint(id))

	if books == nil {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(response).Encode(books)
}
