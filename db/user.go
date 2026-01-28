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

type ShortSets struct {
	Reps       int     `json:"reps"`
	Weight     float32 `json:"weight"`
	TrainId    uint    `json:"trainId"`
	ExerciseId uint    `json:"exerciseId"`
	SetId      uint    `json:"id"`
}

func GetTopSets(userId, exerciseId int) ([]ShortSets, error) {
	var sets []ShortSets

	err := db.
		Table("users u").
		Joins("INNER JOIN trains t ON u.id = t.user_id").
		Joins("INNER JOIN sets s ON s.train_id = t.id").
		Where("u.id = ? AND s.exercise_id = ?", userId, exerciseId).
		Select("s.reps, s.weight, s.train_id, s.exercise_id, s.id as set_id").
		Order("s.created_at DESC").
		Scan(&sets).Error

	if sets == nil {
		return []ShortSets{}, err
	}

	return sets, err
}
