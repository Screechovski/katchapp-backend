package handlers

import (
	"katchapp-backend/db"
	"katchapp-backend/middleware"
	"net/http"
	"strconv"
)

func TrainDelete(w http.ResponseWriter, r *http.Request) {
	userId, err := middleware.GetUserId(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Параметр 'id' обязателен", http.StatusBadRequest)
		return
	}

	trainId, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}

	err = db.DeleteTrain(userId, uint(trainId))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
