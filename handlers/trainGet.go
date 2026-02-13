package handlers

import (
	"encoding/json"
	"katchapp-backend/db"
	"katchapp-backend/helper"
	"katchapp-backend/middleware"
	"log"
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
		helper.HandleError(w, err, http.StatusInternalServerError, "Failed to retrieve trains")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(trains); err != nil {
		log.Printf("Error encoding response: %v", err)
		return
	}
}
