package handlers

import (
	"encoding/json"
	"net/http"
)



func GetAllBooks(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)	
	json.NewEncoder(w).Encode(AllBooks)

}