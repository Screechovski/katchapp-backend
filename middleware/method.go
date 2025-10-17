package middleware

import (
	"fmt"
	"net/http"
)

type MethodConfig struct {
	Get    http.HandlerFunc
	Post   http.HandlerFunc
	Delete http.HandlerFunc
}

func Method(config MethodConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Method middleware:", r.Method)

		switch r.Method {
		case http.MethodOptions:
			return
		case http.MethodPost:
			config.Post(w, r)
			return
		case http.MethodDelete:
			config.Delete(w, r)
			return
		case http.MethodGet:
			config.Get(w, r)
			return
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}
