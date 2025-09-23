package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var db *sql.DB

func Init() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			token TEXT UNIQUE NOT NULL
		);
	`)
	if err != nil {
		log.Fatal(fmt.Errorf("failed creating user table: %w", err))
	}

	InitExercises()

	// Seed first user if none exists
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

func Connect() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	user, userExists := os.LookupEnv("DB_USER")
	password, passwordExists := os.LookupEnv("DB_PASSWORD")
	dbName, dbNameExists := os.LookupEnv("DB_NAME")
	host, hostExists := os.LookupEnv("DB_HOST")
	port, portExists := os.LookupEnv("DB_PORT")

	if !userExists || !passwordExists || !dbNameExists || !hostExists || !portExists {
		log.Fatal("no data for connect, user: ", user, " password: ", password, " dbName: ", dbName)
	}
	portInt, err := strconv.Atoi(port)
	if err != nil {
		log.Fatal("failed to convert port to int: ", err)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, portInt, user, password, dbName)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to PostgreSQL!")
}

func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

func GetDB() *sql.DB {
	return db
}
