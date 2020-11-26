package main

import "fmt"

type Subs struct {
	Ma1, Ma2, Ma3, Ma4 string
}
type Student struct {
	Name     string
	marks    int
	subjects []Subs
	rollno   int
}

func main() {

	var sahi Student
	sahi.marks = 10
	fmt.Println(sahi)
	naveen := Student{
		Name:   "naveen",
		marks:  256,
		rollno: 1,
		subjects: []Subs{
			{
				Ma1: "eng",
				Ma2: "tel",
				Ma3: "maths",
				Ma4: "social",
			},
		},
	}

	marks(&naveen)
	fmt.Println(naveen.marks)
}

func marks(mr *Student) {
	mr.marks = 236
}
