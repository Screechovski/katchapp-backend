package db

import (
	"log"
	"os"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string
	Token  string `gorm:"unique"`
	Role   string
	Trains []Train
}

func initUser() error {
	if err := db.AutoMigrate(&User{}); err != nil {
		panic("failed to migrate train table")
	}

	firstToken, ok := os.LookupEnv("FIRST_USER_TOKEN")

	if !ok || firstToken == "" {
		log.Fatal("FIRST_USER_TOKEN env is required to seed initial user")
	}

	if _, err := GetUser(firstToken); err == nil {
		return nil
	}

	if err := CreateUser("admin", firstToken, "admin"); err != nil {
		return err
	}

	return nil
}

func GetUser(token string) (User, error) {
	var user User
	err := db.Where("token = ?", token).First(&user).Error
	return user, err
}

func CreateUser(name, token, role string) error {
	newUser := User{Name: name, Token: token, Role: role}

	result := db.Create(&newUser)

	return result.Error
}

type ShortNewSets struct {
	Reps       int     `json:"reps"`
	Weight     float32 `json:"weight"`
	TrainId    uint    `json:"trainId"`
	ExerciseId uint    `json:"exerciseId"`
	Date       string  `json:"date"`
	SetId      uint    `json:"id"`
}

func GetSets(userId, exerciseId int) ([]ShortNewSets, error) {
	var sets []ShortNewSets

	err := db.
		Table("sets").
		Select("sets.id, trains.date, sets.reps, sets.weight, sets.train_id, sets.exercise_id").
		Joins("INNER JOIN trains ON trains.id = sets.train_id").
		Joins("INNER JOIN users ON users.id = trains.user_id").
		Where("users.id = ? AND sets.exercise_id = ?", userId, exerciseId).
		Order("sets.id DESC").
		Scan(&sets).Error

	if sets == nil {
		return []ShortNewSets{}, err
	}

	return sets, err
}
