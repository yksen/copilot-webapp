package handler

import (
	"net/http"

	"github.com/yksen/copilot-webapp/utils"
)

func Analytics(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		templates, err := utils.Templates()
		utils.CheckPanic(w, err)

		err = templates.ExecuteTemplate(w, "analytics", nil)
		utils.Check(w, err)
	}
}
