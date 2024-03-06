package main

import (
	"fmt"
	"net/http"
	"speechs/databases"
	"speechs/settings"
)

func main() {
	setting := settings.CreateSettingsFromFlags()
	dbInfo := databases.New(setting.DBUser, setting.DBPassword, setting.DBHost, setting.DBName)
	db, err := databases.CreateDatabase(dbInfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	http.ListenAndServe(setting.AppPort(), nil)
}
