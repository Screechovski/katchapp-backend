package handlers

import (
	"encoding/json"
	"katchapp-backend/db"
	"katchapp-backend/helper"
	"log"
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
		helper.HandleError(w, err, http.StatusBadRequest, "Invalid request data")
		return
	}

	_, err = db.GetUser(data.Token)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := Response{IsValid: err == nil}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %v", err)
		return
	}
}
