package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/redis"
	"github.com/google/uuid"
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

	data,err:=json.Marshal(book)
	var errorredis error=redis.Client.SAdd(redis.Ctx,"BookSet",data,0).Err()

	if errorredis!=nil{
		http.Error(w,"eror occured in adding to redis hash",http.StatusBadRequest)
	}

	
	




	






}



