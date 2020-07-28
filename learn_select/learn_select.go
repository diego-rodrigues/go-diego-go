package main

import (
	"fmt"
	"time"
)

func multiples(i int) (chan int, chan struct{}) {
	out := make(chan int)
	done := make(chan struct{})
	curVal := 0
	go func() {
		for {
			select {
			case out <- curVal * i:
				curVal++
			case <-done:
				fmt.Println("goroutine shutting down...")
				return
			}
		}
	}()
	return out, done
}

func main2() {
	in := make(chan int, 1)
	out := make(chan int, 1)

	out <- 1

	// in <- 2
	// fmt.Println("wrote 2 to in")

	// v := <-out
	// fmt.Println("read", v, "from out")

	select {
	case in <- 2:
		fmt.Println("wrote 2 to in")
	case v := <-out:
		fmt.Println("read", v, "from out")
	}

	// ----------------------
	in2 := make(chan int)
	out2 := make(chan int)

	select {
	case in2 <- 3:
		fmt.Println("wrote 3 to in2")
	case v := <-out2:
		fmt.Println("read", v, "from out2")
	default:
		fmt.Println("nothing worked... =/")
	}
	// ----------------

	twosCh, doneCh := multiples(2)
	for v := range twosCh {
		if v > 20 {
			break
		}
		fmt.Println(v)
	}
	close(doneCh)

	time.Sleep(1 * time.Second)
}
