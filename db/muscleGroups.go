package db

import (
	"gorm.io/gorm"
)

type MuscleGroup struct {
	gorm.Model
	Name      string
	Exercises []Exercise `gorm:"foreignKey:MuscleGroupID"`
}

func initMuscleGroup() error {
	if err := db.AutoMigrate(&MuscleGroup{}); err != nil {
		return err
	}

	var count int64
	db.Model(&MuscleGroup{}).Count(&count)

	type seedGroup struct {
		Name    string
		PathIds []int
	}

	if count == 0 {
		groups := []MuscleGroup{
			{Name: "Грудные мышцы"},
			{Name: "Бицепс"},
			{Name: "Трицепс"},
			{Name: "Передняя дельта"},
			{Name: "Средняя дельта"},
			{Name: "Задняя дельта"},
			{Name: "Предплечья"},
			{Name: "Прямая мышца живота (пресс)"},
			{Name: "Косые мышцы живота"},
			{Name: "Квадрицепсы"},
			{Name: "Бицепсы бедра"},
			{Name: "Ягодицы"},
			{Name: "Икры"},
			{Name: "Трапеции"},
			{Name: "Спина"},
			{Name: "Широчайшая мышца"},
			{Name: "Большая круглая мышца"},
		}
		for _, group := range groups {
			if result := db.Create(&group); result.Error != nil {
				return result.Error
			}
		}
		return nil
	}

	return nil
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
