package main

import (
	"fmt"
	"katchapp-backend/db"
	"katchapp-backend/handlers"
	"katchapp-backend/helper"
	"katchapp-backend/middleware"
	"log"
	"net/http"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	db.Connect()
	defer db.Close()

	// http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.URL.Path == "/admin" || r.URL.Path == "/admin/" {
	// 		http.ServeFile(w, r, "client/index.html")
	// 		return
	// 	}

	// 	http.NotFound(w, r)
	// })

	http.Handle(
		"/admin/",
		http.StripPrefix(
			"/admin/",
			http.FileServer(
				http.Dir("client"),
			),
		),
	)

	http.Handle(
		"/image/",
		http.StripPrefix(
			"/image/",
			http.FileServer(
				http.Dir("images"),
			),
		),
	)

	http.Handle(
		"/api/exercises",
		middleware.Cors(
			middleware.Method(
				middleware.MethodConfig{
					Get: handlers.ExercisesGet,
				},
			),
		),
	)

	http.HandleFunc(
		"/api/train",
		middleware.Cors(
			middleware.Auth(
				middleware.Method(
					middleware.MethodConfig{
						Post: handlers.TrainPost,
						Get:  handlers.TrainGet,
					},
				),
			),
		),
	)

	http.HandleFunc(
		"/api/check-token",
		middleware.Cors(
			middleware.Method(
				middleware.MethodConfig{
					Post: handlers.CheckTokenPost,
				},
			),
		),
	)

	// Hello World endpoint
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method != http.MethodGet {
	// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 		return
	// 	}
	// 	w.Header().Set("Content-Type", "text/plain")
	// 	fmt.Fprint(w, "Hello World")
	// })

	if !helper.IsDev() {
		http.Handle(
			"/",
			http.FileServer(
				http.Dir("../katch-app"),
			),
		)
	}

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
