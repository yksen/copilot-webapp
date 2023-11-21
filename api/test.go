package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func Test(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	rows, err := db.Query("SELECT version()")
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	for rows.Next() {
		var version string
		if err := rows.Scan(&version); err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
		fmt.Fprintf(w, "Version: %v", version)
	}

	defer db.Close()
}
