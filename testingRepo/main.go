package main

import (
	"fmt"
	"github.com/ZscalerCWP/cwp-alert-common/pkg/snowflake"
	"strings"
)

var test1 = "alertsteamtest"
var test2 = "IAC"
var test3 = "SCAN"

func main() {

	scanTableName := snowflake.GetFullyQualifiedName(strings.ToUpper(test1), test2, test1)

	query := "select distinct parse_json(" + snowflake.RepoDetails + "):full_name "
	tmp := fmt.Sprintf("%s FROM %s", query, scanTableName)
	fmt.Println(tmp)

}
