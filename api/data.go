package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
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
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = io.ReadAll(r.Body)
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	check(err)
	defer db.Close()

	switch r.Method {
	case http.MethodGet:
		rows, err := db.Query("SELECT * FROM records ORDER BY created_at DESC")
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
		contentType := r.Header.Get("Content-Type")

		if contentType == "application/json" {
			payload := struct {
				UplinkMessage struct {
					DecodedPayload struct {
						Type  string
						Value string
					}
				}
			}{}
			json.NewDecoder(r.Body).Decode(&payload)

			record.Type = payload.UplinkMessage.DecodedPayload.Type
			record.Value = payload.UplinkMessage.DecodedPayload.Value
		} else {
			record.Type = r.FormValue("type")
			record.Value = r.FormValue("value")
		}

		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		_, err := db.Exec("INSERT INTO records (type, value) VALUES ($1, $2)", record.Type, fmt.Sprintf("%s", r.Body))
		check(err)

		result, err := db.Query("SELECT COUNT(*) FROM records")
		check(err)

		var count int
		for result.Next() {
			err = result.Scan(&count)
			check(err)
		}

		fmt.Fprintf(w, "Record added successfully. Total records: %d\n", count)
	}
}
