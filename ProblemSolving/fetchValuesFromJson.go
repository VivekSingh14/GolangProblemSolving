package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

var batchSize int = 5

func main() {
	var alertList map[string][]string
	batchSize := 5
	Data1 := `{
  "alert_id": [
    "ZS-CLOUD-AZURE-22,ZS-CLOUD-GCP-23,ZS-CLOUD-AWS-24,ZS-CLOUD-AWS-20,ZS-CLOUD-AWS-21,ZS-CLOUD-AWS-15,ZS-CLOUD-AWS-16,ZS-CLOUD-GCP-7,ZS-CLOUD-AZURE-14,ZS-CLOUD-AWS-17"
  ]
}`
	err := json.Unmarshal([]byte(Data1), &alertList)

	if err != nil {
		return
	}
	res := alertList["alert_id"]
	strArr := strings.Split(res[0], ",")
	alertsLength := len(strArr)
	//fmt.Println(len(strArr))
	limit := alertsLength / batchSize
	//remain := alertsLength % batchSize
	totalAlertsInitial := 0
	for limit > 0 {
		var batch []string
		for i := 1; i <= batchSize; i++ {
			batch = append(batch, strArr[totalAlertsInitial])
			totalAlertsInitial++
		}
		//commaSep := "'" + strings.Join(batch, "', '") + "'"

		//fmt.Println(commaSep)
		//fmt.Println(fmt.Sprintf("%s IN ( %s )", "ALERT_ID", commaSep))
		//queryBuilder.Where(fmt.Sprintf("%s IN ( %s )", "ALERT_ID", commaSep))

		//fmt.Println(fmt.Sprintf("%s IN ( %s ) AND %s IN ( %s ) AND %s IN ( %s )", "ALERT_ID", commaSep, "ALERT_ID", commaSep, "ALERT_ID", commaSep))
		fmt.Println(fmt.Sprintf("%s = %s AND %s = %s ", "STATUS", "ENABLED", "ISSOFTDELETE", "FALSE"))
		limit--
	}

}
