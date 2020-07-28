package main

import "fmt"

type Foo struct {
	A int
	B string
}

func (f Foo) String() string {
	return fmt.Sprintf("It has %d fields: A: %d and B: %s", f.fieldCount(), f.A, f.B)
}

func (f *Foo) Double() {
	f.A = f.A * 2
}

func (f Foo) fieldCount() int {
	return 2
}

type Bar struct {
	C bool
	Foo
}

func (b Bar) String() string {
	return fmt.Sprintf("C: %v and Foo: %s", b.C, b.Foo.String())
}

func (b Bar) fieldCount() int {
	return 3
}

type myInt int

func (mi myInt) isEven() bool {
	return mi%2 == 0
}

func (mi *myInt) Double() {
	*mi = *mi * 2
}
func main() {
	f := Foo{
		A: 10,
		B: "Heey",
	}
	fmt.Println(f.String())

	f.Double()

	fmt.Println(f.String())

	b := Bar{
		C:   true,
		Foo: f,
	}
	fmt.Println(b.String())
	b.Double()
	fmt.Println(b.String())

	mi := myInt(5)
	fmt.Println(mi)
	fmt.Println(mi.isEven())
	mi.Double()
	fmt.Println(mi.isEven())

}
