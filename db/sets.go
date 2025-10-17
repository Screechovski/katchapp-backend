package db

import (
	"gorm.io/gorm"
)

type SetsParams struct {
	ExerciseId int     `json:"exerciseId"`
	Reps       int     `json:"reps"`
	Weight     float32 `json:"weight"`
}

type Sets struct {
	gorm.Model
	Reps       int     `json:"reps"`
	Weight     float32 `json:"weight"`
	TrainId    uint    `json:"trainId"`
	ExerciseId uint    `json:"exerciseId"`
	Exercise   Exercise
}

func initSets() {
	err := db.AutoMigrate(&Sets{})
	if err != nil {
		panic("failed to migrate sets table")
	}
}

func WriteSets(sets []SetsParams, trainId uint) ([]uint, error) {
	var err error
	var ids []uint

	for i := range sets {
		id, err := WriteSet(sets[i], trainId)

		if err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	return ids, err
}

func WriteSet(set SetsParams, trainId uint) (uint, error) {
	newSet := Sets{
		Reps:       set.Reps,
		Weight:     set.Weight,
		TrainId:    trainId,
		ExerciseId: uint(set.ExerciseId),
	}

	result := db.Create(&newSet)
	if result.Error != nil {
		return 0, result.Error
	}

	return newSet.ID, nil
}
