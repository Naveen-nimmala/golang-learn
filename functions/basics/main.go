package main

import (
	"fmt"
	"strconv"
)

func main() {

	local := 10
	show(local)
	incrWrng(local)
	show(local)

	local = incrRight(local)
	show(local)

	local = increseBy2(local, 2)
	show(local)
	local, _ = increseByStr(local, "hello")
	show(local)
}

func show(n int) {
	// can't access main's local
	// fmt.Printf("show : n = %d\n", local)
	fmt.Printf("show  → n = %d\n", n)
}

func incrWrng(n int) {
	n++
}

func incrRight(n int) int {
	n++
	return n
}

func increseBy2(n, num int) int {

	return n * num
}

func increseByStr(n int, str string) (int, error) {
	val, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("This is not valid", err)
	}
	return val * n, err
}
