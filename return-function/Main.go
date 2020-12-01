package main

import "fmt"

func main() {
	fmt.Println("Hello, playground!")
	s1 := foo()

	fmt.Println(s1)

	s2 := bar(32)
	fmt.Printf("%T\n", s2)
}

type greeting struct {
	hello string
}

func foo() greeting {
	return greeting{
		hello: "Hello, World!",
	}
}

func bar(x int) func() int {
	return func() int {
		fmt.Println(x)
		return x
	}
}
