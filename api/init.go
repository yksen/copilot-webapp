package handler

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/yksen/copilot-webapp/utils"
)

func Init(w http.ResponseWriter, r *http.Request) {
	db, err := utils.GetDatabase()
	utils.CheckPanic(w, err)
	defer db.Close()

	dropVehicles := `DROP TABLE IF EXISTS vehicles;`
	_, err = db.Exec(dropVehicles)
	utils.Check(w, err)

	dropRecords := `DROP TABLE IF EXISTS records;`
	_, err = db.Exec(dropRecords)
	utils.Check(w, err)

	createVehicles := `
		CREATE TABLE vehicles (
			vehicle_id INT GENERATED ALWAYS AS IDENTITY,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			name VARCHAR(255) NOT NULL,
			api_key VARCHAR(255) NOT NULL,
			application_name VARCHAR(255) NOT NULL,
			webhook_name VARCHAR(255) NOT NULL,
			device_name VARCHAR(255) NOT NULL,
			PRIMARY KEY (vehicle_id));`
	_, err = db.Exec(createVehicles)
	utils.Check(w, err)

	createRecords := `
		CREATE TABLE records (
			record_id INT GENERATED ALWAYS AS IDENTITY,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			type VARCHAR(255) NOT NULL,
			value VARCHAR(255) NOT NULL,
			vehicle_id INT NOT NULL,
			PRIMARY KEY (record_id),
			FOREIGN KEY (vehicle_id) REFERENCES vehicles(vehicle_id));`
	_, err = db.Exec(createRecords)
	utils.Check(w, err)

	populateVehicles := `
		INSERT INTO vehicles (name, api_key, application_name, webhook_name, device_name)
		VALUES ('test', 'test', 'test', 'test', 'test');`
	_, err = db.Exec(populateVehicles)
	utils.Check(w, err)

	populateRecords := `
		INSERT INTO records (type, value, vehicle_id)
		VALUES ('speed', '0', 1);`
	_, err = db.Exec(populateRecords)
	utils.Check(w, err)
}
