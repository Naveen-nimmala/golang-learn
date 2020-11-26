package main

import "fmt"

func main() {

	var a int
	var b string
	var c bool
	var d float64

	fmt.Println(a, b, c, d)

	arys := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arys)

	slcs := []int{1, 2, 3, 4}
	fmt.Println(slcs)

	const val int = 10
	fmt.Println(val)

}
