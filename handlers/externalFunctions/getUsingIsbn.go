package externalFunctions

import (

	"net/http"
)

func GetbookbySendingIsbn(isbn string) (*http.Response,error) {
	response,error:=http.Get("https://openlibrary.org/api/books?bibkeys=ISBN:"+isbn+"&format=json&jscmd=data")
	return response,error
	}