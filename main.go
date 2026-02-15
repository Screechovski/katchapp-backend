package main

import (
	"fmt"
	"katchapp-backend/db"
	"katchapp-backend/handlers"
	"katchapp-backend/middleware"
	"log"
	"net/http"
)

func main() {
	db.Connect()

	http.Handle(
		"/exercises",
		middleware.Cors(
			middleware.Method(
				middleware.MethodConfig{
					Get: handlers.ExercisesGet,
				},
			),
		),
	)

	http.HandleFunc(
		"/exercise/history/{exerciseId}",
		middleware.Cors(
			middleware.Auth(
				middleware.Method(
					middleware.MethodConfig{
						Get: handlers.ExerciseGet,
					},
				),
			),
		),
	)

	http.HandleFunc(
		"/train",
		middleware.Cors(
			middleware.Auth(
				middleware.Method(
					middleware.MethodConfig{
						Post:   handlers.TrainPost,
						Get:    handlers.TrainGet,
						Delete: handlers.TrainDelete,
					},
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

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
