package main

import (
	"fmt"
	"log"
	"net/http"
	"restapi/handlers/channel"
	"restapi/handlers"
	"restapi/redis"
	"github.com/gorilla/mux"
)








func main() {
	go channel.Listener()
	redis.RedisConnection()
	router:=mux.NewRouter()
	router.HandleFunc("/books",handlers.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}",handlers.GetSpecificBook).Methods("GET")
	router.HandleFunc("/books",handlers.PostBook).Methods("POST")
	router.HandleFunc("/books/{id}",handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}",handlers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/externalbooks",handlers.GetBooksFromExternal).Methods("GET")
	fmt.Println("server runing....")
	log.Fatal(http.ListenAndServe(":8080",router))


	
}