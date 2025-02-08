package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"restapi/handlers/dto"
	"restapi/handlers/externalFunctions"
)

func GetBooksFromExternal(w http.ResponseWriter, req *http.Request) {
	isbn := req.URL.Query().Get("isbn")
	title := req.URL.Query().Get("title")
	author := req.URL.Query().Get("author")

	var data []byte
	if isbn != "" {
		response, err := externalFunctions.GetbookbySendingIsbn(isbn)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data, _ = io.ReadAll(response.Body)
	}

	if title != "" {
		response, err := externalFunctions.GetbookbySendingTitle(title)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data, _ = io.ReadAll(response.Body)
	}

	if author != "" {
		response, err := externalFunctions.GetbookbySendingAuther(author)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data, _ = io.ReadAll(response.Body)
	}

	var result = make(map[string]interface{})
	err := json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	
		docsArray := result["docs"]
		var DTO []dto.BooksForAutherDTO
		bytedocs, err := json.Marshal(docsArray)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.Unmarshal(bytedocs, &DTO)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(DTO)

	

}
