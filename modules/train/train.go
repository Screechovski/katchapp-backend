package trainModule

import (
	"encoding/json"
	"katchapp-backend/db"
	"katchapp-backend/helper"
	"katchapp-backend/middleware"
	"log"
	"net/http"
	"strconv"
)

type UserTrain struct {
	Weight    float32                  `json:"weight"`
	Date      string                   `json:"date"`
	Exercises []db.CollectedSetsParams `json:"exercises"`
}

func Post(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserId(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var data UserTrain
	err = json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		helper.HandleError(w, err, http.StatusBadRequest, "Invalid request data")
		return
	}

	createdTrainId, err := db.WriteTrain(data.Date, userID, data.Weight)

	if err != nil {
		helper.HandleError(w, err, http.StatusInternalServerError, "Failed to create train")
		return
	}

	err = db.WriteCollectedSets(data.Exercises, createdTrainId)

	if err != nil {
		helper.HandleError(w, err, http.StatusInternalServerError, "Failed to create collected sets")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(createdTrainId); err != nil {
		log.Printf("Error encoding response: %v", err)
		return
	}
}

// type TrainSecodaryMuscleGroup struct {
// 	id              uint   `json:"id"`
// 	name            string `json:"name"`
// 	engagementLevel int    `json:"engagementLevel"`
// }

// type TrainMuscleGroup struct {
// 	id   uint   `json:"id"`
// 	name string `json:"name"`
// }
// type TrainExerciseSet struct {
// 	id     uint    `json:"id"`
// 	reps   int     `json:"reps"`
// 	weight float32 `json:"weight"`
// 	sets   int     `json:"sets"`
// }
// type TrainExercise struct {
// 	id                    uint                       `json:"id"`
// 	muscleGroup           TrainMuscleGroup           `json:"muscleGroup"`
// 	secondaryMuscleGroups []TrainSecodaryMuscleGroup `json:"secondaryMuscleGroups"`
// 	sets                  []TrainExerciseSet         `json:"sets"`
// 	imageName             string                     `json:"imageName"`
// 	name                  string                     `json:"name"`
// }

// type TrainForUser struct {
// 	id         uint            `json:"id"`
// 	date       string          `json:"date"`
// 	userWeight float32         `json:"userWeight"`
// 	exercises  []TrainExercise `json:"exercises"`
// }

// type TrainsForUser struct {
// 	items []TrainForUser
// 	page  int   `json:"page"`
// 	limit int   `json:"limit"`
// 	total int64 `json:"total"`
// }

func Get(w http.ResponseWriter, r *http.Request) {
	userID, err := middleware.GetUserId(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	page := 1
	if p := r.URL.Query().Get("page"); p != "" {
		if n, err := strconv.Atoi(p); err == nil && n > 0 {
			page = n
		} else {
			http.Error(w, "invalid page", http.StatusBadRequest)
			return
		}
	}
	limit := 20
	pageResult, err := db.GetTrainsWithCollectedSetsPaged(userID, page, limit)
	if err != nil {
		helper.HandleError(w, err, http.StatusInternalServerError, "Failed to retrieve trains")
		return
	}

	type response struct {
		Items []db.Train `json:"items"`
		Page  int        `json:"page"`
		Limit int        `json:"limit"`
		Total int64      `json:"total"`
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response{
		Items: pageResult.Items,
		Page:  page,
		Limit: limit,
		Total: pageResult.Total,
	}); err != nil {
		log.Printf("Error encoding response: %v", err)
		return
	}
}
