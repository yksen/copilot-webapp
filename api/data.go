package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/yksen/copilot-webapp/templates"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func Data(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	check(err)

	rows, err := db.Query("SELECT * FROM records")
	check(err)

	type Record struct {
		Id         int
		CreatedAt  string
		RecordType string
		Value      string
	}

	data := struct {
		Records []Record
	}{
		Records: []Record{},
	}

	for rows.Next() {
		var record Record
		err = rows.Scan(&record.Id, &record.CreatedAt, &record.RecordType, &record.Value)
		check(err)
		data.Records = append(data.Records, record)
	}

	templates.Table.Execute(w, data)
	check(err)
}
