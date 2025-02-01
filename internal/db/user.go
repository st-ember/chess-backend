package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
)

type User struct {
	ID       string
	Username string
	Password string
}

func RetrievePassword(username string) (string, error) {
	query := "SELECT password FROM users WHERE username = $1"

	var dbPwd string
	err := DB.QueryRow(query, username).Scan(&dbPwd)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("User not found: ", err)
			return "", err
		}
		return "", err
	}

	return dbPwd, nil
}

func CreateUser(username, encPwd string, elo int16) (err error) {
	userId := uuid.NewString()
	userQuery := "INSERT INTO users (id, username, password), VALUES (?, ?, ?)"
	_, err = DB.Exec(userQuery, userId, username, encPwd)
	if err != nil {
		return err
	}

	currEloId := uuid.NewString()
	currEloQuery := "INSERT INTO user_current_elos (id, user_id, elo), VALUES (?, ?, ?)"
	_, err = DB.Exec(currEloQuery, currEloId, userId, elo)
	if err != nil {
		return err
	}

	historyEloId := uuid.NewString()
	historyEloQuery := "INSERT INTO user_elos_history" +
		"(id, user_id, elo, start)" +
		"VALUES(?, ?, ?, ?)"
	_, err = DB.Exec(historyEloQuery, historyEloId, userId, elo, time.Now())
	if err != nil {
		return err
	}

	return nil
}
