package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/yksen/copilot-webapp/templates"
	"github.com/yksen/copilot-webapp/utils"
)

type Record struct {
	Id        int
	CreatedAt string
	Type      string
	Value     string
}

func Analytics(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	utils.CheckPanic(w, err)
	defer db.Close()

	switch r.Method {
	case http.MethodGet:
		recordsPerPage := 10
		pageString := r.URL.Query().Get("page")
		page := 0
		if pageString != "" {
			page, err = strconv.Atoi(pageString)
			utils.Check(w, err)
		}

		rows, err := db.Query("SELECT * FROM records ORDER BY created_at DESC LIMIT $1 OFFSET $2", recordsPerPage, page*recordsPerPage)
		utils.Check(w, err)

		data := struct {
			Records []Record
		}{
			Records: []Record{},
		}

		for rows.Next() {
			var record Record
			err = rows.Scan(&record.Id, &record.CreatedAt, &record.Type, &record.Value)
			utils.Check(w, err)
			data.Records = append(data.Records, record)
		}

		templates.Table.Execute(w, data)
		utils.Check(w, err)

	case http.MethodPost:
		record := Record{}
		requestBody := utils.GetRequestBody(r)
		contentType := r.Header.Get("Content-Type")

		switch contentType {
		case "application/json":
			payload := struct {
				UplinkMessage struct {
					DecodedPayload struct {
						Type  string `json:"type"`
						Value string `json:"value"`
					} `json:"decoded_payload"`
				} `json:"uplink_message"`
			}{}

			buffer := bytes.NewBuffer(requestBody)
			json.NewDecoder(buffer).Decode(&payload)

			record.Type = payload.UplinkMessage.DecodedPayload.Type
			record.Value = payload.UplinkMessage.DecodedPayload.Value
		case "application/x-www-form-urlencoded":
			record.Type = r.FormValue("type")
			record.Value = r.FormValue("value")
		}

		_, err := db.Exec("INSERT INTO records (type, value) VALUES ($1, $2)", record.Type, record.Value)
		utils.Check(w, err)

		result, err := db.Query("SELECT COUNT(*) FROM records")
		utils.Check(w, err)

		var count int
		for result.Next() {
			err = result.Scan(&count)
			utils.Check(w, err)
		}

		fmt.Fprintf(w, "Record added successfully. Total records: %d", count)
	}
}
