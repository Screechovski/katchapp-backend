package main

import (
	"fmt"
	"katchapp-backend/db"
	"katchapp-backend/handlers"
	"katchapp-backend/helper"
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
					Post: middleware.Role(
						[]string{"admin"},
						handlers.ExercisesPost,
					),
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

	http.HandleFunc(
		"/muscle-groups",
		middleware.Cors(
			middleware.Auth(
				middleware.Method(
					middleware.MethodConfig{
						Get: handlers.MuscleGroupsGet,
					},
				),
			),
		),
	)

	if helper.IsDev() {
		http.Handle(
			"/admin/",
			http.StripPrefix(
				"/admin/",
				http.FileServer(
					http.Dir("./client"),
				),
			),
		)
	}

	if helper.IsDev() {
		http.Handle(
			"/image/",
			http.StripPrefix(
				"/image/",
				http.FileServer(
					http.Dir("./images"),
				),
			),
		)
	}

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
