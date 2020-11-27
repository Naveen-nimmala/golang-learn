package main

import "fmt"

type region struct {
	Name      string
	countries int
}

func (r region) prints() {
	fmt.Println(r.Name, r.countries)
}
