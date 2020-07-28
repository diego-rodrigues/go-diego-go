package main

import (
	"fmt"
	"os"
)

func main() {
	word := os.Args[1]
	switch word {
	case "hi":
		fmt.Println("Very informal...")
		fallthrough
	case "hello":
		fmt.Println("Hi yourself!")
	case "goodbye":
		fmt.Println("See you next time.")
	case "greetings":
		fmt.Println("Salut!")
	default:
		fmt.Println("I don't know what you mean.")
	}
}
