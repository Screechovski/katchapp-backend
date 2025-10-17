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

func initUser() {
	err := db.AutoMigrate(&User{})

	if err != nil {
		panic("failed to migrate train table")
	}

	firstToken, ok := os.LookupEnv("FIRST_USER_TOKEN")

	if !ok || firstToken == "" {
		log.Fatal("FIRST_USER_TOKEN env is required to seed initial user")
	}

	_, err = GetUser(firstToken)

	if err == nil {
		return
	}

	err = CreateUser("admin", firstToken, "admin")

	if err != nil {
		log.Panic("failed creating first user", err.Error())
	}
}

func GetUser(token string) (User, error) {
	var user User

	result := db.Where("token = ?", token).First(&user)

	return user, result.Error
}

func CreateUser(name, token, role string) error {
	newUser := User{Name: name, Token: token, Role: role}

	result := db.Create(&newUser)

	return result.Error
}
