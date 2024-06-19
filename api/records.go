package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/yksen/copilot-webapp/utils"
)

func Records(w http.ResponseWriter, r *http.Request) {
	db, err := utils.GetDatabase()
	utils.CheckPanic(w, err)
	defer db.Close()

	templates, err := utils.Templates()
	utils.CheckPanic(w, err)

	switch r.Method {
	case http.MethodGet:
		recordsPerPage := 10
		pageString := r.URL.Query().Get("page")
		page := 0
		if pageString != "" {
			page, err = strconv.Atoi(pageString)
			utils.Check(w, err)
		}

		vehicleIdString := r.URL.Query().Get("vehicleId")
		vehicleId := 0
		if vehicleIdString != "" {
			vehicleId, err = strconv.Atoi(vehicleIdString)
			utils.Check(w, err)
		}

		rows, err := db.Query("SELECT * FROM records WHERE vehicle_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3",
			vehicleId, recordsPerPage, page*recordsPerPage)
		utils.Check(w, err)

		data := struct {
			Records []utils.Record
		}{
			Records: []utils.Record{},
		}

		for rows.Next() {
			var record utils.Record
			err = rows.Scan(&record.RecordId, &record.CreatedAt, &record.Type, &record.Value, &record.VehicleId)
			utils.Check(w, err)
			data.Records = append(data.Records, record)
		}

		err = templates.ExecuteTemplate(w, "records", data)
		utils.Check(w, err)

	case http.MethodPost:
		record := utils.Record{}
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

		var vehicleId int
		vehicleIdRow := db.QueryRow("SELECT vehicle_id FROM vehicles ORDER BY" +
			" created_at DESC LIMIT 1")
		err = vehicleIdRow.Scan(&vehicleId)
		utils.Check(w, err)

		_, err = db.Exec("INSERT INTO records (type, value, vehicle_id) VALUES"+
			" ($1, $2, $3)", record.Type, record.Value, vehicleId)
		utils.Check(w, err)

		w.WriteHeader(http.StatusCreated)

		result, err := db.Query("SELECT COUNT(*) FROM records")
		utils.Check(w, err)

		var count int
		for result.Next() {
			err = result.Scan(&count)
			utils.Check(w, err)
		}

		fmt.Fprintf(w, "Record added successfully. Total records: %d",
			count)
	}
}
