package db

import (
	"gorm.io/gorm"
)

type CollectedSetsParams struct {
	Id     int     `json:"id"`
	Reps   int     `json:"reps"`
	Weight float32 `json:"weight"`
	Sets   int     `json:"sets"`
}

type CollectedSets struct {
	gorm.Model

	Reps     *int     `json:"reps"`
	Weight   *float32 `json:"weight"`
	Sets     *int     `json:"sets"`
	Speed    *float32 `json:"speed,omitempty"`
	Incline  *int     `json:"incline,omitempty"`
	Duration *int     `json:"duration,omitempty"` // Время лучше хранить в секундах или минутах
	Distance *float32 `json:"distance,omitempty"`
	Calories *int     `json:"calories,omitempty"`
	Time     *int     `json:"time,omitempty"`

	TrainId uint `json:"trainId"`

	ExerciseId uint `json:"exerciseId"`
	Exercise   Exercise
}

func initCollectedSets() {
	err := db.AutoMigrate(&CollectedSets{})
	if err != nil {
		panic("failed to migrate sets table")
	}
}
