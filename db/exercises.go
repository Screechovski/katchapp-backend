package db

import (
	"fmt"
	"log"
)

type Exercise struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ImageName string `json:"imageName"`
}

func initExercises() {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS exercises (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            image VARCHAR(255)
        );
    `)
	if err != nil {
		log.Fatal(fmt.Errorf("failed creating exercises table: %w", err))
	}

	exercises, err := GetAllExercises()

	if err != nil {
		log.Println("failed to get exercises:", err)
		return
	}

	if len(exercises) == 0 {
		var exercises = []Exercise{
			{
				ID:        1,
				ImageName: "1.jpg",
				Name:      "Шраги со штангой",
			},
			{
				ID:        2,
				ImageName: "2.jpg",
				Name:      "Шраги со штангой за спиной",
			},
			{
				ID:        3,
				ImageName: "3.jpg",
				Name:      "Тяга штанги к подбородку",
			},
			{
				ID:        4,
				ImageName: "4.jpg",
				Name:      "Жим штанги стоя",
			},
			{
				ID:        5,
				ImageName: "5.jpg",
				Name:      "Жим штанги сидя",
			},
			{
				ID:        6,
				ImageName: "6.jpg",
				Name:      "Жим гантелей сидя",
			},
			{
				ID:        7,
				ImageName: "7.jpg",
				Name:      "Жим Арнольда",
			},
			{
				ID:        8,
				ImageName: "8.jpg",
				Name:      "Разведение гантелей стоя",
			},
			{
				ID:        9,
				ImageName: "9.jpg",
				Name:      "Подъем гантелей перед собой",
			},
			{
				ID:        10,
				ImageName: "10.jpg",
				Name:      "Разведение гантелей в наклоне",
			},
			{
				ID:        11,
				ImageName: "11.jpg",
				Name:      "Обратные разведения рук в тренажере",
			},
			{
				ID:        12,
				ImageName: "12.jpg",
				Name:      "Подъем гантелей над головой через стороны",
			},
			{
				ID:        13,
				ImageName: "13.jpg",
				Name:      "Подъем штанги на бицепс стоя",
			},
			{
				ID:        14,
				ImageName: "14.jpg",
				Name:      "Молоток",
			},
			{
				ID:        15,
				ImageName: "15.jpg",
				Name:      "Подъем гантелей на бицепс в скамье Скотта",
			},
			{
				ID:        16,
				ImageName: "16.jpg",
				Name:      "Подъем на бицепс в блочном тренажере стоя",
			},
			{
				ID:        17,
				ImageName: "17.jpg",
				Name:      "Сгибание рук на бицепс в кроссовере",
			},
			{
				ID:        18,
				ImageName: "18.jpg",
				Name:      "Концентрированный подъем на бицепс",
			},
			{
				ID:        19,
				ImageName: "19.jpg",
				Name:      "Подъем штанги на бицепс обратным хватом",
			},
			{
				ID:        20,
				ImageName: "20.jpg",
				Name:      "Подъем гантелей на бицепс стоя",
			},
			{
				ID:        21,
				ImageName: "21.jpg",
				Name:      "Подъем гантелей на бицепс сидя",
			},
			{
				ID:        22,
				ImageName: "22.jpg",
				Name:      "Жим штанги узким хватом лежа",
			},
			{
				ID:        23,
				ImageName: "23.jpg",
				Name:      "Отжимания от скамьи",
			},
			{
				ID:        24,
				ImageName: "24.jpg",
				Name:      "Французский жим лежа",
			},
			{
				ID:        25,
				ImageName: "25.jpg",
				Name:      "Французский жим EZ-штанги сидя",
			},
			{
				ID:        26,
				ImageName: "26.jpg",
				Name:      "Французский жим в тренажере сидя",
			},
			{
				ID:        27,
				ImageName: "27.jpg",
				Name:      "Жим книзу в блочном тренажере",
			},
			{
				ID:        28,
				ImageName: "28.jpg",
				Name:      "Жим книзу одной рукой обратным хватом",
			},
			{
				ID:        29,
				ImageName: "29.jpg",
				Name:      "Разгибание руки с гантелью из-за головы",
			},
			{
				ID:        30,
				ImageName: "30.jpg",
				Name:      "Разгибания руки с гантелью в наклоне",
			},
			{
				ID:        31,
				ImageName: "31.jpg",
				Name:      "Жим штанги лежа",
			},
			{
				ID:        32,
				ImageName: "32.jpg",
				Name:      "Жим штанги лежа вверх",
			},
			{
				ID:        33,
				ImageName: "33.jpg",
				Name:      "Жим штанги лежа вниз",
			},
			{
				ID:        34,
				ImageName: "34.jpg",
				Name:      "Жим гантелей лежа",
			},
			{
				ID:        35,
				ImageName: "35.jpg",
				Name:      "Жим гантелей лежа вверх",
			},
			{
				ID:        36,
				ImageName: "36.jpg",
				Name:      "Жим гантелей лежа вниз",
			},
			{
				ID:        37,
				ImageName: "37.jpg",
				Name:      "Жим от груди в тренажере сидя",
			},
			{
				ID:        38,
				ImageName: "38.jpg",
				Name:      "Разведение гантелей лежа",
			},
			{
				ID:        39,
				ImageName: "39.jpg",
				Name:      "Разведение гантелей лежа вверх",
			},
			{
				ID:        40,
				ImageName: "40.jpg",
				Name:      "Сведения рук в тренажере ",
			},
			{
				ID:        41,
				ImageName: "41.jpg",
				Name:      "Сведение в кроссовере через верхние блоки",
			},
			{
				ID:        42,
				ImageName: "42.jpg",
				Name:      "Сведение в кроссовере через нижние блоки",
			},
			{
				ID:        43,
				ImageName: "43.jpg",
				Name:      "Скручивания",
			},
			{
				ID:        44,
				ImageName: "44.jpg",
				Name:      "Скручивания на римсокм стуле",
			},
			{
				ID:        45,
				ImageName: "45.jpg",
				Name:      "Скручивания на скамье с наклоном вниз",
			},
			{
				ID:        46,
				ImageName: "46.jpg",
				Name:      "Скручивания на коленях в блочном тренажере",
			},
			{
				ID:        47,
				ImageName: "47.jpg",
				Name:      "Обратные скручивания",
			},
			{
				ID:        48,
				ImageName: "48.jpg",
				Name:      "Подъем коленей в висе",
			},
			{
				ID:        49,
				ImageName: "49.jpg",
				Name:      "Подъем ног в висе",
			},
			{
				ID:        50,
				ImageName: "50.jpg",
				Name:      "Косые скручивания",
			},
			{
				ID:        51,
				ImageName: "51.jpg",
				Name:      "Подтягивания на перекладине",
			},
			{
				ID:        52,
				ImageName: "52.jpg",
				Name:      "Тяга штанги в наклоне",
			},
			{
				ID:        53,
				ImageName: "53.jpg",
				Name:      "Тяга штанги в наклоне обратным хватом",
			},
			{
				ID:        54,
				ImageName: "54.jpg",
				Name:      "Тяга Т-штанги",
			},
			{
				ID:        55,
				ImageName: "55.jpg",
				Name:      "Тяга гантели одной рукой в наклоне",
			},
			{
				ID:        56,
				ImageName: "56.jpg",
				Name:      "Вертикальная тяга широким хватом",
			},
			{
				ID:        57,
				ImageName: "57.jpg",
				Name:      "Вертикальная тяга обратным хватом",
			},
			{
				ID:        58,
				ImageName: "58.jpg",
				Name:      "Горизонтальная тяга в блочном тренажере",
			},
			{
				ID:        59,
				ImageName: "59.jpg",
				Name:      "Пуловер в блочном тренажере стоя",
			},
			{
				ID:        60,
				ImageName: "60.jpg",
				Name:      "Становая тяга",
			},
			{
				ID:        61,
				ImageName: "61.jpg",
				Name:      "Наклоны со штангой на плечах",
			},
			{
				ID:        62,
				ImageName: "62.jpg",
				Name:      "Приседания со штангой",
			},
			{
				ID:        63,
				ImageName: "63.jpg",
				Name:      "Приседания в тренажере Смита",
			},
			{
				ID:        64,
				ImageName: "64.jpg",
				Name:      "Приседания со штангой на груди в тренажере Смита",
			},
			{
				ID:        65,
				ImageName: "65.jpg",
				Name:      "Гак-приседания",
			},
			{
				ID:        66,
				ImageName: "66.jpg",
				Name:      "Жим ногами",
			},
			{
				ID:        67,
				ImageName: "67.jpg",
				Name:      "Выпады со штангой",
			},
			{
				ID:        68,
				ImageName: "68.jpg",
				Name:      "Вышагивания на платформу",
			},
			{
				ID:        69,
				ImageName: "69.jpg",
				Name:      "Разгибания ног",
			},
			{
				ID:        70,
				ImageName: "70.jpg",
				Name:      "Рывок штанги на грудь",
			},
			{
				ID:        71,
				ImageName: "71.jpg",
				Name:      "Становая тяга на прямых ногах",
			},
			{
				ID:        72,
				ImageName: "72.jpg",
				Name:      "Румынский подъем",
			},
			{
				ID:        73,
				ImageName: "73.jpg",
				Name:      "Гиперэкстензия для мышц бедра",
			},
			{
				ID:        74,
				ImageName: "74.jpg",
				Name:      "Сгибание ног лежа",
			},
			{
				ID:        75,
				ImageName: "75.jpg",
				Name:      "Сгибания ног стоя",
			},
			{
				ID:        76,
				ImageName: "76.jpg",
				Name:      "Сгибания ног сидя",
			},
			{
				ID:        77,
				ImageName: "77.jpg",
				Name:      "Подъемы на носки стоя",
			},
			{
				ID:        78,
				ImageName: "78.jpg",
				Name:      "Подъемы на носки в тренажере для жимов ногами",
			},
			{
				ID:        79,
				ImageName: "79.jpg",
				Name:      "Подъемы на носки сидя",
			},
			{
				ID:        80,
				ImageName: "80.jpg",
				Name:      "Подъемы носков",
			},
			{
				ID:        81,
				ImageName: "81.jpg",
				Name:      "Сгибания рук в запястьях",
			},
			{
				ID:        82,
				ImageName: "82.jpg",
				Name:      "Отжимания на брусьях",
			},
			{
				ID:        83,
				ImageName: "83.jpg",
				Name:      "Отжимания от пола",
			},
			{
				ID:        84,
				ImageName: "84.jpg",
				Name:      "Приседания",
			},
			{
				ID:        85,
				ImageName: "85.jpg",
				Name:      "Выпады с гантелями",
			},
			{
				ID:        86,
				ImageName: "86.jpg",
				Name:      "Разведение рук в нижнем кроссовере",
			},
			{
				ID:        87,
				ImageName: "87.jpg",
				Name:      "Подъем ног в упоре на локтях",
			},
			{
				ID:        88,
				ImageName: "88.jpg",
				Name:      "Подъем коленей в упоре на локтях",
			},
			{
				ID:        89,
				ImageName: "89.jpg",
				Name:      "Поочередное сгибание рук с гантелями",
			},
			{
				ID:        90,
				ImageName: "90.jpg",
				Name:      "Тяга нижнего блока к поясу сидя",
			},
			{
				ID:        91,
				ImageName: "91.jpg",
				Name:      "Бег",
			},
			{
				ID:        92,
				ImageName: "92.jpg",
				Name:      "Велотренажер",
			},
			{
				ID:        93,
				ImageName: "93.jpg",
				Name:      "Пуловер",
			},
			{
				ID:        94,
				ImageName: "94.jpg",
				Name:      "Шраги с гантелями",
			},
			{
				ID:        95,
				ImageName: "95.jpg",
				Name:      "Подъем EZ-штанги на бицепс на скамье Скотта",
			},
		}
		query := "INSERT INTO exercises (id, name, image) VALUES "

		for i, exercise := range exercises {
			query += fmt.Sprintf("(%d, '%s', '%s')", exercise.ID, exercise.Name, exercise.ImageName)
			if i < len(exercises)-1 {
				query += ", "
			}
		}
		query += ";"
		_, err := db.Exec(query)
		if err != nil {
			log.Println("failed to insert exercise:", err)
		}
	}
}

func GetAllExercises() ([]Exercise, error) {
	list := make([]Exercise, 0)

	rows, err := db.Query("SELECT * FROM exercises ORDER BY id")
	if err != nil {
		log.Println("failed to query exercises:", err)
		return list, err
	}
	defer rows.Close()

	for rows.Next() {
		var e Exercise
		if err := rows.Scan(&e.ID, &e.Name, &e.ImageName); err != nil {
			log.Println("failed to scan exercise:", err)
			return list, err
		}
		list = append(list, e)
	}
	if err := rows.Err(); err != nil {
		log.Println("rows error:", err)
		return list, err
	}

	return list, err
}
