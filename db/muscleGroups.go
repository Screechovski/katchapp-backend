package db

import (
	"gorm.io/gorm"
)

type MuscleGroup struct {
	gorm.Model
	Name string
}

func GetAllMuscleGroups() ([]MuscleGroup, error) {
	var groups []MuscleGroup
	result := db.Find(&groups)
	return groups, result.Error
}

func CreateMuscleGroup(name string) (uint, error) {
	group := &MuscleGroup{Name: name}
	result := db.Create(group)
	return group.ID, result.Error
}
