package main

import (
	"encoding/json"
	"fmt"
)

type Details struct {
	Id, Name, City string
}

func main() {
	m := make(map[string]string)
	m["Name"] = "Vivek"
	m["Id"] = "12Jai"
	m["City"] = "Jaipur"
	m["Department"] = "Engineering"
	fmt.Println(m)

	fmt.Println("***************")

	data, _ := json.Marshal(&m)

	fmt.Println(data)

	fmt.Println("===================")

	var map1 Details
	err := json.Unmarshal(data, &map1)
	if err == nil {
		fmt.Println(map1)

	}
}
