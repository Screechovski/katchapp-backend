package db

import (
	"fmt"
	"log"
	"os"
)

type User struct {
	ID    int
	Token string
}

func initUser() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			token TEXT UNIQUE NOT NULL
		);
	`)

	if err != nil {
		log.Fatal(fmt.Errorf("failed creating user table: %w", err))
	}

	var existingUsers int
	err = db.QueryRow("SELECT COUNT(*) FROM users").Scan(&existingUsers)

	if err != nil {
		log.Fatal(fmt.Errorf("failed seelcting users: %w", err))
	}

	if existingUsers == 0 {
		firstToken, ok := os.LookupEnv("FIRST_USER_TOKEN")

		if !ok || firstToken == "" {
			log.Fatal("FIRST_USER_TOKEN env is required to seed initial user")
		}

		// Insert initial user
		_, err = db.Exec("INSERT INTO users (token) VALUES ($1)", firstToken)

		if err != nil {
			log.Fatal(fmt.Errorf("failed to insert initial user: %w", err))
		}

		fmt.Println("Seeded initial user")
	}
}

func GetUser(token string) (User, error) {
	var user User
	err := db.QueryRow("SELECT id, token FROM users WHERE token = $1", token).Scan(&user.ID, &user.Token)
	return user, err
}
