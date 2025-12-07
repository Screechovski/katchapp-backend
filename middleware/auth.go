package middleware

import (
	"context"
	"fmt"
	"katchapp-backend/db"
	"net/http"
)

type contextKey string

var authUserId contextKey = "userID"

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			return
		}

		authToken := r.Header.Get("authorization")

		if authToken == "" {
			http.Error(w, "Authorization token required", http.StatusUnauthorized)
			return
		}

		user, err := db.GetUser(authToken)

		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), authUserId, user.ID)

		next(w, r.WithContext(ctx))
	}
}

func GetUserId(r *http.Request) (uint, error) {
	value := r.Context().Value(authUserId)
	if value == nil {
		return 0, fmt.Errorf("user ID not found in context")
	}
	userId, ok := value.(uint)
	if !ok {
		return 0, fmt.Errorf("invalid user ID type")
	}
	return userId, nil
}
