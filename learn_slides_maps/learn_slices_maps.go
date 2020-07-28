package main

import "fmt"

func main() {
	// slices
	s := make([]string, 0)
	fmt.Println("len of s: ", len(s))
	s = append(s, "hello")
	fmt.Println("len of s: ", len(s))
	fmt.Println("content of s[0]: ", s[0])
	s[0] = "goodbye"
	fmt.Println("content of s[0]: ", s[0])

	s2 := make([]string, 2)
	fmt.Println("content of s2[0]: ", s2[0])
	s2 = append(s2, "hello")
	fmt.Println("content of s2[0]: ", s2[0])
	fmt.Println("content of s2[2]: ", s2[2])
	fmt.Println("len of s2: ", len(s2))

	s3 := []string{"a", "b", "c"}

	for k, v := range s3 {
		fmt.Println(k, v)
	}

	s4 := s3[0:2]
	fmt.Println("s4: ", s4)
	s3[0] = "d"
	fmt.Println("s4: ", s4)

	var s5 []string
	s5 = s3
	s5[1] = "camel"
	fmt.Println("s3: ", s3)

	modSlice(s3)
	fmt.Println("s3[0]: ", s3[0])
	fmt.Println("-------------------")

	// maps
	m := make(map[string]int)
	m["hello"] = 300
	h := m["hello"]
	fmt.Println("hello in m: ", h)
	fmt.Println("a in m: ", m["a"])

	if v, ok := m["hello"]; ok {
		fmt.Println("hello in m: ", v)
	}

	m["hello"] = 20
	fmt.Println("hello in m: ", m["hello"])

	m2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 50,
	}
	for k, v := range m2 {
		fmt.Println(k, v)
	}

	fmt.Println("b in m2: ", m2["b"])
	delete(m2, "b")
	fmt.Println("b in m2: ", m2["b"])

	m3 := m
	m3["goodbye"] = -1
	fmt.Println("goodbye in m3: ", m3["goodbye"])
	fmt.Println("goodbye in m: ", m["goodbye"])

	fmt.Println("cheese in m: ", m["cheese"])
	modMap(m)
	fmt.Println("cheese in m: ", m["cheese"])
}

func modSlice(s []string) {
	s[0] = "hello"
}

func modMap(m map[string]int) {
	m["cheese"] = 20
}
