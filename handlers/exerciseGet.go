package handlers

import (
	"encoding/json"
	"fmt"
	"katchapp-backend/db"
	"katchapp-backend/middleware"
	"net/http"
	"sort"
	"strconv"
)

type RawRecord struct {
	Reps       int     `json:"reps"`
	Weight     float32 `json:"weight"`
	TrainID    uint    `json:"trainId"`
	ExerciseID uint    `json:"exerciseId"`
	Date       string  `json:"date"`
	ID         int     `json:"id"`
}

type TrainCombatApproaches struct {
	Reps   int     `json:"reps"`
	Weight float32 `json:"weight"`
	Sets   int     `json:"sets"`
}

type TrainCombat struct {
	TrainID    uint                    `json:"trainId"`
	TrainDate  string                  `json:"trainDate"`
	Approaches []TrainCombatApproaches `json:"approaches"`
}

func splitApproachKey(key string) (int, float32) {
	parts := make([]float32, 2)
	for i, s := range split(key, '_') {
		if val, err := strconv.ParseFloat(s, 32); err == nil {
			parts[i] = float32(val)
		}
	}
	return int(parts[0]), parts[1]
}

func split(s string, sep byte) []string {
	var result []string
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == sep {
			result = append(result, s[start:i])
			start = i + 1
		}
	}
	result = append(result, s[start:])
	return result
}

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

	sets, err := db.GetSets(int(userID), exerciseId)
	if err != nil {
		http.Error(w, "getting sets error", http.StatusBadRequest)
		return
	}

	trainMap := make(map[uint]*TrainCombat)
	for _, set := range sets {
		if _, exists := trainMap[set.TrainId]; !exists {
			trainMap[set.TrainId] = &TrainCombat{
				TrainID:    set.TrainId,
				TrainDate:  set.Date,
				Approaches: []TrainCombatApproaches{},
			}
		}
	}

	approachGroups := make(map[uint]map[string]int)
	for _, set := range sets {
		approachKey := fmt.Sprintf("%v_%v", set.Reps, set.Weight)
		if _, exists := approachGroups[set.TrainId]; !exists {
			approachGroups[set.TrainId] = make(map[string]int)
		}
		approachGroups[set.TrainId][approachKey]++
	}

	var trains []TrainCombat
	for trainKey, train := range trainMap {
		groups := approachGroups[trainKey]

		for approachKey, sets := range groups {
			reps, weight := splitApproachKey(approachKey)
			train.Approaches = append(train.Approaches, TrainCombatApproaches{
				Reps:   reps,
				Weight: weight,
				Sets:   sets,
			})
		}

		sort.Slice(train.Approaches, func(i, j int) bool {
			if train.Approaches[i].Weight != train.Approaches[j].Weight {
				return train.Approaches[i].Weight < train.Approaches[j].Weight
			}
			return train.Approaches[i].Reps < train.Approaches[j].Reps
		})

		trains = append(trains, *train)
	}

	sort.Slice(trains, func(i, j int) bool {
		return trains[i].TrainDate < trains[j].TrainDate
	})

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(trains); err != nil {
		http.Error(w, "error encode response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
