package main

import "fmt"

func main() {

	var1 := 10
	var2 := "naveen"

	chnage(&var1, &var2)
	fmt.Println(var1, var2)

}

func chnage(x *int, y *string) {
	fmt.Println(*x, *y)
	*x = 20
	*y = "hello"
}
