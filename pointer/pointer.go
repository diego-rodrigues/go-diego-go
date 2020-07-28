package main

import "fmt"

func setTo10(a *int) {
	*a = 10
}

func main() {
	a := 10
	b := &a
	c := a
	fmt.Println(a, b, *b, c)

	a = 20
	fmt.Println(a, b, *b, c)

	*b = 30
	fmt.Println(a, b, *b, c)

	c = 40
	fmt.Println(a, b, *b, c)

	var b1 *int
	fmt.Println(b1)
	// fmt.Println(*b1)

	c1 := new(int)
	fmt.Println(c1)
	fmt.Println(*c1)

	fmt.Println(a)
	setTo10(&a)
	fmt.Println(a)
}
