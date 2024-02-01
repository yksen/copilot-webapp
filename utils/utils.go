package utils

import (
	"bytes"
	"database/sql"
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type Record struct {
	RecordId  int
	CreatedAt string
	Type      string
	Value     string
	VehicleId int
}

type Vehicle struct {
	VehicleId         int
	CreatedAt         string
	Name              string
	ApiKey            string
	ApplicationName   string
	WebhookName       string
	DeviceName        string
	SelectedVehicleId int
}

func Check(w http.ResponseWriter, err error) {
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CheckPanic(w http.ResponseWriter, err error) {
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
}

func GetRequestBody(r *http.Request) []byte {
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = io.ReadAll(r.Body)
	}
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	return bodyBytes
}

func GetDatabase() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetVehicleById(db *sql.DB, vehicleId string) (*Vehicle, error) {
	var vehicle Vehicle
	err := db.QueryRow("SELECT vehicle_id, api_key, application_name, webhook_name, device_name FROM vehicles WHERE vehicle_id = $1",
		vehicleId).Scan(&vehicle.VehicleId, &vehicle.ApiKey, &vehicle.ApplicationName, &vehicle.WebhookName, &vehicle.DeviceName)
	if err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func GetCheckboxValue(value string) int {
	if value == "on" {
		return 1
	}
	return 0
}

//go:embed templates/*.html
var embedFiles embed.FS

func Templates() (*template.Template, error) {
	templates, err := template.ParseFS(embedFiles, "templates/*.html")
	if err != nil {
		return nil, err
	}
	return templates, nil
}
