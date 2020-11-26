package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func main() {

	var cuttime = 1 * time.Hour

	fileInfo, err := ioutil.ReadDir("/Users/naveen/Desktop/DevOps/Golang/go-exmp/new-go")
	if err != nil {
		log.Fatal(err.Error())
	}
	now := time.Now()

	for _, info := range fileInfo {
		//fmt.Println(info.Name(), now.Sub(info.ModTime()))
		diff := now.Sub(info.ModTime())
		if diff > cuttime {
			fmt.Printf("Deleting %s which is %s old\n", info.Name(), diff)
		} else {
			fmt.Printf("Not Deleting %s which is %s old\n", info.Name(), diff)
		}
	}

}
