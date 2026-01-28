package middleware

import (
	"katchapp-backend/db"
	"net/http"
	"slices"
)

func Role(validRoles []string, next http.HandlerFunc) http.HandlerFunc {
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

		found := slices.Contains(validRoles, user.Role)

		if found {
			next(w, r)
		} else {
			http.Error(w, "no no mr fish", http.StatusUnauthorized)
			return
		}
	}
}
