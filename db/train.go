package db

import (
	"fmt"

	"gorm.io/gorm"
)

type Train struct {
	gorm.Model
	Date       string
	UserWeight float32
	UserID     uint
	Sets       []Sets `gorm:"foreignKey:TrainId;constraint:OnDelete:CASCADE"`
}

type TrainForUserSet struct {
	ID                int     `json:"id"`
	ExerciseId        int     `json:"exerciseId"`
	ExerciseImageName string  `json:"exerciseImageName"`
	ExerciseName      string  `json:"exerciseName"`
	Reps              int     `json:"reps"`
	Weight            float32 `json:"weight"`
}

type TrainForUserWithSets struct {
	ID         int               `json:"id"`
	Date       string            `json:"date"`
	UserWeight float32           `json:"userWeight"`
	Sets       []TrainForUserSet `json:"sets"`
}

func initTrain() {
	err := db.AutoMigrate(&Train{})
	if err != nil {
		panic("failed to migrate train table")
	}
}

func WriteTrain(date string, userId uint, userWeight float32) (uint, error) {
	train := Train{
		Date:       date,
		UserWeight: userWeight,
		UserID:     userId,
	}

	result := db.Create(&train)

	if result.Error != nil {
		return 0, result.Error
	}

	return train.ID, nil
}

func GetTrainsByUserId(userId uint) ([]Train, error) {
	var trains []Train

	err := db.Where("user_id = ?", userId).
		Order("date DESC").
		Preload("Sets.Exercise.MuscleGroup").
		Preload("Sets.Exercise.SecondaryMuscles.MuscleGroup").
		Find(&trains).Error

	return trains, err
}

func DeleteTrain(userId, trainId uint) error {
	result := db.Where("id = ? AND user_id = ?", trainId, userId).Delete(&Train{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("тренировка не найдена или не принадлежит пользователю")
	}

	return nil
}
