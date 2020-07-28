package main

import (
	"fmt"

	"github.com/diego-rodrigues/leftpad"
)

func main() {
	fmt.Println(leftpad.Format("hello, world", 30))
	fmt.Println(leftpad.Format("hello, world 🧐", 30))
	fmt.Println(leftpad.FormatRune("hello, world", 30, '🧐'))
	fmt.Println(leftpad.Format("hello, world. yaaaaas", 30))
}
