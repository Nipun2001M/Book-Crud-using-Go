package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"restapi/handlers/channel"
	"github.com/gorilla/mux"
)

func UpdateBook(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id, error := strconv.Atoi(params["id"])
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}

	var book Book
	var found bool = false
	var indexofbook int
	for index, value := range AllBooks {
		if value.BookID == id {
			book = value
			found = true
			indexofbook = index
			json.NewEncoder(w).Encode(book)

		}

	}

	if !found {
		http.Error(w, "Book Not Found", http.StatusNotFound)
		return
	}

	var newBook Book
	err := json.NewDecoder(req.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if newBook.Author != "" {
		AllBooks[indexofbook].Author = newBook.Author
	}
	if newBook.BookName != "" {
		AllBooks[indexofbook].BookName = newBook.BookName
	}
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "sucessfully updated",
	}
	
	msg:=fmt.Sprintf("/books - Book ID :%d",book.BookID)
	channel.AddToChannel("PUT",msg)
	json.NewEncoder(w).Encode(response)

}
