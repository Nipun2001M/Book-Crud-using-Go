package externalFunctions

import (

	"net/http"
)

func GetbookbySendingAuther(Auther string) (*http.Response,error) {
	response,error:=http.Get("https://openlibrary.org/search.json?author="+Auther+"&sort=new")
	return response,error
	}