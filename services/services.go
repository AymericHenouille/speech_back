package services

import (
	"database/sql"
	"fmt"
	"speechs/databases"
)

func CurrentDate(db *sql.DB) string {
	rows, err := databases.QueryFile(db, "scripts/current_date.sql")
	if err != nil {
		fmt.Printf("Error querying current date: %s", err)
		return ""
	}
	rows.Next()
	var date string
	err = rows.Scan(&date)
	if err != nil {
		fmt.Printf("Error scanning current date: %s", err)
		return ""
	}
	return date
}
