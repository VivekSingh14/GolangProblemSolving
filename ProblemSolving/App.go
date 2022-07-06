package main

import "fmt"

var batchSize1 = 5

func main1() {
	count := 16
	tempCount := 0
	if count > 0 && count >= batchSize1 {
		limit := count / batchSize1
		remain := count % batchSize1
		//fmt.Println("**************** ", remain)
		for limit > 0 {
			for i := 0; i < batchSize1; i++ {
				fmt.Println(tempCount)
				tempCount++
			}
			fmt.Println("===================")
			limit--
		}
		for i := 0; i < remain; i++ {
			fmt.Println(tempCount)
			tempCount++
		}
	}
}
