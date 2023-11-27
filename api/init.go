package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func Init(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	createTable := `
		CREATE TABLE IF NOT EXISTS records (
			id SERIAL PRIMARY KEY,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			type VARCHAR(255) NOT NULL,
			value VARCHAR(255) NOT NULL);`

	if _, err := db.Exec(createTable); err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	addSampleData := "INSERT INTO records (type, value) VALUES (1, UNNEST(ARRAY['a', 'b', 'c']))"

	if _, err := db.Exec(addSampleData); err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
}
