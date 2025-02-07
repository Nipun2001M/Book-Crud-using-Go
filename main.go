package main

import (
	"fmt"
	"log"
	"net/http"

	"restapi/handlers"

	"github.com/gorilla/mux"
)







func main() {
	router:=mux.NewRouter()
	router.HandleFunc("/books",handlers.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}",handlers.GetSpecificBook).Methods("GET")
	router.HandleFunc("/books",handlers.PostBook).Methods("POST")
	router.HandleFunc("/books/{id}",handlers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}",handlers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/externalbooks",handlers.GetBooksFromExternal).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080",router))
	fmt.Println("server runing....")


	
}