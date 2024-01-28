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

//go:embed templates/*.html
var embedFiles embed.FS

func Templates() (*template.Template, error) {
	templates, err := template.ParseFS(embedFiles, "templates/*.html")
	if err != nil {
		return nil, err
	}
	return templates, nil
}
