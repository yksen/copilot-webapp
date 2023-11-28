package handler

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func Data(w http.ResponseWriter, r *http.Request) {
	check := func(err error) {
		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
		}
	}

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

	tmpl, err := template.New("table").Parse(`
		<table>
			<tr>
				<th>Id</th>
				<th>CreatedAt</th>
				<th>RecordType</th>
				<th>Value</th>
			</tr>
			{{range $r := .Records}}
			<tr>
				<td>{{$r.Id}}</td>
				<td>{{$r.CreatedAt}}</td>
				<td>{{$r.RecordType}}</td>
				<td>{{$r.Value}}</td>
			</tr>
			{{end}}
		</table>
	`)
	check(err)

	err = tmpl.Execute(w, data)
	check(err)
}
