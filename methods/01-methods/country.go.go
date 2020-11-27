package main

import "fmt"

type country struct {
	Name       string
	Population int
	courency   string
}

func (c *country) prints(name string) {
	fmt.Println(c.Name, c.Population, c.courency)
	if name == "naveen" {
		c.Name = "Bhrath"
	}

}
