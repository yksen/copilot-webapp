package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/yksen/copilot-webapp/utils"
)

func Tools(w http.ResponseWriter, r *http.Request) {
	db, err := utils.GetDatabase()
	utils.CheckPanic(w, err)
	defer db.Close()

	templates, err := utils.Templates()
	utils.CheckPanic(w, err)

	switch r.Method {
	case http.MethodGet:
		err = templates.ExecuteTemplate(w, "tools", nil)
		utils.Check(w, err)

	case http.MethodPost:
		values := []string{}
		toolId := r.FormValue("toolId")
		switch toolId {
		case "1":
			fallthrough
		case "2":
			values = append(values, r.FormValue("duration"))
		case "3":
			force := utils.GetCheckboxValue(r.FormValue("force"))
			values = append(values, fmt.Sprintf("%d", force), r.FormValue("duration"), r.FormValue("duty"), r.FormValue("frequency"))
		case "4":
			state := utils.GetCheckboxValue(r.FormValue("state"))
			values = append(values, fmt.Sprintf("%d", state))
		default:
			panic("Unknown tool")
		}
		requestBody := fmt.Sprintf(`{"downlinks":[{
			"decoded_payload":{
				"arguments":"%s"
			},
			"f_port":%s
		  }]
		}`, strings.Join(values, ","), toolId)
		bodyReader := io.Reader(strings.NewReader(requestBody))

		vehicleId := r.FormValue("vehicleId")
		vehicle, err := utils.GetVehicleById(db, vehicleId)
		utils.CheckPanic(w, err)

		requestUrl := fmt.Sprintf("%s/api/v3/as/applications/%s/webhooks/%s/devices/%s/down/push",
			os.Getenv("TTN_URL"), vehicle.ApplicationName, vehicle.WebhookName, vehicle.DeviceName)

		request, err := http.NewRequest(http.MethodPost, requestUrl, bodyReader)
		utils.Check(w, err)

		request.Header.Set("Authorization", "Bearer "+vehicle.ApiKey)
		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("User-Agent", "my-integration/my-integration-version")

		_, err = http.DefaultClient.Do(request)
		utils.Check(w, err)
	}
}
