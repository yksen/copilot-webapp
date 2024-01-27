package handler

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/yksen/copilot-webapp/utils"
)

func Settings(w http.ResponseWriter, r *http.Request) {
	db, err := utils.GetDatabase()
	utils.CheckPanic(w, err)
	defer db.Close()

	switch r.Method {
	case http.MethodGet:
		templates, err := utils.Templates()
		utils.CheckPanic(w, err)

		err = templates.ExecuteTemplate(w, "settings", nil)
		utils.Check(w, err)
	}
}
