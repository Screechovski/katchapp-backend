package db

import (
	"gorm.io/gorm"
)

type Exercise struct {
	gorm.Model
	Name             string `json:"name"`
	ImageName        string `json:"imageName"`
	Hidden           bool   `json:"hidden"`
	MuscleGroupID    uint
	MuscleGroup      MuscleGroup
	SecondaryMuscles []ExerciseSecondaryMuscle `gorm:"foreignKey:ExerciseId"`
}

type SubGroup struct {
	Id    uint
	Level int
}

func UpdateImage(id uint, imageName string) error {
	result := db.Model(&Exercise{}).Where("id = ?", id).Update("ImageName", imageName)
	return result.Error
}

func SaveExercise(name, imgPath string, group uint, subGroups []SubGroup) (uint, error) {
	exercise := Exercise{
		ImageName:        imgPath,
		Name:             name,
		Hidden:           false,
		MuscleGroupID:    group,
		SecondaryMuscles: []ExerciseSecondaryMuscle{},
	}

	for _, sub := range subGroups {
		exercise.SecondaryMuscles = append(exercise.SecondaryMuscles, ExerciseSecondaryMuscle{
			MuscleGroupId:   sub.Id,
			EngagementLevel: sub.Level,
		})
	}

	result := db.Create(&exercise)

	return exercise.ID, result.Error
}

func GetAllExercises() ([]Exercise, error) {
	var exercises []Exercise

	result := db.Preload("SecondaryMuscles").Find(&exercises)

	return exercises, result.Error
}

func GetExerciseByID(id uint) (Exercise, error) {
	var exercise Exercise

	result := db.Preload("SecondaryMuscles").First(&exercise, id)

	return exercise, result.Error
}

func UpdateExerciseHidden(id uint, hidden bool) error {
	result := db.Model(&Exercise{}).Where("id = ?", id).Update("hidden", hidden)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func DeleteExercise(id uint) error {
	result := db.Where("id = ?", id).Delete(&Exercise{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
