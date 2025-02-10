package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"restapi/handlers/channel"
	"github.com/gorilla/mux"
)

func DeleteBook(w http.ResponseWriter,req *http.Request){
	params:=mux.Vars(req);
	id,error:=strconv.Atoi(params["id"])
	if error!=nil{
		http.Error(w,error.Error(),http.StatusBadRequest)

	}
	for index,value:=range AllBooks{
		if(value.BookID==id){
			AllBooks=append(AllBooks[:index],AllBooks[index+1:]...)
			json.NewEncoder(w).Encode(AllBooks)
			return
			

		}
	}
	response :=map[string]string{
		"message":"book not found",
	}
	msg:=fmt.Sprintf("/books - Book ID :%d",id)
	channel.AddToChannel("DELETE",msg)
	json.NewEncoder(w).Encode(response)

	

}