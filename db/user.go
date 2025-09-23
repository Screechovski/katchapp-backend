package db

import (
	"database/sql"
)

type User struct {
	ID    int
	Token string
}

func FindUserWithToken(conn *sql.DB, token string) (User, error) {
	var user User
	err := conn.QueryRow("SELECT id, token FROM users WHERE token = $1", token).Scan(&user.ID, &user.Token)
	return user, err
}
