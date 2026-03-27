package handlers

import (
	"encoding/json"
	"katchapp-backend/db"
	"katchapp-backend/helper"
	"katchapp-backend/middleware"
	"log"
	"net/http"
	"strconv"
)

const (
	defaultTrainPageSize = 20
	maxTrainPageSize     = 100
)

func TrainGet(w http.ResponseWriter, r *http.Request) {
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

	limit := defaultTrainPageSize
	if l := r.URL.Query().Get("limit"); l != "" {
		if n, err := strconv.Atoi(l); err == nil && n > 0 {
			limit = n
		} else {
			http.Error(w, "invalid limit", http.StatusBadRequest)
			return
		}
	}
	if limit > maxTrainPageSize {
		limit = maxTrainPageSize
	}

	pageResult, err := db.GetTrainsByUserIdPaged(userID, page, limit)
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
