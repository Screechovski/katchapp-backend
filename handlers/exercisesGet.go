package handlers

import (
	"encoding/json"
	"katchapp-backend/db"
	"net/http"
	"sort"
)

func ExercisesGet(w http.ResponseWriter, r *http.Request) {
	exercises, err := db.GetAllExercises()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	authToken := r.Header.Get("authorization")

	if authToken != "" {
		user, err := db.GetUser(authToken)

		if err == nil {
			trains, err := db.GetTrainsByUserId(user.ID)

			if err == nil {
				frequency := make(map[uint]int)

				for _, train := range trains {
					for _, set := range train.Sets {
						if set.ExerciseId != 0 {
							frequency[set.ExerciseId]++
						}
					}
				}

				if len(frequency) > 0 {
					sort.Slice(exercises, func(i, j int) bool {
						return frequency[exercises[i].ID] > frequency[exercises[j].ID]
					})
				}
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(exercises); err != nil {
		http.Error(w, "Error on response encode", http.StatusInternalServerError)
	}
}
