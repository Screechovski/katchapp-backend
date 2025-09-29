package db

import (
	"database/sql"
	"fmt"
	"log"
)

type TrainWithoutId struct {
	Date       string  `json:"date"`
	UserId     int     `json:"user"`
	UserWeight float32 `json:"userWeight"`
}

type Train struct {
	ID int `json:"id"`
	TrainWithoutId
}

type TrainForUser struct {
	ID         int     `json:"id"`
	Date       string  `json:"date"`
	UserWeight float32 `json:"userWeight"`
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
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS trains (
            id SERIAL PRIMARY KEY,
            date DATE NOT NULL,
            user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
            user_weight DECIMAL(5,2)
        );
    `)

	if err != nil {
		log.Fatal(fmt.Errorf("failed creating trains table: %w", err))
	}
}

func WriteTrain(payload TrainWithoutId) (int, error) {
	return insert(func() *sql.Row {
		return db.QueryRow(
			`INSERT INTO trains (date, user_id, user_weight)
			VALUES ($1,$2,$3)
			RETURNING id`,
			payload.Date,
			payload.UserId,
			payload.UserWeight,
		)
	})
}

/*
SELECT
	trains.id as train_id,
	trains.date,
	trains.user_weight,
	sets.exercise_id,
	sets.reps,
	sets.weight
FROM trains
INNER JOIN sets
ON trains.id = sets.train_id
WHERE trains.user_id = 1
ORDER BY trains.id DESC;
*/

func GetTrainsByUser(userId int) ([]TrainForUserWithSets, error) {
	clientTrains := make([]TrainForUserWithSets, 0)
	trainsList := make([]TrainForUser, 0)

	rows, err := db.Query(
		`SELECT 
			trains.id,
			trains.date,
			trains.user_weight
		FROM trains
		WHERE trains.user_id = $1
		ORDER BY trains.id DESC;`,
		userId,
	)

	if err != nil {
		log.Println("failed to query trains:", err)
		return clientTrains, err
	}

	defer rows.Close()

	for rows.Next() {
		var train TrainForUser

		if err := rows.Scan(&train.ID, &train.Date, &train.UserWeight); err != nil {
			log.Println("failed to scan exercise:", err)
			return clientTrains, err
		}

		trainsList = append(trainsList, train)
	}

	if err := rows.Err(); err != nil {
		log.Println("rows error:", err)
		return clientTrains, err
	}

	for _, train := range trainsList {
		sets := make([]TrainForUserSet, 0)

		rows, err := db.Query(
			`SELECT 
				sets.id,
				sets.exercise_id,
				sets.reps,
				sets.weight,
				exercises.image,
				exercises.name
			FROM sets
			INNER JOIN exercises
			ON exercises.id = sets.exercise_id
			WHERE sets.train_id = $1
			ORDER BY sets.id DESC;`,
			train.ID,
		)

		if err != nil {
			log.Println("failed to query trains:", err)
			return clientTrains, err
		}

		for rows.Next() {
			var set TrainForUserSet

			if err := rows.Scan(
				&set.ID,
				&set.ExerciseId,
				&set.Reps,
				&set.Weight,
				&set.ExerciseImageName,
				&set.ExerciseName,
			); err != nil {
				log.Println("failed to scan exercise:", err)
				return clientTrains, err
			}

			if err != nil {
				log.Println("failed to query trains:", err)
				return clientTrains, err
			}

			sets = append(sets, set)
		}

		clientTrains = append(clientTrains, TrainForUserWithSets{
			ID:         train.ID,
			Date:       train.Date,
			UserWeight: train.UserWeight,
			Sets:       sets,
		})
	}

	return clientTrains, err
}
