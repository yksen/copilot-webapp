package handler

import (
	"net/http"

	"github.com/yksen/copilot-webapp/utils"
)

func Tools(w http.ResponseWriter, r *http.Request) {
	db, err := utils.GetDatabase()
	utils.CheckPanic(w, err)
	defer db.Close()

	templates, err := utils.Templates()
	utils.CheckPanic(w, err)

	switch r.Method {
	case http.MethodPost:

	}

	err = templates.ExecuteTemplate(w, "tools", nil)
	utils.Check(w, err)
}
