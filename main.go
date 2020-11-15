package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learn how to code with golang!")
	numberOfBytesWritten, _ := fmt.Println(0, 0, true)
	fmt.Println(numberOfBytesWritten)

	// x is using the short declaration operator here and this declares AND assigns a value.
	x := 42
	fmt.Println(x)

	y := 100 + 42

	fmt.Println(y)

	z := "Dom, Hallan - November 14, 2021"
	fmt.Println(z)
}
