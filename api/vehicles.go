package handler

import (
	"net/http"
	"strconv"

	"github.com/yksen/copilot-webapp/utils"
)

func Vehicles(w http.ResponseWriter, r *http.Request) {
	db, err := utils.GetDatabase()
	utils.CheckPanic(w, err)
	defer db.Close()

	templates, err := utils.Templates()
	utils.CheckPanic(w, err)

	switch r.Method {
	case http.MethodPost:
		vehicle := utils.Vehicle{
			Name:            r.FormValue("vehicleName"),
			ApiKey:          r.FormValue("apiKey"),
			ApplicationName: r.FormValue("applicationName"),
			WebhookName:     r.FormValue("webhookName"),
			DeviceName:      r.FormValue("deviceName"),
		}

		_, err = db.Exec("INSERT INTO vehicles (name, api_key, application_name, webhook_name, device_name) VALUES ($1, $2, $3, $4, $5)", vehicle.Name, vehicle.ApiKey, vehicle.ApplicationName, vehicle.WebhookName, vehicle.DeviceName)
		utils.Check(w, err)

		w.WriteHeader(http.StatusCreated)

	case http.MethodGet:

	case http.MethodDelete:
		vehicleId := r.FormValue("vehicleId")

		_, err = db.Exec("DELETE FROM vehicles WHERE vehicle_id = $1", vehicleId)
		utils.Check(w, err)
	}

	var vehicleId int
	vehicleIdString := r.URL.Query().Get("vehicleId")
	if vehicleIdString != "" {
		vehicleId, err = strconv.Atoi(vehicleIdString)
		utils.Check(w, err)
	}

	data := struct {
		Vehicles []utils.Vehicle
	}{
		Vehicles: []utils.Vehicle{},
	}

	rows, err := db.Query("SELECT vehicle_id, name FROM vehicles")
	utils.Check(w, err)

	for rows.Next() {
		var vehicle utils.Vehicle
		err = rows.Scan(&vehicle.VehicleId, &vehicle.Name)
		utils.Check(w, err)
		vehicle.SelectedVehicleId = vehicleId
		data.Vehicles = append(data.Vehicles, vehicle)
	}

	if vehicleIdString != "" {
		err = templates.ExecuteTemplate(w, "vehicleSelect", data)
		utils.Check(w, err)
	} else {
		err = templates.ExecuteTemplate(w, "vehicles", data)
		utils.Check(w, err)
	}
}
