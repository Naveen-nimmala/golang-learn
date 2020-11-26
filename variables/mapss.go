package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["room"] = 40
	m["floor"] = 10
	m["block"] = 2

	fmt.Println(m["block"])

	as := m["room"]
	fmt.Println(as)

	for i, v := range m {
		fmt.Println(i, v)
	}
}
