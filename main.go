package main

import (
	"fmt"
	"log"
	"net/http"
	"restapi/handlers"
	"restapi/handlers/channel"
	"restapi/redis"
	"sync"

	"github.com/gorilla/mux"
)








func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go channel.Listener(&wg)
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
	wg.Wait()


	
}