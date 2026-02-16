package handlers

import (
	"encoding/json"
	"fmt"
	"katchapp-backend/db"
	"katchapp-backend/helper"
	"katchapp-backend/middleware"
	"log"
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

type Result struct {
	ExerciseID int           `json:"exerciseId"`
	Rm         float32       `json:"rm"`
	Trains     []TrainCombat `json:"trains"`
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
func getRm(sets []db.ShortNewSets) float32 {
	rms := []float32{}
	for _, set := range sets {
		rms = append(rms, set.Weight/(1.0278-(0.0278*float32(set.Reps))))
	}
	if len(rms) == 0 {
		return 0
	}
	var sum float32
	for _, v := range rms {
		sum += v
	}
	return sum / float32(len(rms))
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
		helper.HandleError(w, err, http.StatusInternalServerError, "Failed to retrieve sets")
		return
	}

	var rm float32
	if len(sets) < 6 {
		rm = 0
	} else {
		rm = getRm(sets[:min(6, len(sets))])
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
		return trains[i].TrainID > trains[j].TrainID
	})

	result := Result{
		Trains:     trains,
		ExerciseID: exerciseId,
		Rm:         rm,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Printf("Error encoding response: %v", err)
		return
	}
}
