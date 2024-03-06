package databases

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func CreateDatabase(info *OpenDBInfo) (*sql.DB, error) {
	datasources := info.Datasource()
	return sql.Open("postgres", datasources)
}

func QueryFile(db *sql.DB, file string, args ...any) (*sql.Rows, error) {
	query, err := findQueryFile(file)
	if err != nil {
		return nil, err
	}
	return db.Query(query, args...)
}

var queryFiles = map[string]string{}

func findQueryFile(file string) (string, error) {
	query, ok := queryFiles[file]
	if !ok {
		bytes, err := os.ReadFile(file)
		if err != nil {
			return "", fmt.Errorf("query file %s not found", file)
		}
		query = string(bytes)
		queryFiles[file] = query
	}
	return query, nil
}
