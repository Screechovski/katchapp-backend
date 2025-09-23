package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	user, userExists := os.LookupEnv("POSTGRES_USER")
	password, passwordExists := os.LookupEnv("POSTGRES_PASSWORD")
	dbName, dbNameExists := os.LookupEnv("POSTGRES_DB")

	if !userExists || !passwordExists || !dbNameExists {
		log.Fatal("no data for connect")
	}

	connStr := fmt.Sprintf("postgresql://%s:%s@db:5432/%s?sslmode=disable", user, password, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to PostgreSQL!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go and PostgreSQL!")
	})

	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
