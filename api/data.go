package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/yksen/copilot-webapp/templates"
)

type Record struct {
	Id        int
	CreatedAt string
	Type      string
	Value     string
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func Data(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	check(err)
	defer db.Close()

	switch r.Method {
	case http.MethodGet:
		rows, err := db.Query("SELECT * FROM records")
		check(err)

		data := struct {
			Records []Record
		}{
			Records: []Record{},
		}

		for rows.Next() {
			var record Record
			err = rows.Scan(&record.Id, &record.CreatedAt, &record.Type, &record.Value)
			check(err)
			data.Records = append(data.Records, record)
		}

		templates.Table.Execute(w, data)
		check(err)

	case http.MethodPost:
		record := Record{}
		record.Type = r.FormValue("type")
		record.Value = r.FormValue("value")

		_, err := db.Exec("INSERT INTO records (type, value) VALUES ($1, $2)", record.Type, record.Value)
		check(err)

		result, err := db.Query("SELECT COUNT(*) FROM records")
		check(err)

		var count int
		for result.Next() {
			err = result.Scan(&count)
			check(err)
		}

		fmt.Fprintf(w, "Record added successfully. Total records: %d", count)
	}
}
