package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetSpecificBook(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, error := strconv.Atoi(params["id"])
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
	}

	var book Book
	for _, value := range AllBooks {
		if value.BookID == id {
			book = value
			json.NewEncoder(w).Encode(book)
			return


			}

	}
	http.Error(w,"Book Not Found",http.StatusNotFound)
	

}

