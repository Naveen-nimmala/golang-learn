package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func main() {

	fmt.Println(runtime.GOOS) // to check OS name

	if runtime.GOOS == "linux" || runtime.GOOS == "windows" {
		fmt.Println("THis is no mac os")
	} else {
		execute()
		executepwd()
	}
}

func execute() {
	out, err := exec.Command("ls", "-lrt").Output()
	if err != nil {
		log.Fatal(err)
	}
	output := string(out[:])
	fmt.Println(output)
}

func executepwd() {
	out, err := exec.Command("pwd").Output()
	if err != nil {
		log.Fatal(err)
	}
	output := string(out)
	fmt.Println(output)
}
