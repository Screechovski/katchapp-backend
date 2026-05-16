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

type TrainsPage struct {
	Items []Train
	Total int64
}

func GetTrainsByUserIdPaged(userId uint, page, pageSize int) (TrainsPage, error) {
	var out TrainsPage
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}

	q := db.Model(&Train{}).Where("user_id = ?", userId)
	if err := q.Count(&out.Total).Error; err != nil {
		return out, err
	}

	offset := (page - 1) * pageSize
	err := db.Where("user_id = ?", userId).
		Order("date DESC").
		Preload("Sets.Exercise.MuscleGroup").
		Preload("Sets.Exercise.SecondaryMuscles.MuscleGroup").
		Offset(offset).
		Limit(pageSize).
		Find(&out.Items).Error

	return out, err
}

func GetTrainsWithCollectedSetsPaged(userId uint, page, pageSize int) (TrainsPage, error) {
	out := TrainsPage{
		Items: []Train{},
		Total: 0,
	}
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}

	q := db.Model(&Train{}).Where("user_id = ?", userId)
	if err := q.Count(&out.Total).Error; err != nil {
		return out, err
	}

	offset := (page - 1) * pageSize
	type PTrain struct {
		ID         uint    `gorm:"primaryKey"`
		UserWeight float64 `gorm:"column:user_weight"`
	}
	var trains []PTrain

	err := db.Table("trains").
		Select("id", "user_weight").
		Where("user_id = ?", userId).
		Order("id DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&trains).Error

	fmt.Println(trains)

	if err != nil {
		return out, err
	}

	var trainIds []uint

	for _, train := range trains {
		trainIds = append(trainIds, train.ID)
	}

	type Set struct {
		ID         uint    `gorm:"primaryKey;column:id"`
		Reps       int     `gorm:"column:reps"`
		Weight     float64 `gorm:"column:weight"` // замените на int/decimal, если в БД другой тип
		ExerciseID uint    `gorm:"column:exercise_id"`
		TrainID    uint    `gorm:"column:train_id"`
	}

	var oldSetsTable []Set

	err = db.Table("sets").
		Model(&Set{}).
		Select("id, reps, weight, exercise_id, train_id").
		Where("train_id IN ?", trainIds).
		Find(&oldSetsTable).Error

	fmt.Println(oldSetsTable)

	if err != nil {
		return out, err
	}

	type SetCombat struct {
		ID         uint
		Reps       int
		Sets       int
		Weight     float64
		ExerciseID uint
		TrainID    uint
	}

	groupedSets := make(map[string]*SetCombat)

	for _, s := range oldSetsTable {
		key := fmt.Sprintf("%d_%f_%d", s.Reps, s.Weight, s.ExerciseID)

		if existing, ok := groupedSets[key]; ok {
			existing.Sets++
		} else {
			groupedSets[key] = &SetCombat{
				ID:         s.ID,
				Sets:       1,
				Reps:       s.Reps,
				Weight:     s.Weight,
				ExerciseID: s.ExerciseID,
				TrainID:    s.TrainID,
			}
		}
	}

	var result []SetCombat
	for _, val := range groupedSets {
		result = append(result, *val)
	}

	return out, nil
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
