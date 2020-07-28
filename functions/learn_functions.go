package main

import "fmt"

func addNumbers(a, b int) int {
	return a + b
}

func divAndRemainder(a, b int) (int, int) {
	return a / b, a % b
}

func callByValue(a int, arr [2]int, s string) {
	a = a + a
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] + arr[i]
	}
	s = s + s
	fmt.Println(a, arr, s)
}

func main() {
	a := addNumbers(2, 3)
	fmt.Println(a)
	a = addNumbers(2, 7)
	fmt.Println(a)
	a = addNumbers(10, -3)
	fmt.Println(a)

	b, c := divAndRemainder(10, 3)
	fmt.Println(b, c)
	b, _ = divAndRemainder(94, 3)
	fmt.Println(b)

	arr := [2]int{1, 2}
	s := "hello"
	callByValue(a, arr, s)
	fmt.Println(a, arr, s)
}
