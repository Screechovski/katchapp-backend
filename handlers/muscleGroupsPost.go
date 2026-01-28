package handlers

import (
	"fmt"
	"katchapp-backend/db"
	"katchapp-backend/middleware"
	"net/http"
)

func MuscleGroupsPost(w http.ResponseWriter, r *http.Request) {
	_, err := middleware.GetUserId(r)

	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	groups, err := db.GetAllMuscleGroups()

	fmt.Println(groups)

	w.Header().Set("Content-Type", "application/json")
}
