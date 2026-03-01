package handlers

import (
	"errors"
	"katchapp-backend/db"
	"katchapp-backend/helper"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"gorm.io/gorm"
)

func ExerciseDelete(w http.ResponseWriter, r *http.Request) {
	exerciseID, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}

	exercise, err := db.GetExerciseByID(uint(exerciseID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "exercise not found", http.StatusNotFound)
			return
		}
		helper.HandleError(w, err, http.StatusInternalServerError, "Failed to load exercise")
		return
	}

	err = db.DeleteExercise(uint(exerciseID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "exercise not found", http.StatusNotFound)
			return
		}
		helper.HandleError(w, err, http.StatusInternalServerError, "Failed to delete exercise")
		return
	}

	if exercise.ImageName != "" {
		imagePath := filepath.Join("images/exercise", filepath.Base(exercise.ImageName))
		if err := os.Remove(imagePath); err != nil && !errors.Is(err, os.ErrNotExist) {
			helper.HandleError(w, err, http.StatusInternalServerError, "Failed to delete exercise image")
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
