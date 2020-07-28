package main

import (
	"fmt"

	"github.com/diego-rodrigues/language/mapper"
)

func main() {
	fmt.Println(mapper.Greet("howdy, what's new?"))
	fmt.Println(mapper.Greet("Comment allez vous?"))
}
