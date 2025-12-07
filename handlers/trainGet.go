package handlers

import (
	"encoding/json"
	"katchapp-backend/db"
	"katchapp-backend/middleware"
	"net/http"
)

func TrainGet(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserId(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	trains, err := db.GetTrainsByUserId(userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(trains); err != nil {
		http.Error(w, "Error on response encode", http.StatusInternalServerError)
	}
}
