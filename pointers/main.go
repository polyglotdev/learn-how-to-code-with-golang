package main

import "fmt"

func main() {
	fmt.Println("Pointers")
	a := 23
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
}
