package main

import "fmt"

func main() {
	fmt.Println("Pointers")
	a := 23
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b)

	x := 0
	foo(&x)
	fmt.Println(x)
}

func foo(y *int) {
	fmt.Println(y)
	*y = 43
	fmt.Println(y)
}
