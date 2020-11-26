package main

import (
	"fmt"
	"os"
)

func main() {
	items := os.Args
	var item string
	if len(items) == 1 {
		fmt.Println("Enter one string")
		return
	} else if len(items) == 2 {
		item = os.Args[1]
	}

	switch item {
	case "one":
		fmt.Println("ONE")
	case "two":
		fmt.Println("TWO")
	}
}
