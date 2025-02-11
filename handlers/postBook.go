package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/redis"
	"github.com/google/uuid"
	"restapi/handlers/channel"

)


type Book struct {
	BookID   int    `json:"bookId" redis:"bookId"`
	BookName string `json:"bookName" redis:"bookName"`
	Author   string `json:"author" redis:"author"`
}

//all books array
var AllBooks=[]Book {}

func PostBook(w http.ResponseWriter,req *http.Request){

	var book Book;
	err:=json.NewDecoder(req.Body).Decode(&book)
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
	fmt.Println(book.BookID)
	msg:=fmt.Sprintf("/books - Book ID :%d",book.BookID)
	channel.AddToChannel("POST",msg)
	// data,err:=json.Marshal(book)
	key:=fmt.Sprintf("Book:%d",book.BookID)
	errorredis := redis.Client.HSet(redis.Ctx, key, map[string]interface{}{
		"bookId":   book.BookID,
		"bookName": book.BookName,
		"author":   book.Author,
	}).Err()

	if errorredis!=nil{
		http.Error(w,"eror occured in adding to redis hash",http.StatusBadRequest)
	}


	
	



}



