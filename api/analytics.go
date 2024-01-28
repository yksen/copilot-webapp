package handler

import (
	"net/http"

	"github.com/yksen/copilot-webapp/utils"
)

func Analytics(w http.ResponseWriter, r *http.Request) {
	db, err := utils.GetDatabase()
	utils.CheckPanic(w, err)
	defer db.Close()

	templates, err := utils.Templates()
	utils.CheckPanic(w, err)

	switch r.Method {
	case http.MethodGet:
		err = templates.ExecuteTemplate(w, "analytics", nil)
		utils.Check(w, err)
	}
}
