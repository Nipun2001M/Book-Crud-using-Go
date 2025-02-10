package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"restapi/handlers/channel"
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
			msg:=fmt.Sprintf("/books - Book ID :%d",id)
			channel.AddToChannel("GET",msg)
			json.NewEncoder(w).Encode(book)
			return


			}

	}
	http.Error(w,"Book Not Found",http.StatusNotFound)
	

}

