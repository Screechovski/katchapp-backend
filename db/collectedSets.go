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

	Reps   int     `json:"reps"`
	Weight float32 `json:"weight"`
	Sets   int     `json:"sets"`

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

func WriteCollectedSets(sets []CollectedSetsParams, trainId uint) error {
	var newSets []CollectedSets

	for _, set := range sets {
		newSets = append(newSets, CollectedSets{
			Reps:       set.Reps,
			Weight:     set.Weight,
			Sets:       set.Sets,
			TrainId:    trainId,
			ExerciseId: uint(set.Id),
		})
	}

	result := db.Create(&newSets)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func WriteCollectedSet(set CollectedSetsParams, trainId uint) (uint, error) {
	newSet := CollectedSets{
		Reps:       set.Reps,
		Weight:     set.Weight,
		Sets:       set.Sets,
		TrainId:    trainId,
		ExerciseId: uint(set.Id),
	}
	result := db.Create(&newSet)
	if result.Error != nil {
		return 0, result.Error
	}
	return newSet.ID, nil
}
