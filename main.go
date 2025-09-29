package main

import (
	"fmt"
	"katchapp-backend/db"
	"katchapp-backend/handlers"
	"katchapp-backend/middleware"
	"log"
	"net/http"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	db.Connect()
	defer db.Close()

	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/admin" || r.URL.Path == "/admin/" {
			http.ServeFile(w, r, "client/index.html")
			return
		}

		http.NotFound(w, r)
	})

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

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
