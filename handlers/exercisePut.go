package handlers

import (
	"encoding/json"
	"errors"
	"katchapp-backend/db"
	"katchapp-backend/helper"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func ExercisePut(w http.ResponseWriter, r *http.Request) {
	exerciseID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}

	var body struct {
		Hidden *bool `json:"hidden"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		helper.HandleError(w, err, http.StatusBadRequest, "Invalid request data")
		return
	}

	if body.Hidden == nil {
		http.Error(w, "hidden is required", http.StatusBadRequest)
		return
	}

	err = db.UpdateExerciseHidden(uint(exerciseID), *body.Hidden)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "exercise not found", http.StatusNotFound)
			return
		}
		helper.HandleError(w, err, http.StatusInternalServerError, "Failed to update exercise")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
