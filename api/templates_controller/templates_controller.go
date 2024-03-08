package templates_controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"speechs/api/api_controller"
	"speechs/services/templates"
)

func GetTemplates(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Proto)
		readTemplates(db, w, r)
	}
}

func PostTemplate(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("PostTemplate"))
	}
}

func PutTemplate(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("PutTemplate"))
	}
}

func DeleteTemplate(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("DeleteTemplate"))
	}
}

func CreateTemplateController(db *sql.DB) *api_controller.ApiController {
	return &api_controller.ApiController{
		Get:    GetTemplates(db),
		Post:   PostTemplate(db),
		Put:    PutTemplate(db),
		Delete: DeleteTemplate(db),
	}
}

func readTemplates(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	limit, offset := api_controller.ExtractLimitOffset(r)
	templates, err := templates.ReadTemplates(db, limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(templates)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response: %s", err), http.StatusInternalServerError)
	}
}

func readTemplateByIds(db *sql.DB, w http.ResponseWriter, r *http.Request) {

}

func refreshTemplates(db *sql.DB, w http.ResponseWriter, r *http.Request) {

}
