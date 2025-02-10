package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/redis"
)



func GetAllBooks(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)	
	w.Header().Set("Content-Type","application/json")
	books,resulterrro:=redis.Client.SMembers(redis.Ctx,"BookSet").Result()
	if resulterrro!=nil{
		fmt.Println(resulterrro)
	}
	fmt.Println(books)
	var booksStruct []Book
	for _, bookStr := range books {
		var book Book
		if json.Unmarshal([]byte(bookStr), &book) == nil {
			booksStruct = append(booksStruct, book)
		}
	}

	json.NewEncoder(w).Encode(booksStruct)

}