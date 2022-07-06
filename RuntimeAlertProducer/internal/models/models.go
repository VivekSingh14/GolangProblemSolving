package models

type RuntimeAlert struct {
	ID               string
	Alert_ID         string
	ALERT_STATUS     string
	ALERT_CREATED_DT string
	ALERT_UPDATED_DT string
	CSP              string
	ACCOUNT_ID       string
	RESOURCE_NAME    string
	RESOURCE_ID      string
	POLICY_ID        string
	POLICY_NAME      string
	CSP_EVENT_TYPE   string
}
