package db

import (
	"gorm.io/gorm"
)

// TODO поправить связь для MuscleGroup.Paths
type MuscleGroup struct {
	gorm.Model
	Name      string
	Paths     []MuscleGroupPath `gorm:"foreignKey:MuscleGroupID"`
	Exercises []Exercise        `gorm:"foreignKey:MuscleGroupID"`
}

type MuscleGroupPath struct {
	gorm.Model
	MuscleGroupID uint
	PathId        int
}

// TODO 1 добавить id (из свг для каждой мышцы)

func initMuscleGroup() error {
	if err := db.AutoMigrate(&MuscleGroup{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&MuscleGroupPath{}); err != nil {
		return err
	}

	var count int64
	db.Model(&MuscleGroup{}).Count(&count)

	type Item struct {
		Name   string
		SvgIds []int
	}
	var groups = []MuscleGroup{
		{Name: "Грудные мышцы", []int{78, 77}},
		{Name: "Бицепс", []int{49, 48}},
		{Name: "Трицепс", []int{109, 108}},

		{Name: "Передняя дельта", []int{113, 79, 112, 80}},
		{Name: "Средняя дельта", []int{113, 79, 112, 80}},
		{Name: "Задняя дельта", []int{113, 79, 112, 80}},

		{Name: "Предплечья", []int{31, 26, 30, 27, 127, 134, 139, 133, 126, 140}},

		{Name: "Прямая мышца живота (пресс)", []int{64, 63, 56, 55, 44, 45, 29, 28}},
		{Name: "Косые мышцы живота", []int{66, 58, 53, 47, 40, 65, 60, 54, 46, 41}},

		{Name: "Квадрицепсы", Paths: []int{14, 15, 12, 11, 16, 13}},
		{Name: "Бицепсы бедра", Paths: []int{106, 111, 110, 107}},
		{Name: "Ягодицы", Paths: []int{102, 101}},
		{Name: "Икры", Paths: []int{116, 119, 117, 118}},
		{Name: "Трапеции", Paths: []int{114, 115, 92, 93, 90, 91, 100, 99}},
		{Name: "Спина"},
		{Name: "Широчайшая мышца", Paths: []int{98, 97}},
		{Name: "Большая круглая мышца", Paths: []int{142, 141}},
	}

	if count == 0 {
		result := db.Create(groups)

		if result.Error != nil {
			return result.Error
		}
	}

	var countPaths int64
	db.Model(&MuscleGroupPath{}).Count(&countPaths)
	if countPaths == 0 {
		var group MuscleGroup
		if result := db.Model(&MuscleGroup{}).Where("name = ?", groups[0].Name).Find(&group); result.Error != nil {
			return result.Error
		}
		path := MuscleGroupPath{MuscleGroupID: group.ID, PathId: 77}
		if result := db.Create(&path); result.Error != nil {
			return result.Error
		}
		path2 := MuscleGroupPath{MuscleGroupID: group.ID, PathId: 78}
		if result := db.Create(&path2); result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func GetAllMuscleGroups() ([]MuscleGroup, error) {
	var groups []MuscleGroup

	result := db.Find(&groups)

	return groups, result.Error
}
