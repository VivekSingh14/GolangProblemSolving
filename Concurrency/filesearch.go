package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

func main1() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	fmt.Println(homeDir)

	fileInhomedir, err := ioutil.ReadDir(homeDir)

	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	wg.Add(len(fileInhomedir))
	for _, file := range fileInhomedir {
		go func(f os.FileInfo) {
			fmt.Println(file)
			defer wg.Done()
		}(file)
	}
	wg.Wait()
}
