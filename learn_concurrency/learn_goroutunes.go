package main

import (
	"fmt"
	"sync"
)

func runMe(name string) {
	fmt.Println("Hello ", name, " from a goroutine")
}

func main() {
	// go runMe()
	// time.Sleep(1 * time.Second)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(name string) {
		runMe(name)
		wg.Done()
	}("diego")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(localI int) {
			fmt.Println(localI)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
