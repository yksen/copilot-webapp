package handler

import (
	"html/template"
	"net/http"

	"github.com/yksen/copilot-webapp/utils"
)

var templates *template.Template

func Index(w http.ResponseWriter, r *http.Request) {
	templates, err := utils.Templates()
	utils.CheckPanic(w, err)

	err = templates.ExecuteTemplate(w, "index", nil)
	utils.Check(w, err)
}
