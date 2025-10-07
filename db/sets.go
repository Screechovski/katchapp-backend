package db

import (
	"database/sql"
	"fmt"
	"log"
)

type SetsParams struct {
	ExerciseId int `json:"exerciseId"`
	Reps       int `json:"reps"`
	Weight     int `json:"weight"`
}

type Sets struct {
	ID         int `json:"id"`
	TrainId    int `json:"trainId"`
	ExerciseId int `json:"exerciseId"`
	Reps       int `json:"reps"`
	Weight     int `json:"weight"`
}

func initSets() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS sets (
			id SERIAL PRIMARY KEY,
			train_id INT NOT NULL REFERENCES trains(id) ON DELETE CASCADE,
			exercise_id INT NOT NULL REFERENCES exercises(id) ON DELETE RESTRICT,
			reps INT CHECK (reps BETWEEN 1 AND 1000),
			weight DECIMAL(6,2) CHECK (weight BETWEEN 0.1 AND 1000)
		);
	`)

	if err != nil {
		log.Fatal(fmt.Errorf("failed creating sets table: %w", err))
	}
}

func WriteSets(sets []SetsParams, trainId int) ([]int, error) {
	var err error
	var ids []int

	for i := range sets {
		id, err := WriteSet(sets[i], trainId)

		if err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	return ids, err
}

func WriteSet(set SetsParams, trainId int) (int, error) {
	return insert(func() *sql.Row {
		return db.QueryRow(
			`INSERT INTO sets (train_id, exercise_id, reps, weight)
			VALUES ($1, $2, $3, $4)
			RETURNING id`,
			trainId, set.ExerciseId, set.Reps, set.Weight,
		)
	})
}
