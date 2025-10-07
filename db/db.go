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

func initTables() {
	initUser()
	initTrain()
	initExercises()
	initSets()
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

	initTables()

	fmt.Println("Successfully connected to PostgreSQL!")
}

func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

func insert(queryRowFunc func() *sql.Row) (int, error) {
	var id int
	tx, err := db.Begin()

	if err != nil {
		return id, fmt.Errorf("failed to begin transaction: %w", err)
	}

	err = queryRowFunc().Scan(&id)

	if err != nil {
		return id, fmt.Errorf("error on inserting: %w", err)
	}

	defer tx.Rollback()

	return id, err
}
