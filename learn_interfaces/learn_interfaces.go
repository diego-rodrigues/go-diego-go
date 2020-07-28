package main

import "fmt"

type tester interface {
	test(int) bool
}

func runTests(i int, tests []tester) bool {
	result := true
	for _, test := range tests {
		result = result && test.test(i)
	}
	return result
}

type rangeTest struct {
	min int
	max int
}

func (rt rangeTest) test(i int) bool {
	return rt.min <= i && i <= rt.max
}

type divTest int

func (dt divTest) test(i int) bool {
	return i%int(dt) == 0
}

func doStuff(i interface{}) {
	switch i := i.(type) {
	case int:
		fmt.Println("Double i is ", i*2)
	case string:
		fmt.Println("i is ", len(i), " characters long.")
	default:
		fmt.Println("don't know what to do with this")
	}
}

type testerFunc func(int) bool

func (tf testerFunc) test(i int) bool {
	return tf(i)
}

func main() {
	result := runTests(10, []tester{
		rangeTest{min: 5, max: 30},
		divTest(5),
	})
	fmt.Println(result)

	var i interface{}
	i = "Hello"
	j := i.(string)
	k, ok := i.(int)
	fmt.Println(j, k, ok)
	// m := i.(int)
	// fmt.Println(m)

	fmt.Println("-------------")
	doStuff(10)
	doStuff("hello")
	doStuff(true)

	fmt.Println("-------------")
	result = runTests(10, []tester{
		testerFunc(func(i int) bool {
			return i%2 == 0
		}),
		testerFunc(func(i int) bool {
			return i < 20
		}),
	})
	fmt.Println(result)
}
