package app

import (
	"RuntimeAlertProducer/internal/models"
	"RuntimeAlertProducer/internal/service"
)

func Start() {
	runtimeAlerts := newRuntimeAlert()
	service.ProcessRuntimeAlerts(runtimeAlerts)

}

func newRuntimeAlert() []models.RuntimeAlert {
	var alertArray []models.RuntimeAlert
	ra := models.RuntimeAlert{
		ID:               "123",
		Alert_ID:         "234",
		ALERT_STATUS:     "Open",
		ALERT_CREATED_DT: "1801234",
		ALERT_UPDATED_DT: "1801234",
		CSP:              "AWS",
		ACCOUNT_ID:       "123456",
		RESOURCE_NAME:    "S3",
		RESOURCE_ID:      "fj123",
		POLICY_ID:        "45678",
		POLICY_NAME:      "Test",
		CSP_EVENT_TYPE:   "Any",
	}
	alertArray = append(alertArray, ra)
	ra2 := models.RuntimeAlert{
		ID:               "123",
		Alert_ID:         "234",
		ALERT_STATUS:     "Open",
		ALERT_CREATED_DT: "1801234",
		ALERT_UPDATED_DT: "1801234",
		CSP:              "AWS",
		ACCOUNT_ID:       "123456",
		RESOURCE_NAME:    "S3",
		RESOURCE_ID:      "fj123",
		POLICY_ID:        "45678",
		POLICY_NAME:      "Test",
		CSP_EVENT_TYPE:   "Any",
	}
	alertArray = append(alertArray, ra2)

	return alertArray
}
