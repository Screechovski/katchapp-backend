package handlers

import (
	"encoding/json"
	"fmt"
	"katchapp-backend/db"
	"katchapp-backend/middleware"
	"net/http"
)

type UserTrain struct {
	Weight float32         `json:"weight"`
	Date   string          `json:"date"`
	Sets   []db.SetsParams `json:"sets"`
}

func TrainPost(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserId(r)

	var data UserTrain
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		fmt.Println("1 error")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdTrainId, err := db.WriteTrain(db.TrainWithoutId{
		Date:       data.Date,
		UserId:     userID,
		UserWeight: data.Weight,
	})

	if err != nil {
		fmt.Println("2 error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	createdSetsIds, err := db.WriteSets(data.Sets, createdTrainId)

	if err != nil {
		fmt.Println("3 error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(createdSetsIds); err != nil {
		fmt.Println("4 error")
		http.Error(w, "Error on response encode", http.StatusInternalServerError)
	}
}
