package handler

import (
	"database/sql"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/yksen/copilot-webapp/utils"
)

func Init(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	utils.Check(w, err)
	defer db.Close()

	createTable := `
		CREATE TABLE IF NOT EXISTS records (
			id SERIAL PRIMARY KEY,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			type VARCHAR(255) NOT NULL,
			value VARCHAR(255) NOT NULL);`
	_, err = db.Exec(createTable)
	utils.Check(w, err)
}
