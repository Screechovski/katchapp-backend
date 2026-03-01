package handlers

import (
	"encoding/json"
	"katchapp-backend/db"
	"katchapp-backend/helper"
	"katchapp-backend/middleware"
	"net/http"
)

func MuscleGroupsPost(w http.ResponseWriter, r *http.Request) {
	_, err := middleware.GetUserId(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var body struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		helper.HandleError(w, err, http.StatusBadRequest, "Invalid request data")
		return
	}

	if body.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	id, err := db.CreateMuscleGroup(body.Name)
	if err != nil {
		helper.HandleError(w, err, http.StatusInternalServerError, "Failed to create muscle group")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]uint{"id": id}); err != nil {
		helper.HandleError(w, err, http.StatusInternalServerError, "Failed to write response")
		return
	}
}
