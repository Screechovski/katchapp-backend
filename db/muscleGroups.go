package db

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type MuscleGroup struct {
	gorm.Model
	Name      string
	Exercises []Exercise `gorm:"foreignKey:MuscleGroupID"`
}

func initMuscleGroup() {
	err := db.AutoMigrate(&MuscleGroup{})

	if err != nil {
		log.Fatal(fmt.Errorf("failed creating muscle_groups table: %w", err))
	}

	var count int64
	db.Model(&MuscleGroup{}).Count(&count)

	if count == 0 {
		var groups = []MuscleGroup{
			{Name: "Грудные мышцы"},
			{Name: "Спина"},
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
		}

		result := db.Create(groups)

		if result.Error != nil {
			log.Println("failed to insert muscle_groups:", result.Error)
		}
	}
}
