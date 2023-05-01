package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

var batchSize int = 5

func main5() {
	var alertList map[string][]string
	batchSize := 5
	Data1 := `{
  "alert_id": [
    ""
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
