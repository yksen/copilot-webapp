package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func Data(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	rows, err := db.Query("SELECT * FROM records")
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
		fmt.Fprintf(w, "Record: %v, %v, %v, %v<br>", id, createdAt, recordType, value)
	}

	for i := 0; i < 100; i++ {
		fmt.Fprint(w, "test<br>")
	}
}
