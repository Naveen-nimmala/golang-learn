package main

import "fmt"

type human struct {
	speak bool
	eyes  int
	legs  int
	name  string
	hight float64
}

func main() {

	naveen := human{
		speak: true,
		eyes:  2,
		legs:  2,
		name:  "Naveen",
		hight: 5.8,
	}
	printe(naveen)
	naveen.met()
}

func printe(data human) {
	fmt.Println(data.eyes)
}

func (data human) met() {
	fmt.Println(data.name)
	fmt.Println(data.hight)
}
