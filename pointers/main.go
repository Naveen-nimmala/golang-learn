package main

import "fmt"

func main() {

	var (
		Avalue int
		P      *int
		V      int
	)

	Avalue = 10

	P = &Avalue

	V = *P

	fmt.Println(Avalue, P, V, &P, *P)

	Avalue = 1

	fmt.Println(Avalue, P, V, &P, *P)

	changeValue(P)
	fmt.Println(Avalue, P, V, &P, *P)

}

func changeValue(AV *int) {
	*AV = 1000

	*AV++

	fmt.Println("this is def func", AV, &AV, *AV)

}
