package db

import (
	"gorm.io/gorm"
)

type ExerciseSecondaryMuscle struct {
	gorm.Model
	ExerciseId      uint        `json:"exerciseId"`
	MuscleGroupId   uint        `json:"muscleGroupId"`
	MuscleGroup     MuscleGroup `gorm:"foreignKey:MuscleGroupId"`
	EngagementLevel int         `json:"engagementLevel"` // 1-5 уровень вовлеченности
}

func initExerciseSecondaryMuscles() {
	err := db.AutoMigrate(&ExerciseSecondaryMuscle{})

	if err != nil {
		panic("failed to migrate exercise_secondary_muscles table")
	}
}
