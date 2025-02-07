package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)


type Book struct {
	BookID   int    `json:"bookId"`
	BookName string `json:"bookName"`
	Author   string `json:"author"`
}

//all books array
var AllBooks=[]Book {}

func PostBook(w http.ResponseWriter,req *http.Request){

	var book Book;
	err:=json.NewDecoder(req.Body).Decode(&book)
	fmt.Println("Request Body:", req.Body)

	fmt.Println("book",book)
	book.BookID=uuid.New().ClockSequence();
	
	if err != nil  {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
		}
	if book.Author==""||book.BookName==""{
		http.Error(w,"All fields required",http.StatusBadRequest)
		return
	}	



	AllBooks = append(AllBooks, book)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	response:=map[string]string{
		"message":"book added sucessfully",
	}
	json.NewEncoder(w).Encode(response)






}
