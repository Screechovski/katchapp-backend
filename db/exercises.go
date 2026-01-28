package db

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Exercise struct {
	gorm.Model
	Name             string `json:"name"`
	ImageName        string `json:"imageName"`
	MuscleGroupID    uint
	MuscleGroup      MuscleGroup
	SecondaryMuscles []ExerciseSecondaryMuscle `gorm:"foreignKey:ExerciseId"`
	Sets             []Sets                    `gorm:"foreignKey:ExerciseId"`
}

func getExercisesInitial() []Exercise {
	return []Exercise{
		{
			ImageName:        "1.jpg",
			Name:             "Шраги со штангой",
			MuscleGroupID:    15,
			SecondaryMuscles: []ExerciseSecondaryMuscle{},
		},
		{
			ImageName:        "2.jpg",
			Name:             "Шраги со штангой за спиной",
			MuscleGroupID:    15,
			SecondaryMuscles: []ExerciseSecondaryMuscle{},
		},
		{
			ImageName:     "3.jpg",
			Name:          "Тяга штанги к подбородку",
			MuscleGroupID: 5,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   6,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   15,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "4.jpg",
			Name:          "Жим штанги стоя",
			MuscleGroupID: 5,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   6,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   4,
					EngagementLevel: 2,
				},
			},
		},
		{
			ImageName:     "5.jpg",
			Name:          "Жим штанги сидя",
			MuscleGroupID: 5,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   6,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   4,
					EngagementLevel: 2,
				},
			},
		},
		{
			ImageName:     "6.jpg",
			Name:          "Жим гантелей сидя",
			MuscleGroupID: 5,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   6,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   4,
					EngagementLevel: 2,
				},
			},
		},
		{
			ImageName:     "7.jpg",
			Name:          "Жим Арнольда",
			MuscleGroupID: 5,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   6,
					EngagementLevel: 3,
				},
				{
					MuscleGroupId:   3,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "8.jpg",
			Name:          "Разведение гантелей стоя",
			MuscleGroupID: 6,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   5,
					EngagementLevel: 2,
				},
			},
		},
		{
			ImageName:     "9.jpg",
			Name:          "Подъем гантелей перед собой",
			MuscleGroupID: 5,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   6,
					EngagementLevel: 2,
				},
			},
		},
		{
			ImageName:     "10.jpg",
			Name:          "Разведение гантелей в наклоне",
			MuscleGroupID: 7,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   2,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:        "11.jpg",
			Name:             "Обратные разведения рук в тренажере",
			MuscleGroupID:    7,
			SecondaryMuscles: []ExerciseSecondaryMuscle{},
		},
		{
			ImageName:     "12.jpg",
			Name:          "Подъем гантелей над головой через стороны",
			MuscleGroupID: 5,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   6,
					EngagementLevel: 3,
				},
			},
		},
		{
			ImageName:     "13.jpg",
			Name:          "Подъем штанги на бицепс стоя",
			MuscleGroupID: 3,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   8,
					EngagementLevel: 2,
				},
			},
		},
		{
			ImageName:     "14.jpg",
			Name:          "Молоток",
			MuscleGroupID: 3,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   8,
					EngagementLevel: 3,
				},
			},
		},
		{
			ImageName:        "15.jpg",
			Name:             "Подъем гантелей на бицепс в скамье Скотта",
			MuscleGroupID:    3,
			SecondaryMuscles: []ExerciseSecondaryMuscle{},
		},
		{
			ImageName:     "16.jpg",
			Name:          "Подъем на бицепс в блочном тренажере стоя",
			MuscleGroupID: 3,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   8,
					EngagementLevel: 2,
				},
			},
		},
		{
			ImageName:        "17.jpg",
			Name:             "Сгибание рук на бицепс в кроссовере",
			MuscleGroupID:    3,
			SecondaryMuscles: []ExerciseSecondaryMuscle{},
		},
		{
			ImageName:        "18.jpg",
			Name:             "Концентрированный подъем на бицепс",
			MuscleGroupID:    3,
			SecondaryMuscles: []ExerciseSecondaryMuscle{},
		},
		{
			ImageName:     "19.jpg",
			Name:          "Подъем штанги на бицепс обратным хватом",
			MuscleGroupID: 3,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   8,
					EngagementLevel: 2,
				},
			},
		},
		{
			ImageName:     "20.jpg",
			Name:          "Подъем гантелей на бицепс стоя",
			MuscleGroupID: 3,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   8,
					EngagementLevel: 2,
				},
			},
		},

		{
			ImageName:     "21.jpg",
			Name:          "Подъем гантелей на бицепс сидя",
			MuscleGroupID: 3,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   8,
					EngagementLevel: 2,
				},
			},
		},
		{
			ImageName:     "22.jpg",
			Name:          "Жим штанги узким хватом лежа",
			MuscleGroupID: 4,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   1,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   5,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "23.jpg",
			Name:          "Отжимания от скамьи",
			MuscleGroupID: 4,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   1,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   5,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "24.jpg",
			Name:          "Французский жим лежа",
			MuscleGroupID: 4,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   5,
					EngagementLevel: 1,
				},
				{
					MuscleGroupId:   6,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "25.jpg",
			Name:          "Французский жим EZ-штанги сидя",
			MuscleGroupID: 4,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   5,
					EngagementLevel: 1,
				},
				{
					MuscleGroupId:   6,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "26.jpg",
			Name:          "Французский жим в тренажере сидя",
			MuscleGroupID: 4,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   5,
					EngagementLevel: 1,
				},
				{
					MuscleGroupId:   6,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "27.jpg",
			Name:          "Жим книзу в блочном тренажере",
			MuscleGroupID: 4,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   5,
					EngagementLevel: 1,
				},
				{
					MuscleGroupId:   6,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "28.jpg",
			Name:          "Жим книзу одной рукой обратным хватом",
			MuscleGroupID: 4,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   5,
					EngagementLevel: 1,
				},
				{
					MuscleGroupId:   6,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "29.jpg",
			Name:          "Разгибание руки с гантелью из-за головы",
			MuscleGroupID: 4,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   5,
					EngagementLevel: 1,
				},
				{
					MuscleGroupId:   6,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "30.jpg",
			Name:          "Разгибания руки с гантелью в наклоне",
			MuscleGroupID: 4,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   5,
					EngagementLevel: 1,
				},
				{
					MuscleGroupId:   6,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "31.jpg",
			Name:          "Жим штанги лежа",
			MuscleGroupID: 1,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   4,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   5,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "32.jpg",
			Name:          "Жим штанги лежа вверх",
			MuscleGroupID: 1,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   5,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   6,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "33.jpg",
			Name:          "Жим штанги лежа вниз",
			MuscleGroupID: 1,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   4,
					EngagementLevel: 3,
				},
				{
					MuscleGroupId:   5,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "34.jpg",
			Name:          "Жим гантелей лежа",
			MuscleGroupID: 1,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   4,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   6,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "35.jpg",
			Name:          "Жим гантелей лежа вверх",
			MuscleGroupID: 1,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   5,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   6,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "36.jpg",
			Name:          "Жим гантелей лежа вниз",
			MuscleGroupID: 1,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   4,
					EngagementLevel: 3,
				},
				{
					MuscleGroupId:   5,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "37.jpg",
			Name:          "Жим от груди в тренажере сидя",
			MuscleGroupID: 1,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   4,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   5,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "38.jpg",
			Name:          "Разведение гантелей лежа",
			MuscleGroupID: 1,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   6,
					EngagementLevel: 2,
				},
			},
		},
		{
			ImageName:     "39.jpg",
			Name:          "Разведение гантелей лежа вверх",
			MuscleGroupID: 1,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   5,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   6,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "40.jpg",
			Name:          "Сведения рук в тренажере",
			MuscleGroupID: 1,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   6,
					EngagementLevel: 2,
				},
			},
		},

		{
			ImageName:     "41.jpg",
			Name:          "Сведение в кроссовере через верхние блоки",
			MuscleGroupID: 1,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   6,
					EngagementLevel: 2,
				},
			},
		},
		{
			ImageName:     "42.jpg",
			Name:          "Сведение в кроссовере через нижние блоки",
			MuscleGroupID: 1,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   6,
					EngagementLevel: 2,
				},
			},
		},
		{
			ImageName:     "43.jpg",
			Name:          "Скручивания",
			MuscleGroupID: 9,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   10,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "44.jpg",
			Name:          "Скручивания на римском стуле",
			MuscleGroupID: 9,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   10,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "45.jpg",
			Name:          "Скручивания на скамье с наклоном вниз",
			MuscleGroupID: 9,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   10,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "46.jpg",
			Name:          "Скручивания на коленях в блочном тренажере",
			MuscleGroupID: 9,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   10,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "47.jpg",
			Name:          "Обратные скручивания",
			MuscleGroupID: 9,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   10,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "48.jpg",
			Name:          "Подъем коленей в висе",
			MuscleGroupID: 9,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   10,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "49.jpg",
			Name:          "Подъем ног в висе",
			MuscleGroupID: 9,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   10,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "50.jpg",
			Name:          "Косые скручивания",
			MuscleGroupID: 10,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   9,
					EngagementLevel: 2,
				},
			},
		},
		{
			ImageName:     "51.jpg",
			Name:          "Подтягивания на перекладине",
			MuscleGroupID: 2,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   3,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   7,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "52.jpg",
			Name:          "Тяга штанги в наклоне",
			MuscleGroupID: 2,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   15,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   7,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "53.jpg",
			Name:          "Тяга штанги в наклоне обратным хватом",
			MuscleGroupID: 2,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   3,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   15,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "54.jpg",
			Name:          "Тяга Т-штанги",
			MuscleGroupID: 2,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   15,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   7,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "55.jpg",
			Name:          "Тяга гантели одной рукой в наклоне",
			MuscleGroupID: 2,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   15,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   7,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "56.jpg",
			Name:          "Вертикальная тяга широким хватом",
			MuscleGroupID: 2,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   3,
					EngagementLevel: 1,
				},
				{
					MuscleGroupId:   7,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "57.jpg",
			Name:          "Вертикальная тяга обратным хватом",
			MuscleGroupID: 2,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   3,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   7,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "58.jpg",
			Name:          "Горизонтальная тяга в блочном тренажере",
			MuscleGroupID: 2,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   15,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   7,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "59.jpg",
			Name:          "Пуловер в блочном тренажере стоя",
			MuscleGroupID: 2,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   1,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "60.jpg",
			Name:          "Становая тяга",
			MuscleGroupID: 2,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   13,
					EngagementLevel: 3,
				},
				{
					MuscleGroupId:   11,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   15,
					EngagementLevel: 2,
				},
			},
		},

		{
			ImageName:     "61.jpg",
			Name:          "Наклоны со штангой на плечах",
			MuscleGroupID: 2,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   13,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   15,
					EngagementLevel: 2,
				},
			},
		},
		{
			ImageName:     "62.jpg",
			Name:          "Приседания со штангой",
			MuscleGroupID: 11,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   13,
					EngagementLevel: 3,
				},
				{
					MuscleGroupId:   12,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "63.jpg",
			Name:          "Приседания в тренажере Смита",
			MuscleGroupID: 11,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   13,
					EngagementLevel: 3,
				},
				{
					MuscleGroupId:   12,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "64.jpg",
			Name:          "Приседания со штангой на груди в тренажере Смита",
			MuscleGroupID: 11,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   13,
					EngagementLevel: 3,
				},
				{
					MuscleGroupId:   5,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "65.jpg",
			Name:          "Гак-приседания",
			MuscleGroupID: 11,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   13,
					EngagementLevel: 2,
				},
			},
		},
		{
			ImageName:     "66.jpg",
			Name:          "Жим ногами",
			MuscleGroupID: 11,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   13,
					EngagementLevel: 2,
				},
			},
		},
		{
			ImageName:     "67.jpg",
			Name:          "Выпады со штангой",
			MuscleGroupID: 11,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   13,
					EngagementLevel: 3,
				},
				{
					MuscleGroupId:   12,
					EngagementLevel: 2,
				},
			},
		},
		{
			ImageName:     "68.jpg",
			Name:          "Вышагивания на платформу",
			MuscleGroupID: 11,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   13,
					EngagementLevel: 3,
				},
				{
					MuscleGroupId:   12,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:        "69.jpg",
			Name:             "Разгибания ног",
			MuscleGroupID:    11,
			SecondaryMuscles: []ExerciseSecondaryMuscle{},
		},
		{
			ImageName:     "70.jpg",
			Name:          "Рывок штанги на грудь",
			MuscleGroupID: 2,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   11,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   13,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   15,
					EngagementLevel: 2,
				},
			},
		},
		{
			ImageName:     "71.jpg",
			Name:          "Становая тяга на прямых ногах",
			MuscleGroupID: 12,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   2,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   13,
					EngagementLevel: 2,
				},
			},
		},
		{
			ImageName:     "72.jpg",
			Name:          "Румынский подъем",
			MuscleGroupID: 12,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   2,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   13,
					EngagementLevel: 3,
				},
			},
		},
		{
			ImageName:     "73.jpg",
			Name:          "Гиперэкстензия для мышц бедра",
			MuscleGroupID: 12,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   13,
					EngagementLevel: 3,
				},
				{
					MuscleGroupId:   2,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:        "74.jpg",
			Name:             "Сгибание ног лежа",
			MuscleGroupID:    12,
			SecondaryMuscles: []ExerciseSecondaryMuscle{},
		},
		{
			ImageName:        "75.jpg",
			Name:             "Сгибания ног стоя",
			MuscleGroupID:    12,
			SecondaryMuscles: []ExerciseSecondaryMuscle{},
		},
		{
			ImageName:        "76.jpg",
			Name:             "Сгибания ног сидя",
			MuscleGroupID:    12,
			SecondaryMuscles: []ExerciseSecondaryMuscle{},
		},
		{
			ImageName:        "77.jpg",
			Name:             "Подъемы на носки стоя",
			MuscleGroupID:    14,
			SecondaryMuscles: []ExerciseSecondaryMuscle{},
		},
		{
			ImageName:        "78.jpg",
			Name:             "Подъемы на носки в тренажере для жимов ногами",
			MuscleGroupID:    14,
			SecondaryMuscles: []ExerciseSecondaryMuscle{},
		},
		{
			ImageName:        "79.jpg",
			Name:             "Подъемы на носки сидя",
			MuscleGroupID:    14,
			SecondaryMuscles: []ExerciseSecondaryMuscle{},
		},
		{
			ImageName:        "80.jpg",
			Name:             "Подъемы носков",
			MuscleGroupID:    14,
			SecondaryMuscles: []ExerciseSecondaryMuscle{},
		},

		{
			ImageName:        "81.jpg",
			Name:             "Сгибания рук в запястьях",
			MuscleGroupID:    8,
			SecondaryMuscles: []ExerciseSecondaryMuscle{},
		},
		{
			ImageName:     "82.jpg",
			Name:          "Отжимания на брусьях",
			MuscleGroupID: 4,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   1,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   5,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "83.jpg",
			Name:          "Отжимания от пола",
			MuscleGroupID: 1,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   4,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   5,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "84.jpg",
			Name:          "Приседания",
			MuscleGroupID: 11,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   13,
					EngagementLevel: 3,
				},
			},
		},
		{
			ImageName:     "85.jpg",
			Name:          "Выпады с гантелями",
			MuscleGroupID: 11,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   13,
					EngagementLevel: 3,
				},
				{
					MuscleGroupId:   12,
					EngagementLevel: 2,
				},
			},
		},
		{
			ImageName:     "86.jpg",
			Name:          "Разведение рук в нижнем кроссовере",
			MuscleGroupID: 7,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   2,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "87.jpg",
			Name:          "Подъем ног в упоре на локтях",
			MuscleGroupID: 9,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   10,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "88.jpg",
			Name:          "Подъем коленей в упоре на локтях",
			MuscleGroupID: 9,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   10,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "89.jpg",
			Name:          "Поочередное сгибание рук с гантелями",
			MuscleGroupID: 3,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   8,
					EngagementLevel: 2,
				},
			},
		},
		{
			ImageName:     "90.jpg",
			Name:          "Тяга нижнего блока к поясу сидя",
			MuscleGroupID: 2,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   15,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   7,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "91.jpg",
			Name:          "Бег",
			MuscleGroupID: 11,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   13,
					EngagementLevel: 2,
				},
				{
					MuscleGroupId:   14,
					EngagementLevel: 2,
				},
			},
		},
		{
			ImageName:     "92.jpg",
			Name:          "Велотренажер",
			MuscleGroupID: 11,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   13,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:     "93.jpg",
			Name:          "Пуловер",
			MuscleGroupID: 2,
			SecondaryMuscles: []ExerciseSecondaryMuscle{
				{
					MuscleGroupId:   1,
					EngagementLevel: 1,
				},
			},
		},
		{
			ImageName:        "94.jpg",
			Name:             "Шраги с гантелями",
			MuscleGroupID:    15,
			SecondaryMuscles: []ExerciseSecondaryMuscle{},
		},
		{
			ImageName:        "95.jpg",
			Name:             "Подъем EZ-штанги на бицепс на скамье Скотта",
			MuscleGroupID:    3,
			SecondaryMuscles: []ExerciseSecondaryMuscle{},
		},
	}
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

func initExercises() {
	err := db.AutoMigrate(&Exercise{})

	if err != nil {
		log.Fatal(fmt.Errorf("failed creating exercises table: %w", err))
	}

	var count int64
	db.Model(&Exercise{}).Count(&count)

	if count == 0 {
		exercises := getExercisesInitial()

		result := db.Create(&exercises)

		if result.Error != nil {
			log.Printf("failed to insert exercises: %v", result.Error)
			return
		}

		log.Printf("Successfully inserted %d exercises with secondary muscles", len(exercises))
	}
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
