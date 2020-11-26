package main

import "fmt"

type (
	Airways  []string
	Air      [3]string
	Airway   []int
	airindia string
	numbers  int
	result   bool
	ratio    float64
)

func main() {

	flights := Airways{"Indigo", "airasia", "airindia"}
	fmt.Println(flights)
	flight := Air{"hello", "world", "im"}
	fmt.Println(flight)

	airint := Airway{1, 2, 3}
	fmt.Println(airint)

	myname := airindia("naveen")
	fmt.Println(myname)
	mynumber := numbers(10)
	fmt.Println(mynumber)

	newfun()

}

func newfun() {

	var (
		aa Air
		bb Airways
		cc Airway
	)

	fmt.Println(aa, bb, cc)

}
