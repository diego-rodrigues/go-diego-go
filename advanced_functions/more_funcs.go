package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func addTwo(a int) int {
	return a + 2
}

func printOperation(a int, f func(int) int) {
	fmt.Println(f(a))
}

func makeAdder(b int) func(int) int {
	return func(a int) int {
		return b + a
	}
}

func makeDoubler(f func(int) int) func(int) int {
	return func(a int) int {
		return f(a) * 2
	}
}

func main() {
	myAnnon := func(a int) int {
		return a + 2
	}
	myAddOne := addOne
	fmt.Println(addOne(1))
	fmt.Println(myAddOne(1))
	fmt.Println(myAnnon(2))

	printOperation(10, addOne)
	printOperation(10, addTwo)

	// closure
	b := 2
	myNewAdd := func(a int) {
		b = a + b
	}
	myNewAdd(1)
	fmt.Println(b)

	newestAddOne := makeAdder(1)
	newestAddTwo := makeAdder(2)

	fmt.Println(newestAddOne(10))
	fmt.Println(newestAddTwo(10))

	addAndDouble := makeDoubler(newestAddOne)
	fmt.Println(addAndDouble(20))
}
