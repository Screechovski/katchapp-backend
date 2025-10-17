package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func initTables() {
	initUser()
	initTrain()
	initMuscleGroup()
	initExerciseSecondaryMuscles()
	initExercises()
	initSets()
}

func Connect() {
	err := godotenv.Load()

	if err != nil {
		log.Print("No .env file found")
	}

	user, userExists := os.LookupEnv("DB_USER")
	password, passwordExists := os.LookupEnv("DB_PASSWORD")
	dbName, dbNameExists := os.LookupEnv("DB_NAME")
	host, hostExists := os.LookupEnv("DB_HOST")
	port, portExists := os.LookupEnv("DB_PORT")

	if !userExists || !passwordExists || !dbNameExists || !hostExists || !portExists {
		log.Fatal("no data for connect to DB")
	}

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		user,
		password,
		dbName,
		host,
		port,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panicf("Failed to connect to database: %v", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")

	initTables()
}
