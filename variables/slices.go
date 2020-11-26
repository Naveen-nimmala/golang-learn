package main

import "fmt"

func main() {

	// sl := []int{1, 2, 3, 4, 5, 6}
	sl := make([]int, 5, 10)

	for i := 0; i < 5; i++ {
		sl[i] = i
	}
	fmt.Println(sl)
	fmt.Println(len(sl))

	sl = append(sl, 99)
	fmt.Println(sl)

	ml := make([]int, len(sl))
	fmt.Println(ml)
	copy(ml, sl)
	fmt.Println(sl)

	kl := ml[:3]
	fmt.Println(kl)
}
