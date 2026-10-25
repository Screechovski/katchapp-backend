package main

import (
	"fmt"
	"katchapp-backend/db"
	"katchapp-backend/handlers"
	"katchapp-backend/middleware"
	"log"
	"net/http"
	"os"
)

func main() {
	db.Connect()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ok":true}`))
	})

	http.Handle(
		"/exercises",
		middleware.Cors(
			middleware.Method(
				middleware.MethodConfig{
					Get: handlers.ExercisesGet,
					Post: middleware.Role(
						[]string{"admin"},
						handlers.ExercisesPost,
					),
				},
			),
		),
	)

	http.HandleFunc(
		"/exercises/{id}",
		middleware.Cors(
			middleware.Auth(
				middleware.Role(
					[]string{"admin"},
					middleware.Method(
						middleware.MethodConfig{
							Put:    handlers.ExercisePut,
							Delete: handlers.ExerciseDelete,
						},
					),
				),
			),
		),
	)

	http.HandleFunc(
		"/check-token",
		middleware.Cors(
			middleware.Method(
				middleware.MethodConfig{
					Post: handlers.CheckTokenPost,
				},
			),
		),
	)

	http.HandleFunc(
		"/muscle-groups",
		middleware.Cors(
			middleware.Auth(
				middleware.Method(
					middleware.MethodConfig{
						Get:  handlers.MuscleGroupsGet,
						Post: handlers.MuscleGroupsPost,
					},
				),
			),
		),
	)

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8088"
	}

	fmt.Printf("Server starting on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
