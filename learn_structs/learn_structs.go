package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	A int
	b string
}

type Bar struct {
	C string
	F Foo
}

type Baz struct {
	D string
	Foo
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	f := Foo{}
	fmt.Println(f)

	g := Foo{10, "hello"}
	fmt.Println(g)

	h := Foo{
		b: "goodbye",
	}
	fmt.Println(h)

	h.A = 1000
	fmt.Println(h)

	var f2 Foo
	f2 = f
	f2.A = 100
	fmt.Println(f2)
	fmt.Println(f)

	var f3 *Foo = &f
	f3.A = 300
	fmt.Println(*f3)
	fmt.Println(f)

	fmt.Println("---------------------")
	f = Foo{A: 10, b: "hello"}
	b1 := Bar{C: "Fred", F: f}
	fmt.Println(b1.F)
	b2 := Baz{D: "Nano", Foo: f}
	fmt.Println(b2.A)

	fmt.Println("---------------------")
	bob := `{"name":"Bob", "age": 20}`
	var b Person
	json.Unmarshal([]byte(bob), &b)
	fmt.Println(b.Age, b.Name)
	bob2, _ := json.Marshal(b)
	fmt.Println(string(bob2))
}
