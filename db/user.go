package db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Token string `gorm:"unique"`
	Role  string
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
