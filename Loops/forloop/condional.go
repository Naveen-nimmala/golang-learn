package main

import (
	"fmt"
)

func main() {
	var i int
	for i := 0; i < 1; i++ {
		fmt.Println("1st Loop", i+1)
	}
	for ;i < 1;i++{
		fmt.Println("2nd Loop", i+1)
	}
	fmt.Println(i)
	for i <= 1 {
		fmt.Println("3rd loop", i+1)
		i++
	}
}

