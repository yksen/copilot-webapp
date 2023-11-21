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
		fmt.Fprintf(w, "Version: %v<br>", version)
	}

	rows, err = db.Query("SELECT COUNT(*) FROM records")
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	for rows.Next() {
		var count int
		if err := rows.Scan(&count); err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
		fmt.Fprintf(w, "Count: %v<br>", count)
	}

	rows, err = db.Query("SELECT * FROM records")
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	for rows.Next() {
		var id int
		var createdAt string
		var recordType string
		var value string
		if err := rows.Scan(&id, &createdAt, &recordType, &value); err != nil {
			fmt.Fprintf(w, "Error: %v", err)
			return
		}
		fmt.Fprintf(w, "ID: %v, Created at: %v, Type: %v, Value: %v<br>", id, createdAt, recordType, value)
	}

	defer db.Close()
}
