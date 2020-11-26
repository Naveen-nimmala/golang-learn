package main

import "fmt"

func main() {

	//students := []string{"adddres", "why", "gdroutd", "hello"}

	// for _, val := range students {

	// 	for i, _ := range val {
	// 		fmt.Print(i)
	// 	}
	// 	fmt.Println()

	// }

	clgs := [][]string{
		{"adddres", "why", "gdroutd", "hello"},
		{"add", "dfd", "ad"},
	}
	count := 0
	// for _, vl := range clgs {
	// 	for indz, vlu := range vl {
	// 		fmt.Println(indz, vlu)
	// 	}
	// }
	// fmt.Println(len(clgs))
	for i := 0; i < len(clgs); i++ {
		for j := 0; j < len(clgs[i]); j++ {
			for k := 0; k < len(clgs[i][j]); k++ {
				if clgs[i][j][k] == 'a' {
					count++
					fmt.Printf("%c\n", clgs[i][j][k])

				}
			}
		}

	}
	fmt.Println(count)
}
