package handlers

import (
	"encoding/json"
	"katchapp-backend/db"
	"katchapp-backend/helper"
	"log"
	"net/http"
)

func ExercisesGet(w http.ResponseWriter, r *http.Request) {
	exercises, err := db.GetAllExercises()

	if err != nil {
		helper.HandleError(w, err, http.StatusInternalServerError, "Failed to retrieve exercises")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(exercises); err != nil {
		log.Printf("Error encoding response: %v", err)
		return
	}
}
