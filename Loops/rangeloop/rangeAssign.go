package main

import "fmt"

func main() {

	india := []string{"hello", "naveen", "how are you"}
	chars := []string{}
	// lenths := len(india)
	for _, v := range india {
		for _, vv := range v {
			chars = append(chars, string(vv))
		}
	}
	fmt.Println(chars)
	for _, v := range chars {
		switch v {
		case "a", "n":
			fmt.Println("naveen")
		}
	}
}
