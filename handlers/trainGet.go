package handlers

import (
	"encoding/json"
	"fmt"
	"katchapp-backend/db"
	"katchapp-backend/middleware"
	"net/http"
)

func TrainGet(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserId(r)
	trains, err := db.GetTrainsByUserId(userID)

	if err != nil {
		fmt.Println("1 error TrainGet", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(trains); err != nil {
		fmt.Println("2 error TrainGet")
		http.Error(w, "Error on response encode", http.StatusInternalServerError)
	}
}
