package main

import (
	"encoding/json"
	"fmt"
	"katchapp-backend/db"
	"log"
	"net/http"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	db.Connect()
	db.Init()
	defer db.Close()

	// http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method != http.MethodGet {
	// 		w.WriteHeader(http.StatusMethodNotAllowed)
	// 		return
	// 	}

	// 	w.Header().Set("Content-Type", "text/html")
	// 	http.ServeFile(w, r, "client/index.html")
	// })

	// Serve client assets under /admin/* and index.html at /admin
	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/admin" || r.URL.Path == "/admin/" {
			http.ServeFile(w, r, "client/index.html")
			return
		}
		http.NotFound(w, r)
	})
	http.Handle("/admin/",
		http.StripPrefix("/admin/",
			http.FileServer(http.Dir("client"))))

	// Serve images under /image/* from images directory
	http.Handle("/image/",
		http.StripPrefix("/image/",
			http.FileServer(http.Dir("images"))))

	http.HandleFunc("/api/exercises", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		exercises, err := db.GetAllExercises()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(exercises)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go and PostgreSQL 9999!")
	})

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
