package handlers

import (
	"encoding/json"
	"fmt"
	"katchapp-backend/db"
	"net/http"
)

type Data struct {
	Token string `json:"token"`
}

type Response struct {
	IsValid bool `json:"isValid"`
}

func CheckTokenPost(w http.ResponseWriter, r *http.Request) {
	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		fmt.Println("error on read")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = db.GetUser(data.Token)

	w.Header().Set("Content-Type", "application/json")

	response := Response{IsValid: err == nil}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Println("error on get user")
		http.Error(w, "Error on response encode", http.StatusInternalServerError)
	}
}
