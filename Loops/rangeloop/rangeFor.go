package main

import (
	"fmt"
)

func main() {
	words := [...]string{"hello", "world"}
	for i,v := range words {
		fmt.Println(i)
		fmt.Println(v)
	}
	books := [5]string{"hello", "world"}
	for i := range books {
		fmt.Println(i)
	}
	games := []string{"hello", "world", "naaveen"}
	for i := range games {
		fmt.Println(i)
	}
	twoarr := [2][3]int{ {1,2,4},{2,3,4} }
	
	for i , _ := range twoarr {
		fmt.Println(i)
	}
}