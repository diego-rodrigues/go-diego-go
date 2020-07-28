package main

import "fmt"

func main() {
	a := 10
	if b := a / 2; a > 5 {
		fmt.Println("a is bigger than 5: ", a, b)
	} else {
		fmt.Println("a is not bigger than 5: ", a, b)
	}

	for i := 0; i < 10; i++ {
		if i > 7 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("--------")
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println("--------")
	s := "Hello, world!"
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}
}
