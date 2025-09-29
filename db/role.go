package db

import (
	"fmt"
	"log"
)

type Role struct {
	ID   int
	Slug string
}

func initRole() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS roles (
			id SERIAL PRIMARY KEY,
			slug TEXT UNIQUE NOT NULL
		);
	`)

	if err != nil {
		log.Fatal(fmt.Errorf("failed creating role table: %w", err))
	}

	var existingRoles int
	err = db.QueryRow("SELECT COUNT(*) FROM roles").Scan(&existingRoles)

	if err != nil {
		log.Fatal(fmt.Errorf("failed seelcting role: %w", err))
	}

	if existingRoles == 0 {

	}
}
