package externalFunctions

import (
	"net/http"
	"strings"
)

func GetbookbySendingTitle(title string) (*http.Response,error) {
	sendingTitle:=strings.ReplaceAll(title," ","+")
	response,error:=http.Get("https://openlibrary.org/search.json?title="+sendingTitle)
	return response,error
	

}