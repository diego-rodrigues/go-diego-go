package main

import (
	"fmt"
	"os"
	"strconv"
)

func divAndMod(a int, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, &myError{A: a, B: b, Message: "Cannot divide by zero"}
	}
	return a / b, a % b, nil
}

type myError struct {
	A       int
	B       int
	Message string
}

func (me *myError) Error() string {
	return fmt.Sprintf("%d and %d produced error %s", me.A, me.B, me.Message)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Expected two input parameters")
		os.Exit(1)
	}
	a, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	b, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	div, mod, err := divAndMod(int(a), int(b))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("div: ", div, " mod: ", mod, " a: ", a, " b: ", b)

}
