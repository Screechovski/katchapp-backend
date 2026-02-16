package handlers

import (
	"encoding/json"
	"katchapp-backend/db"
	"katchapp-backend/helper"
	"katchapp-backend/middleware"
	"log"
	"net/http"
)

type UserTrain struct {
	Weight float32         `json:"weight"`
	Date   string          `json:"date"`
	Sets   []db.SetsParams `json:"sets"`
}

func TrainPost(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserId(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var data UserTrain
	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		helper.HandleError(w, err, http.StatusBadRequest, "Invalid request data")
		return
	}

	createdTrainId, err := db.WriteTrain(data.Date, userID, data.Weight)

	if err != nil {
		helper.HandleError(w, err, http.StatusInternalServerError, "Failed to create train")
		return
	}

	createdSetsIds, err := db.WriteSets(data.Sets, createdTrainId)

	if err != nil {
		helper.HandleError(w, err, http.StatusInternalServerError, "Failed to create sets")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(createdSetsIds); err != nil {
		log.Printf("Error encoding response: %v", err)
		return
	}
}
