package api

import (
	"database/sql"
	"net/http"
	"speechs/api/templates_controller"
)

func BindControllers(db *sql.DB) {
	templatesController := templates_controller.CreateTemplateController(db)
	http.HandleFunc("/templates", templatesController.HandlerFunc())
}
