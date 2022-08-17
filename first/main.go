package main

import (
	"fmt"
	"os"
)

//adding comment.
func main() {
	f, err := os.Create("/Users/viveksingh/Documents/WorkDir/repositories/repo2/demo1/first/test.txt")

	if err != nil {
		panic("cannot create a file")
	} else {
		fmt.Printf("err = %+v \n", err)
	}
	defer f.Close()
	fmt.Fprintf(f, "Hello")
}
