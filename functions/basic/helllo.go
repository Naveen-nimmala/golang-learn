package main

import "fmt"

func hello() {
	fmt.Println("this is 2nd function calling")
}

func intr(a, b int) (az int, sx int) {
	a++
	b++
	return a, b
}
