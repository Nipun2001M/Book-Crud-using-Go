package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restapi/handlers/channel"
	"restapi/redis"
	"strconv"
)

func GetAllBooks(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	var BookKeys = []string{}
	Books := []Book{}
	var cursor uint64

	for {
		key, nextCursor, err := redis.Client.Scan(redis.Ctx, cursor, "Book:*", 0).Result()
		if err != nil {
			fmt.Println("error in getting keys", err)
		}
		BookKeys = append(BookKeys, key...)
		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}
	for _, key := range BookKeys {
		fmt.Println("key", key)
		bookData, err := redis.Client.HGetAll(redis.Ctx, key).Result()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Bd", bookData)
		id := bookData["bookId"]
		bookid, err1 := strconv.Atoi(id)
		if err1 != nil {
			fmt.Println("error in id onvertion to int", err)
		}
		book := Book{
			BookID:   bookid,
			BookName: bookData["bookName"],
			Author:   bookData["author"],
		}
		Books = append(Books, book)
	}
	channel.AddToChannel("GET", "/books - All Books Taken")
	fmt.Println(Books)
	json.NewEncoder(w).Encode(Books)

}
