package api_controller

import (
	"database/sql"
	"speechs/databases"
)

func CurrentTime(db *sql.DB) (string, error) {
	time, err := databases.QueryFile(db, "scripts/time.sql")
	if err != nil {
		return "", err
	}
	time.Next()
	var currentTime string
	err = time.Scan(&currentTime)
	if err != nil {
		return "", err
	}
	return currentTime, nil
}
