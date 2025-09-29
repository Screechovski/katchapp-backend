package handlers

import (
	"encoding/json"
	"katchapp-backend/db"
	"net/http"
)

func ExercisesGet(w http.ResponseWriter, r *http.Request) {
	exercises, err := db.GetAllExercises()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(exercises); err != nil {
		http.Error(w, "Error on response encode", http.StatusInternalServerError)
	}
}
