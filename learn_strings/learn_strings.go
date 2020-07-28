package main

import "fmt"

func main() {
	s := "Hello, World"
	b := s[4]
	fmt.Println(s, b)

	var vals [3]int
	vals[0] = 1
	vals[1] = 3
	vals[2] = 5
	fmt.Println(vals, vals[0], vals[1], vals[2])
}
