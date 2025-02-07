package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"restapi/handlers/externalFunctions"
)

func GetBooksFromExternal(w http.ResponseWriter,req *http.Request) {
	isbn:=req.URL.Query().Get("isbn")
	title:=req.URL.Query().Get("title")
	author:=req.URL.Query().Get("author")

	var data []byte
	if isbn!=""{
		response,error:=externalFunctions.GetbookbySendingIsbn(isbn)
	if error!=nil {
			http.Error(w,error.Error(),http.StatusBadRequest)
		}else{
			fmt.Println(response)
			data,_=io.ReadAll(response.Body)
			

		}

	}
	if title!=""{
		response,error:=externalFunctions.GetbookbySendingTitle(title)
		if error!=nil {
			http.Error(w,error.Error(),http.StatusBadRequest)
		}else{
			data,_=io.ReadAll(response.Body)
		}

	}
	if author!=""{
		response,error:=externalFunctions.GetbookbySendingAuther(author)
		if error!=nil {
			http.Error(w,error.Error(),http.StatusBadRequest)
		}else{
			data,_=io.ReadAll(response.Body)
	}

	}
	var result map[string]interface{}
			error:=json.Unmarshal(data,&result)
			if error!=nil{
				http.Error(w,error.Error(),http.StatusBadRequest)


			}else{
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(result)
			}



	







}