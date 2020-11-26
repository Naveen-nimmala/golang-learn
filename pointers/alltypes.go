package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("ARRAYS")
	Arrys()
	slicess()
	maps()
	structs()

}

func Arrys() {
	nums := [...]int{1, 2, 3}

	incrNums(nums)
	fmt.Println("This is Pass By Value Function", nums)

	incrPtr(&nums)
	fmt.Println("This is Pass By Pointer Function", nums)

}

func incrNums(n [3]int) {
	for i := range n {
		n[i]++
	}
}

func incrPtr(n *[3]int) {
	for i := range n {
		n[i]++
	}
}

func slicess() {
	lists := []string{"one", "two", "three", "four"}
	upValue(lists)
	fmt.Println(lists)
	upPtr(&lists)
	fmt.Println(lists)
}

func upValue(slc []string) {
	for i, v := range slc {
		slc[i] = strings.ToUpper(v)
	}
	slc = append(slc, "FIVE")
}
func upPtr(slc *[]string) {
	lv := *slc
	for i, v := range lv {
		lv[i] = strings.ToUpper(v)
	}
	*slc = append(lv, "FIVE")
}

func maps() {
	mp := map[string]int{"one": 1, "two": 2}
	mapfn(mp)
	fmt.Println(mp)
}

func mapfn(mp map[string]int) {
	fmt.Println(mp["one"])
	mp["three"] = 3
	mp["four"] = 4
}

type Student struct {
	Rollno int
	clg    string
}

type Person struct {
	Name   string
	Sex    string
	Height float32
	stu    Student
}

func structs() {
	naveen := Person{
		Name:   "Naveen",
		Sex:    "Male",
		Height: 5.10,
		stu: Student{
			Rollno: 1,
			clg:    "Audisankara",
		},
	}

	addRollno(naveen)
	fmt.Println(naveen)
	addRollnoPtr(&naveen)
	fmt.Println(naveen)
}

func addRollno(n Person) {
	n.stu.Rollno = 2
	n.Name = "Naveen Nimmala"
}

func addRollnoPtr(n *Person) {
	n.stu.Rollno = 2
	n.Name = "Naveen Nimmala"
}
