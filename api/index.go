package handler

import (
	"net/http"

	"github.com/yksen/copilot-webapp/templates"
)

func Index(w http.ResponseWriter, r *http.Request) {
	templates.Index.Execute(w, nil)
}
