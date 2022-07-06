package main

import "fmt"

func main() {

	var map1 map[int]string = map[int]string{
		1: "Delhi",
		2: "Pune",
		3: "Mumbai",
		4: "Manali",
		5: "Gangtok",
	}

	for key, value := range map1 {
		fmt.Println(key, " : ", value)
	}

	_, ok := map1[6]

	fmt.Println("Is map having the key? ", ok)

	value, key := map1[2]

	fmt.Println("*********************")

	fmt.Println("if map is having the key and then value: ", key, " ", value)

	delete(map1, 5)

	fmt.Println("*********************")

	fmt.Println(map1)

}
