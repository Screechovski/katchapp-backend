package handlers

import (
	"encoding/json"
	"fmt"
	"katchapp-backend/db"
	"katchapp-backend/middleware"
	"net/http"
	"strconv"
)

func ExerciseGet(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserId(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	exerciseId, err := strconv.Atoi(r.PathValue("exerciseId"))

	if err != nil {
		http.Error(w, "decode error", http.StatusBadRequest)
		return
	}

	sets, err := db.GetTopSets(int(userID), exerciseId)

	type SubRes struct {
		Weight int `json:"weight"`
		Reps   int `json:"reps"`
	}
	type Res struct {
		Top  SubRes `json:"top"`
		Last SubRes `json:"last"`
	}
	result := Res{
		Top: SubRes{
			Weight: -1,
			Reps:   -1,
		},
		Last: SubRes{
			Weight: -1,
			Reps:   -1,
		},
	}

	fmt.Println(sets)

	for _, set := range sets {
		if int(set.Weight) > result.Top.Weight {
			result.Top.Weight = int(set.Weight)
			result.Top.Reps = set.Reps
		}
	}

	if len(sets) > 0 {
		result.Last.Weight = int(sets[0].Weight)
		result.Last.Reps = sets[0].Reps
	}

	if err != nil {
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "error encode response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
