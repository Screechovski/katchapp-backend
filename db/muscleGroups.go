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

	type seedGroup struct {
		Name    string
		PathIds []int
	}
	seeds := []seedGroup{
		{Name: "Грудные мышцы", PathIds: []int{78, 77}},
		{Name: "Бицепс", PathIds: []int{49, 48}},
		{Name: "Трицепс", PathIds: []int{109, 108}},
		{Name: "Передняя дельта", PathIds: []int{113, 79, 112, 80}},
		{Name: "Средняя дельта", PathIds: []int{113, 79, 112, 80}},
		{Name: "Задняя дельта", PathIds: []int{113, 79, 112, 80}},
		{Name: "Предплечья", PathIds: []int{31, 26, 30, 27, 127, 134, 139, 133, 126, 140}},
		{Name: "Прямая мышца живота (пресс)", PathIds: []int{64, 63, 56, 55, 44, 45, 29, 28}},
		{Name: "Косые мышцы живота", PathIds: []int{66, 58, 53, 47, 40, 65, 60, 54, 46, 41}},
		{Name: "Квадрицепсы", PathIds: []int{14, 15, 12, 11, 16, 13}},
		{Name: "Бицепсы бедра", PathIds: []int{106, 111, 110, 107}},
		{Name: "Ягодицы", PathIds: []int{102, 101}},
		{Name: "Икры", PathIds: []int{116, 119, 117, 118}},
		{Name: "Трапеции", PathIds: []int{114, 115, 92, 93, 90, 91, 100, 99}},
		{Name: "Спина", PathIds: nil},
		{Name: "Широчайшая мышца", PathIds: []int{98, 97}},
		{Name: "Большая круглая мышца", PathIds: []int{142, 141}},
	}

	if count == 0 {
		for _, s := range seeds {
			group := MuscleGroup{Name: s.Name}
			if result := db.Create(&group); result.Error != nil {
				return result.Error
			}
			for _, pathId := range s.PathIds {
				if result := db.Create(&MuscleGroupPath{MuscleGroupID: group.ID, PathId: pathId}); result.Error != nil {
					return result.Error
				}
			}
		}
		return nil
	}

	var countPaths int64
	db.Model(&MuscleGroupPath{}).Count(&countPaths)
	if countPaths > 0 {
		return nil
	}

	for _, s := range seeds {
		if len(s.PathIds) == 0 {
			continue
		}
		var group MuscleGroup
		if result := db.Where("name = ?", s.Name).First(&group); result.Error != nil {
			continue
		}
		for _, pathId := range s.PathIds {
			if result := db.Create(&MuscleGroupPath{MuscleGroupID: group.ID, PathId: pathId}); result.Error != nil {
				return result.Error
			}
		}
	}

	return nil
}

func GetAllMuscleGroups() ([]MuscleGroup, error) {
	var groups []MuscleGroup

	result := db.Find(&groups)

	return groups, result.Error
}
