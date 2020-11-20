package main

import (
	"fmt"
)

type hotdog int

var b hotdog

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	b = 52
	fmt.Printf("%T\n", b)
}
