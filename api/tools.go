package handler

import (
	"net/http"

	"github.com/yksen/copilot-webapp/utils"
)

func Tools(w http.ResponseWriter, r *http.Request) {
	templates, err := utils.Templates()
	utils.CheckPanic(w, err)

	err = templates.ExecuteTemplate(w, "tools", nil)
	utils.Check(w, err)
}
