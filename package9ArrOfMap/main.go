package main

import (
	"fmt"
	"strings"
)

type Alert struct {
	ID, Alert_ID, ALERT_STATUS, ALERT_CREATED_DT, ALERT_UPDATED_DT, CSP, ACCOUNT_ID string
	RESOURCE_NAME, RESOURCE_ID, POLICY_ID, POLICY_NAME, CSP_EVENT_TYPE              string
}

func main() {

	test := Alert{}
	//m := make(map[Alert]string)
	test.ID = "123"
	test.Alert_ID = "ZS-CLOUD-A-00001"
	test.ALERT_STATUS = "Open"
	test.ALERT_CREATED_DT = "131313331"
	test.ALERT_UPDATED_DT = "313131331"
	test.CSP = "AWS"
	test.ACCOUNT_ID = "a1234"
	test.RESOURCE_NAME = "test-bucket"
	test.RESOURCE_ID = "123"
	test.POLICY_ID = "p1"
	test.POLICY_NAME = "s3 bucket is violating policy as public access is enabled"
	test.CSP_EVENT_TYPE = "create"

	fmt.Println(test)

	fmt.Println("*************************************")
	idAry := []string{"1", "2", "3"}
	ids := strings.Join(idAry, "','")

	fmt.Println(ids)

	fmt.Println("=====================================")
	inq := "6,7"
	sql := fmt.Sprintf("SELECT * FROM Person WHERE ALERT_ID IN ( %s )", inq)
	sql2 := fmt.Sprintf("Hello %s", "vivek")
	sql3 := sql + sql2
	fmt.Println(sql, sql2, " : ", sql3)

	fmt.Println("::::::::::::::::::::::::::::::::::::::::::::::::::::")

	temp := fmt.Sprintf("%s IN ( %s )", "ALERT_ID", inq)
	fmt.Println(temp)

	fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")

	example := []string{"123456", "234567", "345678", "456789"}
	commaSep := "'" + strings.Join(example, "', '") + "'"

	fmt.Println(commaSep)
}
