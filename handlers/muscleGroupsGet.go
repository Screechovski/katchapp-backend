package handlers

import (
	"encoding/json"
	"katchapp-backend/db"
	"katchapp-backend/middleware"
	"net/http"
)

func MuscleGroupsGet(w http.ResponseWriter, r *http.Request) {
	_, err := middleware.GetUserId(r)

	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	groups, err := db.GetAllMuscleGroups()

	type ClientGroups struct {
		Id   uint   `json:"id"`
		Name string `json:"name"`
	}
	var clientGroups []ClientGroups

	for _, g := range groups {
		clientGroups = append(clientGroups, ClientGroups{g.ID, g.Name})
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(clientGroups); err != nil {
		http.Error(w, "Error on response encode", http.StatusInternalServerError)
	}
}
